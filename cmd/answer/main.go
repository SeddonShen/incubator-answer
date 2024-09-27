/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

//go:generate go run github.com/swaggo/swag/cmd/swag init -g ./cmd/answer/main.go -d ../../ -o ../../docs

package main

import (
	_ "github.com/Anan1225/incubator-answer-plugins/user-center-slack"
	_ "github.com/Anan1225/incubator-answer-plugins/user-center-wecom"
	_ "github.com/SeddonShen/slack_plugin_seddon"
	_ "github.com/apache/incubator-answer-plugins/connector-github"
	answercmd "github.com/apache/incubator-answer/cmd"
)

// main godoc
// @title 	"apache answer"
// @description = "apache answer api"
// @version = "v0.0.1"
// @BasePath = "/"
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	answercmd.Main()
}
