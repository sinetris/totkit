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

package vm

import "fmt"

var Provisioners = make(map[string]Provisioner)

type Provisioner interface {
	GetName() string
	GetVersion() (string, error)
	VirtualMachine() VirtualMachine
	Network() Network
}

type VirtualMachine interface {
	Create(VirtualMachineConfig) error
	GetState(name string) error
}

type VirtualMachineConfig struct {
	Name  string
	State string
}

type Network interface {
	Create(NetworkConfig) error
}

type NetworkConfig struct {
	Name string
	CIDR string
}

func RegisterProvisioner(name string, provisioner Provisioner) {
	if _, found := Provisioners[name]; found {
		panic(fmt.Errorf("registering duplicate provisioner '%s'", name))
	}
	Provisioners[name] = provisioner
}

func GetProvisioner(name string) (*Provisioner, error) {
	var err error
	provisioner, found := Provisioners[name]
	if !found {
		err = fmt.Errorf("provisioner '%s' doesn't exist", name)
	}
	return &provisioner, err
}
