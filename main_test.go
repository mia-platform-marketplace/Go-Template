/*
 * Copyright 2019 Mia srl
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
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

		t.Setenv("HTTP_PORT", "3000")
		t.Setenv("SERVICE_VERSION", "myVersion")

		go func() {
			entrypoint(shutdown)
		}()
		defer func() { shutdown <- syscall.SIGTERM }()

		time.Sleep(1 * time.Second)

		resp, err := http.DefaultClient.Get("http://localhost:3000/-/healthz")
		require.Equal(t, nil, err)
		require.Equal(t, 200, resp.StatusCode)
	})

	t.Run("sets correct path prefix", func(t *testing.T) {
		shutdown := make(chan os.Signal, 1)

		t.Setenv("SERVICE_VERSION", "myVersion")
		t.Setenv("SERVICE_PREFIX", "/prefix")

		go func() {
			entrypoint(shutdown)
		}()
		defer func() { shutdown <- syscall.SIGTERM }()

		time.Sleep(1 * time.Second)

		resp, err := http.DefaultClient.Get("http://localhost:8080/prefix/")
		require.Equal(t, nil, err)
		require.Equal(t, 404, resp.StatusCode)
	})

	t.Run("shutdown works properly", func(t *testing.T) {
		t.Setenv("SERVICE_VERSION", "myVersion")
		t.Setenv("DELAY_SHUTDOWN_SECONDS", "3")

		shutdown := make(chan os.Signal, 1)
		done := make(chan bool, 1)

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
