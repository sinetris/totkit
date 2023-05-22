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
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	rootCmd "sinetris.info/totkit/cmd"
)

// docsCmd represents the vm command
var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generate documentation for totkit CLI",
	Long:  "Generate documentation for totkit CLI",
	Run: func(cmd *cobra.Command, args []string) {
		err := doc.GenMarkdownTree(rootCmd.RootCmd, "./docs/generated")
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	generateCmd.AddCommand(docsCmd)
}
