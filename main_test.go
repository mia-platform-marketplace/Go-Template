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

	"gotest.tools/assert"
)

func TestEntrypoint(t *testing.T) {
	t.Run("opens server on port 3000", func(t *testing.T) {
		shutdown := make(chan os.Signal, 1)

		os.Setenv("HTTP_PORT", "3000")

		go func() {
			entrypoint(shutdown)
		}()
		defer func() {
			os.Unsetenv("HTTP_PORT")
			shutdown <- syscall.SIGTERM
		}()

		time.Sleep(1 * time.Second)

		resp, err := http.DefaultClient.Get("http://localhost:3000/")
		assert.Equal(t, err, nil)
		assert.Equal(t, resp.StatusCode, 200)
	})

	t.Run("sets correct path prefix", func(t *testing.T) {
		shutdown := make(chan os.Signal, 1)

		os.Setenv("SERVICE_PREFIX", "/prefix")
		go func() {
			entrypoint(shutdown)
		}()
		defer func() {
			os.Unsetenv("SERVICE_PREFIX")
			shutdown <- syscall.SIGTERM
		}()

		time.Sleep(1 * time.Second)

		resp, err := http.DefaultClient.Get("http://localhost:8080/prefix/")
		assert.Equal(t, err, nil)
		assert.Equal(t, resp.StatusCode, 200)
	})
	
	t.Run("GracefulShutdown works properly", func(t *testing.T) {
		shutdown := make(chan os.Signal, 1)
		flag := false

		go func() {
			entrypoint(shutdown)
			time.Sleep(1 * time.Second)
			flag = true
		}()
		shutdown <- syscall.SIGTERM

		time.Sleep(2 * time.Second)
		assert.Equal(t, flag, true)
	})
}
