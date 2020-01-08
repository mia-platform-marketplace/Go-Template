
// Package helpers is a set of utility commonly used by http servers
package helpers

import (
	"context"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

// GracefulShutdown waits on notified signal to shutdown until all connections are closed.
func GracefulShutdown(srv *http.Server, interruptChan chan os.Signal, logger *logrus.Logger) {
	// Block until we receive our signal.
	<-interruptChan

	if err := srv.Shutdown(context.Background()); err != nil {
		logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Error during shutdown, forcing close.")
		if err := srv.Close(); err != nil {
			logger.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("Error during server close.")
		}
	}
}
