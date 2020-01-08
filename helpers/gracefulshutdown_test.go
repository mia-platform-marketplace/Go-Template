package helpers

import (
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
	"time"

	"github.com/sirupsen/logrus/hooks/test"
	"gotest.tools/assert"
)

func TestGracefuleShutdown(t *testing.T) {
	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
	}

	var mtx sync.Mutex
	var listenerError error
	go func() {
		mtx.Lock()
		defer mtx.Unlock()

		listenerError = srv.ListenAndServe()
	}()

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, syscall.SIGTERM)
	log, _ := test.NewNullLogger()

	go func() {
		GracefulShutdown(srv, interruptChan, log)
	}()

	interruptChan <- syscall.SIGTERM

	time.Sleep(500 * time.Millisecond)

	mtx.Lock()
	defer mtx.Unlock()
	assert.Equal(t, listenerError, http.ErrServerClosed, "Listener server not close correctly")
}
