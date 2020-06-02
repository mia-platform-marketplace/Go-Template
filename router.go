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
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type HelloWorld struct {
	Msg string `json:"msg"`
}

func setupRouter(router *mux.Router) {
	// Setup your routes here.
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		helloWorld := HelloWorld{
			Msg: "Hello world!",
		}
		body, err := json.Marshal(&helloWorld)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write(nil)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	})
}
