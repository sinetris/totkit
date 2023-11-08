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

type MockNetworkProvisioner struct {
	list map[string]Network
}

func (p *MockNetworkProvisioner) Create(c NetworkConfig) error {
	if p.list == nil {
		p.list = make(map[string]Network)
	}

	n := new(Network)
	n.CIDR = c.CIDR
	n.Name = c.Name
	p.list[c.Name] = *n
	return nil
}

func (p *MockNetworkProvisioner) Destroy(name string) error {
	var err error
	if p.list == nil {
		p.list = make(map[string]Network)
	}
	if _, ok := p.list[name]; ok {
		delete(p.list, name)
	} else {
		err = fmt.Errorf("Network '%s' doesn't exist", name)
	}
	return err
}

func (p *MockNetworkProvisioner) State(name string) (Network, error) {
	var err error
	n, found := p.list[name]
	if !found {
		err = fmt.Errorf("provisioner '%s' doesn't exist", name)
	}
	return n, err
}

func (n *MockNetworkProvisioner) List() (map[string]Network, error) {
	return n.list, nil
}

func TestNetwork(t *testing.T) {
	name := "mocked-network"
	c := NetworkConfig{
		Name: name,
		CIDR: "10.0.0.0/8",
	}

	p := new(MockNetworkProvisioner)
	err := p.Create(c)
	if err != nil {
		t.Errorf("Unexpected error %v.\n", err)
	}

	res, err := p.State(name)
	if err != nil {
		t.Errorf("Unexpected error %v.\n", err)
	}

	if res.CIDR != c.CIDR {
		t.Errorf("Expected Network CIDR to be the same as NetworkConfig CIDR, got %s instead.\n", res.CIDR)
	}
}
