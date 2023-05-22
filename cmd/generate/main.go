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

package generate

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"sinetris.info/totkit/cmd"
)

// generateCmd represents the vm command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "CLI command to Generate code",
	Long: heredoc.Doc(`
		CLI code generator

		Look at the sub commands to generate configs, docs, etc
	`),
}

func init() {
	cmd.RootCmd.AddCommand(generateCmd)
}
