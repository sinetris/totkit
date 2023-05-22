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

// statusVMCmd represents the 'vm status' command
var statusVMCmd = &cobra.Command{
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
	virtualizationCmd.AddCommand(vmCmd)
	vmCmd.AddCommand(createVMCmd, destroyVMCmd, statusVMCmd)
}

func runECreateVM(cmd *cobra.Command, args []string) error {
	settings := viper.AllSettings()
	log.Println(settings)
	viper.GetString("logfile")
	name := "netname"
	cidr := "192.168.123.0/24"
	natnet := fmt.Sprintf("natnetwork add --netname %s --network %s --ipv6 off --dhcp on --enable", name, cidr)
	networkResult, err := vbox.VBoxManage(natnet)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("VirtualBox version %s\n", networkResult)

	// VBoxManage createvm --name ${vbox_machine_id} --ostype Ubuntu22_LTS_64 --register
	// VBoxManage createvm <--name=name> [--basefolder=basefolder] [--group=group-ID,...] [--ostype=ostype] \
	//   [--register] [--uuid=uuid] [--cipher cipher] [--password-id password-id] [--password file]
	// VBoxManage modifyvm ${vbox_machine_id} --audio none
	// VBoxManage modifyvm ${vbox_machine_id} --hpet on
	// VBoxManage modifyvm ${vbox_machine_id} --nested-hw-virt on
	// VBoxManage modifyvm ${vbox_machine_id} --hwvirtex on
	// VBoxManage clonevm ${vbox_machine_id} --register --name <MACHINE_CLONE_TMP_ID --snapshot base --options link

	fmt.Println("VMs: " + strings.Join(args, " "))
	version, err := vbox.VBoxManage("--version")
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("VirtualBox version %s\n", version)
	return nil
}
