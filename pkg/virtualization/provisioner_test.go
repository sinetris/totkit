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
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockProvisioner struct {
	networkProvisioner *MockNetworkProvisioner
	vmProvisioner      *MockVirtualMachineProvisioner
}

func (p *MockProvisioner) Name() string {
	return "mocked-provisioner"
}

func (p *MockProvisioner) Version() (string, error) {
	return "v0.0.1", nil
}

func (p *MockProvisioner) VirtualMachine() VirtualMachineProvisioner {
	if p.vmProvisioner == nil {
		p.vmProvisioner = new(MockVirtualMachineProvisioner)
	}
	return p.vmProvisioner
}

func (p *MockProvisioner) Network() NetworkProvisioner {
	if p.networkProvisioner == nil {
		p.networkProvisioner = new(MockNetworkProvisioner)
	}
	return p.networkProvisioner
}

func TestProvisioner(t *testing.T) {
	// name := "mocked-provisioner"
	// p := new(MockProvisioner)
	assert.Equal(t, 123, 123, "they should be equal")

	// if diff := cmp.Diff(SetDescription{
	// 	Version:        "1.1.1",
	// 	SDKVersion:     sdkVersion.String(),
	// 	APIVersion:     "x" + APIVersionMajor + "." + APIVersionMinor,
	// 	Builders:       []string{"example", "example-2"},
	// 	PostProcessors: []string{"example", "example-2"},
	// 	Provisioners:   []string{"example", "example-2"},
	// 	Datasources:    []string{"example", "example-2"},
	// }, outputDesc); diff != "" {
	// 	t.Fatalf("Unexpected description: %s", diff)
	// }
	// if err != nil {
	// 	t.Errorf("Unexpected error %d instead.\n", err)
	// }

	// if res.CIDR != nc.CIDR {
	// 	t.Errorf("Expected Network CIDR to be the same as NetworkConfig CIDR, got %d instead.\n", res.CIDR)
	// }
}
