package common

import (
	"bufio"
	"io/ioutil"
	"os"
	"os/exec"
)

// stdout,stderr遇到中文有乱码，先保留，观察是否需要在服务端处理
func RunExe(cmdString string) (stdout, stderr string, err error) {
	// 将命令写入文件再执行 防止命令的参数丢失

	f, err := ioutil.TempFile("", "regular")
	if err != nil {
		return
	}
	f.Close()
	tmpFile := f.Name()
	os.Remove(tmpFile)
	tmpFile = tmpFile + ".bat"
	d1 := []byte("@echo off \r\n" + cmdString)
	err = ioutil.WriteFile(tmpFile, d1, 0644)
	if err != nil {
		return
	}
	defer os.Remove(tmpFile)
	cmd := exec.Command("cmd.exe", "/C", tmpFile)
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
	for scannerOut.Scan() {
		stdout += "\r\n" + scannerOut.Text()
	}
	scannerErr := bufio.NewScanner(errPipe)
	for scannerErr.Scan() {
		stderr += "\r\n" + scannerErr.Text()
	}

	return
}
