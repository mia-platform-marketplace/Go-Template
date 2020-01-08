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
	"github.com/mia-platform/configlib"
)

// EnvironmentVariables struct with the mapping of desired
// environment variables.
type EnvironmentVariables struct {
	LogLevel            string
	HTTPPort            string
	UserIDHeaderKey     string
	GroupsHeaderKey     string
	ClienttypeHeaderKey string
	BackofficeHeaderKey string
	ServicePrefix       string
	ServiceVersion      string
}

var envVariablesConfig = []configlib.EnvConfig{
	{
		Key:          "LOG_LEVEL",
		Variable:     "LogLevel",
		DefaultValue: "info",
	},
	{
		Key:          "HTTP_PORT",
		Variable:     "HTTPPort",
		DefaultValue: "8080",
	},
	{
		Key:          "USERID_HEADER_KEY",
		Variable:     "UserIdHeaderKey",
		DefaultValue: "userid",
	},
	{
		Key:          "GROUPS_HEADER_KEY",
		Variable:     "GroupsHeaderKey",
		DefaultValue: "usergroups",
	},
	{
		Key:          "CLIENTTYPE_HEADER_KEY",
		Variable:     "ClienttypeHeaderKey",
		DefaultValue: "clienttype",
	},
	{
		Key:          "BACKOFFICE_HEADER_KEY",
		Variable:     "BackofficeHeaderKey",
		DefaultValue: "isbackoffice",
	},
	{
		Key:      "SERVICE_PREFIX",
		Variable: "ServicePrefix",
	},
	{
		Key:      "SERVICE_VERSION",
		Variable: "ServiceVersion",
	},
}
