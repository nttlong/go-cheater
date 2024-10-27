// This is the implementation of the ProcessManager class in Go.
package process_manager

import (
	"github.com/mitchellh/go-ps"
)

type ProcessManager struct {
	processes []ps.Process
}

func NewProcessManager() *ProcessManager {
	return &ProcessManager{}
}
