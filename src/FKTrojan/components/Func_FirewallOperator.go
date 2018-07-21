/*
Author: FreeKnight
防火墙设置和管理
 */
//------------------------------------------------------------
package components
//------------------------------------------------------------
import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"

	"github.com/NebulousLabs/go-upnp"
	"errors"
)
//------------------------------------------------------------
// 添加一个文件到防火墙信任中
/*
	name: 进程名，例如："TestApp"
	file: 文件名，例如："C:\\Users\\FreeKnight\\Desktop\\TestApp.exe"
	comment: 必须拥有 管理员 权限
 */
func addFileToFirewall(name string, file string) bool {
	if isAdmin {
		cmd := fmt.Sprintf(`netsh advfirewall firewall add rule name="%s" dir=in action=allow program="%s" enable=yes`, name, file)
		CommandWork := exec.Command("cmd", "/C", cmd)
		CommandWork.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		History, _ := CommandWork.Output()
		if strings.Contains(string(History), "Ok.") {
			FKDebugLog("Add file to firewall successed: " + name)
			return true
		}
		return false
	}
	return false
}
//------------------------------------------------------------
// 使用UPnp开启指定端口
func openPort(port int) (string, error) {
	prt := uint16(port)
	name := "Server" + randomString(5, false)
	d, err := upnp.Discover()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	err = d.Forward(prt, name)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return "Open port successed.", nil
}
//------------------------------------------------------------
// 修改Host文件
func editHost(data string, redution bool)(string, error) {
	if !isAdmin{
		FKDebugLog("You are not a Admin, so edit host file failed.")
		return "", errors.New("Edit host file need a admin account.")
	}

	if redution {
		// 恢复原本的Host文件
		if checkIsFileExist(winDirPath + hostFilePath + "hosts.bak") {
			err := deleteFile(winDirPath + hostFilePath + "hosts")
			if err != nil{
				return "", err
			}
			err = renameFile(winDirPath+hostFilePath+"hosts.bak", "hosts")
			if err != nil{
				return "", err
			}
		} else{
			return "", errors.New("Can't find hosts.bak.")
		}
	} else {
		// 备份并修改Host文件
		if !checkIsFileExist(winDirPath + hostFilePath + "hosts.bak") {
			err := renameFile(winDirPath+hostFilePath+"hosts", "hosts.bak")
			if err != nil{
				return "", err
			}
			err = createFileAndWriteData(winDirPath+hostFilePath+"hosts", []byte(data))
			if err != nil{
				return "", err
			}
		} else{
			return "", errors.New("Already has hosts.bak, maybe host file already changed")
		}
	}

	_, err := runThirdExe("ipconfig //flushdns")
	if err != nil{
		return "", err
	}
	return "Edit host successed.", nil
}
//------------------------------------------------------------