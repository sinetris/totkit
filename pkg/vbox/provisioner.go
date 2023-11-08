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

	vrt "sinetris.info/totkit/pkg/virtualization"
)

// Implement vrt.NetworkProvisioner

type VBoxkNetworkProvisioner struct {
	provisioner *VBoxProvisioner
	Name        string
	CIDR        string
}

func (n *VBoxkNetworkProvisioner) Create(nc vrt.NetworkConfig) error {
	vboxArgs := []string{"natnetwork", "add", "--netname", nc.Name, "--network", nc.CIDR, "--ipv6", "off"}
	res, err := VBoxManage(vboxArgs...)
	if err != nil {
		log.Fatalf("VBoxManage: %v\nError: %v\n", res, err)
	} else {
		fmt.Printf("VirtualBox: %s\n", res.Stdout)
	}
	return err
}

func (vm *VBoxkNetworkProvisioner) Destroy(name string) error {
	vboxArgs := []string{"natnetwork", "remove", "--netname", name}
	res, err := VBoxManage(vboxArgs...)
	if err != nil {
		log.Fatalf("VBoxManage: %v\nError: %v\n", res, err)
	} else {
		fmt.Printf("VirtualBox: %s\n", res.Stdout)
	}

	return err
}

func (n *VBoxkNetworkProvisioner) State(name string) (vrt.Network, error) {
	res := new(vrt.Network)
	return *res, nil
}

func (n *VBoxkNetworkProvisioner) List() (map[string]vrt.Network, error) {
	vboxArgs := []string{"natnetwork", "list"}
	res, err := VBoxManage(vboxArgs...)
	if err != nil {
		log.Fatalf("VBoxManage: %v\nError: %v\n", res, err)
	}
	fmt.Printf("VirtualBox: %s\n", res.Stdout)
	networks := make(map[string]vrt.Network)
	return networks, err
}

// Implement vrt.VirtualMachineProvisioner

type VBoxVMProvisioner struct {
	provisioner *VBoxProvisioner
	uuid        string
	name        string
}

func (vm *VBoxVMProvisioner) Create(vmConfig vrt.VirtualMachineConfig) error {
	vm.uuid = ""
	vm.name = vmConfig.Name
	return nil
}

func (vm *VBoxVMProvisioner) Destroy(name string) error {
	vboxArgs := []string{"destroyvm", "--name", name}
	res, err := VBoxManage(vboxArgs...)
	if err != nil {
		log.Fatalf("VBoxManage: %v\nError: %v\n", res, err)
	} else {
		fmt.Printf("VirtualBox: %s\n", res.Stdout)
	}

	return err
}

func (vm *VBoxVMProvisioner) State(name string) (vrt.VirtualMachine, error) {
	log.Printf("VM %s : %v", name, vm.uuid)
	res := new(vrt.VirtualMachine)
	return *res, nil
}

func (vm *VBoxVMProvisioner) List() (map[string]vrt.VirtualMachine, error) {
	res := make(map[string]vrt.VirtualMachine)

	return res, nil
}

// Implement vrt.Provisioner

type VBoxProvisioner struct{}

func (p *VBoxProvisioner) Name() string {
	return "vbox"
}

func (p *VBoxProvisioner) Version() (string, error) {
	res, err := VBoxManage("--version")
	if err != nil {
		log.Fatalf("VBoxManage: %v\nError: %v\n", res, err)
		return "", err
	}
	return res.Stdout, nil
}

func (p *VBoxProvisioner) VirtualMachine() vrt.VirtualMachineProvisioner {
	return &VBoxVMProvisioner{
		provisioner: p,
	}
}

func (p *VBoxProvisioner) Network() vrt.NetworkProvisioner {
	return &VBoxkNetworkProvisioner{
		provisioner: p,
	}
}

func init() {
	provisioner := new(VBoxProvisioner)
	vrt.RegisterProvisioner(provisioner.Name(), provisioner)
}
