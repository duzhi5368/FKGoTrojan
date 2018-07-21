package client_singleton

import (
	"os"
	"os/exec"
)

func slaveStart() error {
	cmd := exec.Command(os.Args[0])
	go cmd.Run()
	return nil
}
