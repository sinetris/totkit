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
	"fmt"
	"log"

	"github.com/MakeNowJust/heredoc"
	"github.com/google/go-jsonnet"
	"github.com/spf13/cobra"
)

type Config struct {
	ConfigFile string
	JSON       string
}

var cfg Config

// cfgCmd represents the vm command
var cfgCmd = &cobra.Command{
	Use:   "config",
	Short: "Generate CLI configurations using Jsonnet",
	Long: heredoc.Doc(`
		Generate CLI configurations using Jsonnet

		This configuration can be later passed to the CLI using '--config'
	`),
	RunE: func(cmd *cobra.Command, args []string) error {
		vmResult, err := JsonnetVM(cmd)
		if err != nil {
			fmt.Printf("JsonnetVM config Error in file: %s\n%v", cfg.ConfigFile, err)
			return err
		}

		cfg.JSON = vmResult
		fmt.Printf("JsonnetVM config from file: %s\n%v", cfg.ConfigFile, cfg.JSON)
		return nil
	},
}

func init() {
	generateCmd.AddCommand(cfgCmd)

	cfgCmd.PersistentFlags().StringVar(&cfg.ConfigFile, "jsonnet-config", "", "Path to a jsonnet file to configure VMs.")
	err := cfgCmd.MarkPersistentFlagFilename("jsonnet-config", "jsonnet")
	if err != nil {
		completeError := fmt.Sprintf("Configuration Error for 'jsonnet-config' \n%v", err)
		log.Fatal(completeError)
	}
}

func JsonnetVM(cmd *cobra.Command) (string, error) {
	vm := jsonnet.MakeVM()

	jsonStr, err := vm.EvaluateFile(cfg.ConfigFile)
	if err != nil {
		return "", err
	}

	return jsonStr, nil
}
