/* plumber: a deployment pipeline template
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
package pipelines

import (
	"testing"

	"github.com/recruit-tech/plumber/stages"
)

func createCommandStage(command string, arguments ...string) *stages.CommandStage {
	in := make(chan stages.Mediator)
	out := make(chan stages.Mediator)
	return &stages.CommandStage{
		Command:   "echo",
		Arguments: []string{"echo", "baz"},
		BaseStage: stages.BaseStage{
			InputCh:  &in,
			OutputCh: &out,
		},
	}
}

func TestAddPipeline(t *testing.T) {
	pipeline := NewPipeline()
	pipeline.AddStage(stages.InitStage("command"))
	pipeline.AddStage(stages.InitStage("command"))
	expected := 2
	actual := pipeline.Size()
	if expected != actual {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestRunPipeline(t *testing.T) {
	pipeline := NewPipeline()
	pipeline.AddStage(createCommandStage("echo", "foobar"))
	pipeline.AddStage(createCommandStage("echo", "baz"))
	expected := true
	actual := pipeline.Run()
	if expected != actual {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	// TODO: check the output from stage
}