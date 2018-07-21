package main

import (
	"bufio"
	"os/exec"
)

func ExecuteWindowsCmd(cmdString string) (stdout, stderr []string, err error) {
	cmd := exec.Command("cmd.exe", "/c", cmdString)
	outPipe, err := cmd.StdoutPipe()
	if err != nil {
		return
	}
	errPipe, err := cmd.StderrPipe()
	if err != nil {
		return
	}
	err = cmd.Start()
	if err != nil {
		return
	}
	defer cmd.Wait()
	scannerOut := bufio.NewScanner(outPipe)
	stdout = make([]string, 0)
	for scannerOut.Scan() {
		stdout = append(stdout, scannerOut.Text())
	}
	scannerErr := bufio.NewScanner(errPipe)
	stderr = make([]string, 0)
	for scannerErr.Scan() {
		stderr = append(stderr, scannerErr.Text())
	}
	return
}
