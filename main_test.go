/*
 * Copyright 2023 Mia srl
 * All rights reserved.
 */

package main

import (
	"net/http"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestEntryPoint(t *testing.T) {
	t.Run("opens server on port 3000", func(t *testing.T) {
		shutdown := make(chan os.Signal, 1)

		os.Setenv("HTTP_PORT", "3000")
		os.Setenv("SERVICE_VERSION", "myVersion")

		go func() {
			entrypoint(shutdown)
		}()
		defer func() {
			os.Unsetenv("HTTP_PORT")
			os.Unsetenv("SERVICE_VERSION")
			shutdown <- syscall.SIGTERM
		}()

		time.Sleep(1 * time.Second)

		resp, err := http.DefaultClient.Get("http://localhost:3000/-/healthz")
		require.Equal(t, nil, err)
		require.Equal(t, 200, resp.StatusCode)
	})

	t.Run("sets correct path prefix", func(t *testing.T) {
		shutdown := make(chan os.Signal, 1)

		os.Setenv("SERVICE_VERSION", "myVersion")
		os.Setenv("SERVICE_PREFIX", "/prefix")

		go func() {
			entrypoint(shutdown)
		}()
		defer func() {
			os.Unsetenv("SERVICE_VERSION")
			os.Unsetenv("SERVICE_PREFIX")
			shutdown <- syscall.SIGTERM
		}()

		time.Sleep(1 * time.Second)

		resp, err := http.DefaultClient.Get("http://localhost:8080/prefix/")
		require.Equal(t, nil, err)
		require.Equal(t, 404, resp.StatusCode)
	})

	t.Run("shutdown works properly", func(t *testing.T) {
		os.Setenv("SERVICE_VERSION", "myVersion")
		os.Setenv("DELAY_SHUTDOWN_SECONDS", "3")

		shutdown := make(chan os.Signal, 1)
		done := make(chan bool, 1)

		defer func() {
			os.Unsetenv("SERVICE_VERSION")
			os.Unsetenv("DELAY_SHUTDOWN_SECONDS")
		}()

		go func() {
			time.Sleep(5 * time.Second)
			done <- false
		}()

		go func() {
			entrypoint(shutdown)
			done <- true
		}()

		shutdown <- syscall.SIGTERM

		flag := <-done
		require.True(t, flag)
	})
}
