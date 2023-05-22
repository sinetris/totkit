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
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"sinetris.info/totkit/pkg/vbox"
)

// networkCmd represents the 'network' command
var networkCmd = &cobra.Command{
	Use:   "network",
	Short: "Manage Virtual Machine resources",
	Long:  "Manage Virtual Machine resources",
}

// createNetworkCmd represents the 'network create' command
var createNetworkCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Virtual Machines network",
	Long:  "Create a new Virtual Machines network",
	RunE:  runECreateNetwork,
}

// destroyNetworkCmd represents the 'network destroy' command
var destroyNetworkCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy a Virtual Machines",
	Long:  "Destroy a Virtual Machines",
	Run: func(cmd *cobra.Command, args []string) {
		// VBoxManage unregistervm ${vbox_machine_id} --delete --machinereadable
		// VBoxManage destroyvm --name ${vbox_machine_id} --ostype Ubuntu22_LTS_64 --register
		// VBoxManage natnetwork remove --netname ${net_name}

		fmt.Println("VMs: " + strings.Join(args, " "))
		version, err := vbox.VBoxManage("--version")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("VirtualBox version %s\n", version)
	},
}

// statusNetworkCmd represents the 'network status' command
var statusNetworkCmd = &cobra.Command{
	Use:   "status",
	Short: "Check Virtual Machines status",
	Long:  "Check Virtual Machines status",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("VMs: " + strings.Join(args, " "))
		// VBoxManage list vms
		// VBoxManage list systemproperties
		// VBoxManage list bridgedifs
		// VBoxManage list hdds
		version, err := vbox.VBoxManage("--version")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("VirtualBox version %s\n", version)
	},
}

func init() {
	virtualizationCmd.AddCommand(networkCmd)
	networkCmd.AddCommand(createNetworkCmd, destroyNetworkCmd, statusNetworkCmd)
}

func runECreateNetwork(cmd *cobra.Command, args []string) error {
	viper.GetString("logfile")
	name := "netname1"
	cidr := "192.168.123.0/24"
	vbox_args := []string{"natnetwork", "add", "--netname", name, "--network", cidr, "--ipv6", "off"}
	networkResult, err := vbox.VBoxManage(vbox_args...)
	if err != nil {
		fmt.Printf("Virtualization error for %v\n", vbox_args)
		log.Fatal(err)
	}
	fmt.Printf("VirtualBox network %s\n", networkResult)
	return err
}
