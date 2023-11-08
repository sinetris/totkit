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
	"testing"
)

type MockVirtualMachineProvisioner struct {
	list map[string]VirtualMachine
}

func (p *MockVirtualMachineProvisioner) Create(c VirtualMachineConfig) error {
	if p.list == nil {
		p.list = make(map[string]VirtualMachine)
	}
	n := new(VirtualMachine)
	n.Name = c.Name
	p.list[c.Name] = *n
	return nil
}

func (p *MockVirtualMachineProvisioner) Destroy(name string) error {
	var err error
	if p.list == nil {
		p.list = make(map[string]VirtualMachine)
	}
	if _, ok := p.list[name]; ok {
		delete(p.list, name)
	} else {
		err = fmt.Errorf("VM '%s' doesn't exist", name)
	}
	return err
}

func (p *MockVirtualMachineProvisioner) State(name string) (VirtualMachine, error) {
	var err error
	n, found := p.list[name]
	if !found {
		err = fmt.Errorf("provisioner '%s' doesn't exist", name)
	}
	return n, err
}

func (n *MockVirtualMachineProvisioner) List() (map[string]VirtualMachine, error) {
	return n.list, nil
}

func TestVirtualMachine(t *testing.T) {
	name := "mocked-vm"
	c := VirtualMachineConfig{
		Name: name,
	}

	p := new(MockVirtualMachineProvisioner)
	err := p.Create(c)
	if err != nil {
		t.Errorf("Unexpected error %v.\n", err)
	}

	res, err := p.State(name)
	if err != nil {
		t.Errorf("Unexpected error %v.\n", err)
	}

	if res.Name != c.Name {
		t.Errorf("Expected VirtualMachine Name to be the same as VirtualMachineConfig Name, got %s instead.\n", res.Name)
	}
}
