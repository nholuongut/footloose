package ignite

import (
	"github.com/nholuongut/footloose/pkg/exec"
)

// Stop stops an Ignite VM
func Stop(name string) error {
	return exec.CommandWithLogging(execName, "stop", name)
}
