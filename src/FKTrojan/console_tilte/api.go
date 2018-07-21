package console_tilte

import (
	"FKTrojan/common"
	"fmt"
	"syscall"
	"unsafe"
)

var (
	kernel32 = syscall.NewLazyDLL(common.Deobfuscate("lfsofm43/emm")) //kernel32.dll
	// 参见Test_NameDeo
	procGetConsoleTitle = kernel32.NewProc(common.Deobfuscate("HfuDpotpmfUjumfX")) //GetConsoleTitle
	procSetConsoleTitle = kernel32.NewProc(common.Deobfuscate("TfuDpotpmfUjumfX")) //SetConsoleTitle
)

func Get() (result string, err error) {

	_buf := make([]uint16, 256)
	_addr := uintptr(unsafe.Pointer(&_buf[0]))

	ret, _, callErr := procGetConsoleTitle.Call(_addr, 256)
	if ret == 0 {
		err = fmt.Errorf("GetConsoleTitle() error %v", callErr)
		return
	}
	_buf2 := syscall.UTF16ToString(_buf)
	result = _buf2
	return
}
func Set(title string) error {
	b, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		return err
	}
	ret, _, callErr := procSetConsoleTitle.Call(
		uintptr(unsafe.Pointer(b)),
	)
	if ret == 0 {
		return fmt.Errorf("callErr %v", callErr)
	}
	return nil
}
