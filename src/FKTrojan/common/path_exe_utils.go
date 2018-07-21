package common

import (
	"path/filepath"

	"fmt"

	"os/exec"

	"github.com/StackExchange/wmi"
)

type Win32_Process struct {
	Name           string
	ExecutablePath *string
	ProcessID      int
}

func RunningCount(exePath string) (int, error) {
	if !PathExist(exePath) {
		return 0, fmt.Errorf("path %s not exist", exePath)
	}
	var dst []Win32_Process
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		return 0, err
	}
	myName := filepath.Base(exePath)
	myMd5 := Md5HashStringFile(exePath)
	count := 0
	for _, v := range dst {
		exeName := filepath.Base(*v.ExecutablePath)
		if myName != exeName {
			continue
		}
		exeMd5 := Md5HashStringFile(*v.ExecutablePath)
		if exeMd5 == myMd5 {
			count++
		}
	}
	return count, nil
}
func KillPathExe(exePath string) error {
	if !PathExist(exePath) {
		return fmt.Errorf("path %s not exist", exePath)
	}
	var dst []Win32_Process
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		return err
	}
	myName := filepath.Base(exePath)
	myMd5 := Md5HashStringFile(exePath)
	ids := make([]int, 0)
	for _, v := range dst {
		exeName := filepath.Base(*v.ExecutablePath)
		if myName != exeName {
			continue
		}
		exeMd5 := Md5HashStringFile(*v.ExecutablePath)
		if exeMd5 == myMd5 {
			ids = append(ids, v.ProcessID)
		} else {
			fmt.Printf("file %s md5 is %s not %s\n", *v.ExecutablePath, exeMd5, myMd5)
		}
	}
	if len(ids) == 0 {
		return nil
	}
	pids := make([]string, 0)
	pids = append(pids, "/F")
	for _, id := range ids {
		pids = append(pids, "/PID")
		pids = append(pids, fmt.Sprintf("%d", id))
	}
	//fmt.Println(pids)
	cmd := exec.Command("taskkill", pids...)

	return cmd.Run()
}
