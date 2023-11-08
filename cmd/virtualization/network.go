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

// networkCmd represents the 'network' command
var networkCmd = &cobra.Command{
	Use:   "network",
	Short: "Manage Virtualized network resources",
	Long:  "Manage Virtualized network resources",
}

// createNetworkCmd represents the 'network create' command
var createNetworkCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Virtualized network",
	Long:  "Create a new Virtualized network",
	RunE:  runECreateNetwork,
}

// destroyNetworkCmd represents the 'network destroy' command
var destroyNetworkCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy a Virtualized network",
	Long:  "Destroy a Virtualized network",
	RunE:  runEDestroyNetwork,
}

// statusNetworkCmd represents the 'network status' command
var statusNetworkCmd = &cobra.Command{
	Use:   "status",
	Short: "Check Virtualized network status",
	Long:  "Check Virtualized network status",
	RunE:  runEStatusNetwork,
}

// listNetworkCmd represents the 'network list' command
var listNetworkCmd = &cobra.Command{
	Use:   "list",
	Short: "List Virtualized network",
	Long:  "List Virtualized network",
	RunE:  runEListNetwork,
}

func init() {
	virtualizationCmd.AddCommand(networkCmd)
	networkCmd.AddCommand(createNetworkCmd, destroyNetworkCmd, statusNetworkCmd, listNetworkCmd)
}

func runECreateNetwork(cmd *cobra.Command, args []string) error {
	provisioner, err := virtualization.GetProvisioner("vbox")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	name := "netname1"
	cidr := "192.168.123.0/24"
	err = provisioner.Network().Create(
		virtualization.NetworkConfig{
			Name: name,
			CIDR: cidr,
		})
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	return err
}

func runEDestroyNetwork(cmd *cobra.Command, args []string) error {
	provisioner, err := virtualization.GetProvisioner("vbox")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	name := "netname1"
	err = provisioner.Network().Destroy(name)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	return err
}

func runEStatusNetwork(cmd *cobra.Command, args []string) error {
	provisioner, err := virtualization.GetProvisioner("vbox")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	name := "netname1"
	state, err := provisioner.Network().State(name)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Result: %v\n", state)
	}
	return err
}

func runEListNetwork(cmd *cobra.Command, args []string) error {
	provisioner, err := virtualization.GetProvisioner("vbox")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	state, err := provisioner.Network().List()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Result: %v\n", state)
	}
	return err
}
