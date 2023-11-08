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

	"github.com/spf13/cobra"
	"sinetris.info/totkit/pkg/virtualization"
)

// vmCmd represents the 'vm' command
var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "Manage Virtual Machine resources",
	Long:  "Manage Virtual Machine resources",
}

// createVMCmd represents the 'vm create' command
var createVMCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Virtual Machine",
	Long:  "Create a new Virtual Machine",
	RunE:  runECreateVM,
}

// destroyVMCmd represents the 'vm destroy' command
var destroyVMCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy a Virtual Machines",
	Long:  "Destroy a Virtual Machines",
	RunE:  runEDestroyVM,
}

// statusVMCmd represents the 'vm status' command
var statusVMCmd = &cobra.Command{
	Use:   "status",
	Short: "Check Virtual Machines status",
	Long:  "Check Virtual Machines status",
	RunE:  runEStateVM,
}

func init() {
	virtualizationCmd.AddCommand(vmCmd)
	vmCmd.AddCommand(createVMCmd, destroyVMCmd, statusVMCmd)
}

func runECreateVM(cmd *cobra.Command, args []string) error {
	provisioner, err := virtualization.GetProvisioner("vbox")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	vmProvisioner := provisioner.VirtualMachine()
	name := "vm1"
	err = vmProvisioner.Create(
		virtualization.VirtualMachineConfig{
			Name: name,
		})
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	return err
}

func runEDestroyVM(cmd *cobra.Command, args []string) error {
	provisioner, err := virtualization.GetProvisioner("vbox")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	vmProvisioner := provisioner.VirtualMachine()
	name := "vm1"
	err = vmProvisioner.Destroy(name)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	return err
}

func runEStateVM(cmd *cobra.Command, args []string) error {
	provisioner, err := virtualization.GetProvisioner("vbox")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	vmProvisioner := provisioner.VirtualMachine()
	name := "vm1"
	state, err := vmProvisioner.State(name)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Result: %v\n", state)
	}
	return err
}
