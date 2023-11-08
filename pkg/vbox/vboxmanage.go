// Copyright © 2023 Duilio Ruggiero <duilio@sinetris.info>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vbox

import (
	"bytes"
	"context"
	"os/exec"
	"time"
)

const VBoxCLI string = "VBoxManage"

type VBoxManageRawResult struct {
	Args     []string
	Stdout   string
	Stderr   string
	ExitCode int
}

func VBoxManage(args ...string) (VBoxManageRawResult, error) {
	res := new(VBoxManageRawResult)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	var stdout, stderr bytes.Buffer
	cmd := exec.CommandContext(ctx, VBoxCLI, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	res.Stdout = stdout.String()
	res.Stderr = stderr.String()

	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			res.ExitCode = exiterr.ExitCode()
		} else {
			res.ExitCode = -1
		}
	} else {
		res.ExitCode = 0
	}

	return *res, err
}
