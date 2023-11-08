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

package virtualization

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"sinetris.info/totkit/cmd"
	"sinetris.info/totkit/pkg/virtualization"
)

type VMOptions struct {
	VMConfigFile string
	VMConfig     string
	DryRun       bool
}

// virtualizationCmd represents the virtualization command
var virtualizationCmd = &cobra.Command{
	Use:   "virtualization",
	Short: "Manage virtualized resources",
	Long: heredoc.Doc(`
		Manage virtualized resources

		Virtual Machines, Networks, etc.
	`),
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show Virtual Machines hypervisor version",
	Long:  "Check Virtual Machines hypervisor version",
	Run: func(cmd *cobra.Command, args []string) {
		provisioner, err := virtualization.GetProvisioner("vbox")
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		version, err := provisioner.Version()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		fmt.Printf("VirtualBox version %s\n", version)
	},
	Args: cobra.NoArgs,
}

func init() {
	cmd.RootCmd.AddCommand(virtualizationCmd)
	virtualizationCmd.AddCommand(versionCmd)
}
