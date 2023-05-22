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
	"fmt"
	"log"

	vn "sinetris.info/totkit/pkg/virtualization"
)

type VBox struct{}

type VBoxVM struct {
	vbox *VBox
	uuid string
	name string
}

type VBoxNetwork struct {
	vbox *VBox
	Name string
	CIDR string
}

func (vbox *VBox) GetName() string {
	return "vbox"
}

func (vbox *VBox) GetVersion() (string, error) {
	version, err := VBoxManage("--version")
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return version, nil
}

func (vbox *VBox) VirtualMachine() vn.VirtualMachine {
	return &VBoxVM{
		vbox: vbox,
	}
}

func (vbox *VBox) Network() vn.Network {
	return &VBoxNetwork{
		vbox: vbox,
	}
}

func (box *VBoxVM) Create(vm_config vn.VirtualMachineConfig) error {
	box.uuid = ""
	box.name = vm_config.Name
	return nil
}

func (box *VBoxVM) GetState(name string) error {
	log.Printf("VM %s : %v", name, box.uuid)
	return nil
}

func (net *VBoxNetwork) Create(nw vn.NetworkConfig) error {
	vbox_args := []string{"natnetwork", "add", "--netname", nw.Name, "--network", nw.CIDR, "--ipv6", "off"}
	networkResult, err := VBoxManage(vbox_args...)
	if err != nil {
		fmt.Printf("Virtualization error for %v\n", vbox_args)
		log.Fatal(err)
	}
	fmt.Printf("VirtualBox network %s\n", networkResult)
	return err
}

func init() {
	vbox := new(VBox)
	vn.RegisterProvisioner(vbox.GetName(), vbox)
}
