// +build linux darwin

package qemu

import (
	"os/exec"

	"github.com/nanovms/ops/types"
)

func checkExists(key string) bool {
	_, err := exec.LookPath(key)
	if err != nil {
		return false
	}
	return true
}

// HypervisorInstance provides available hypervisor
func HypervisorInstance() Hypervisor {
	for k := range hypervisors {
		if checkExists(k) {
			hypervisor := hypervisors[k]()
			return hypervisor
		}
	}
	return nil
}

// Hypervisor interface
type Hypervisor interface {
	Start(rconfig *types.RunConfig) error
	Command(rconfig *types.RunConfig) *exec.Cmd
	Stop()
	PID() (string, error)
}

// available hypervisors
var hypervisors = map[string]func() Hypervisor{
	"qemu-system-x86_64": newQemu,
}
