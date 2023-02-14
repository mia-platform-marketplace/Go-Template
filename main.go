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
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mia-platform/configlib"
	"github.com/mia-platform/glogger/v3"
	"github.com/sirupsen/logrus"
)

func main() {
	entrypoint(make(chan os.Signal, 1))
	os.Exit(0)
}

func entrypoint(shutdown chan os.Signal) {
	var env EnvironmentVariables
	err := configlib.GetEnvVariables(envVariablesConfig, &env)
	if err != nil {
		panic(err.Error())
	}

	// Init logger instance.
	log, err := glogger.InitHelper(glogger.InitOptions{Level: env.LogLevel})
	if err != nil {
		panic(err.Error())
	}

	app, err := setupRouter(env, log)
	if err != nil {
		panic(err.Error())
	}

	go func(app *fiber.App, log *logrus.Logger, env EnvironmentVariables) {
		log.WithField("port", env.HTTPPort).Info("starting server")
		if err := app.Listen(fmt.Sprintf(":%s", env.HTTPPort)); err != nil {
			log.Println(err)
		}
	}(app, log, env)

	signal.Notify(shutdown, syscall.SIGTERM)
	<-shutdown
	time.Sleep(time.Duration(env.DelayShutdownSeconds) * time.Second)
	log.Info("Gracefully shutting down...")
	if err := app.Shutdown(); err != nil {
		panic(err.Error())
	}
}
