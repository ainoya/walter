/* walter: a deployment pipeline template
 * Copyright (C) 2014 Recruit Technologies Co., Ltd. and contributors
 * (see CONTRIBUTORS.md)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package config

import (
	"flag"
	"io/ioutil"

	"github.com/go-yaml/yaml"
	"github.com/recruit-tech/walter/log"
)

var (
	fs = flag.NewFlagSet("walter", flag.ExitOnError)
)

type Opts struct {
	PipelineFilePath string
	StopOnAnyFailure bool
	PrintVersion     bool
}

func LoadOpts(arguments []string) (*Opts, error) {
	var pipelineFilePath string
	var stopOnAnyFailure bool
	var printVersion bool

	fs.StringVar(&pipelineFilePath, "c", "./pipeline.yml", "pipeline.yml file")
	fs.BoolVar(&stopOnAnyFailure, "f", false, "Skip execution of subsequent stage after failing to exec the upstream stage.")
	fs.BoolVar(&printVersion, "v", false, "Print the version information and exit.")

	if err := fs.Parse(arguments); err != nil {
		return nil, err
	}

	return &Opts{
		PipelineFilePath: pipelineFilePath,
		StopOnAnyFailure: stopOnAnyFailure,
		PrintVersion:     printVersion,
	}, nil
}

func ReadConfig(configFilePath string) *map[interface{}]interface{} {
	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Errorf("error :%v \n", err)
	}

	return ReadConfigBytes(data)
}

func ReadConfigBytes(configSetting []byte) *map[interface{}]interface{} {
	configData := make(map[interface{}]interface{})
	err := yaml.Unmarshal(configSetting, &configData)
	if err != nil {
		log.Errorf("error :%v \n", err)
	}
	return &configData
}
