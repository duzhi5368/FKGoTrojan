/*
Author: FreeKnight
单例检查
*/
//------------------------------------------------------------
package components

//------------------------------------------------------------
import (
	"FKTrojan/common"
	"os"
	"syscall"
	"unsafe"

	"github.com/StackExchange/wmi"
	//"strconv"
)

//------------------------------------------------------------
// 尝试创建同步锁
func createMutex(name string) (uintptr, error) {
	ret, _, err := procCreateMutex.Call(
		0,
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name))),
	)
	switch int(err.(syscall.Errno)) {
	case 0:
		return ret, nil
	default:
		return ret, err
	}
}

//------------------------------------------------------------
// 扫描进程检查是否有另外一个自身在执行
func checkForAnotherMe() (bool, string, string) {
	var dst []Win32_Process
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		return false, "", ""
	}
	for _, v := range dst {
		if string(common.Md5HashFile(*v.ExecutablePath)) == string(common.Md5HashFile(os.Args[0])) {
			if *v.ExecutablePath != os.Args[0] {
				return true, v.Name, *v.ExecutablePath
			}
		}
	}
	return false, "", ""
}

//------------------------------------------------------------
// 检查是否是单例
func checkIsSingleInstance(name string) bool {
	_, err := createMutex(name)
	if err != nil {
		return false
	}

	//alreadyHave, proname, proexepath := checkForAnotherMe()
	//if alreadyHave {
	//	FKDebugLog("Process name : " + proname + " Exe path: " + proexepath)
	//	return false
	//}

	return true
}

//------------------------------------------------------------
