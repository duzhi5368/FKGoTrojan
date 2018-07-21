/*
Author: FreeKnight
*/
//------------------------------------------------------------
package components
//------------------------------------------------------------
import (
	/*
		extern void AutoDoAnvir_SetEditText();
		extern void AutoDoAnvir_SelectFristListItem();
		extern void AutoDoAnvir_ClickBlock();
		extern void AutoDoAnvirWindowStep4();
		extern void MinSizeOfAnvir();
		extern void WriteFileModifyTime(char* filename, long long iModifyTime);
	*/
	"C"
	//"unsafe"
)
//------------------------------------------------------------
// 最小化Anvir窗口
func CMinSizeOfAnvir(){
	go C.MinSizeOfAnvir()
}
//------------------------------------------------------------
func CAutoDoAnvir_SetEditText(){
	go C.AutoDoAnvir_SetEditText()
}
//------------------------------------------------------------
func CAutoDoAnvir_SelectFristListItem(){
	go C.AutoDoAnvir_SelectFristListItem()
}
//------------------------------------------------------------
func CAutoDoAnvir_ClickBlock(){
	go C.AutoDoAnvir_ClickBlock()
}
//------------------------------------------------------------
func CWriteFileModifyTime(filename string, iModifyTime int64){
	cs := C.CString(filename)
	C.WriteFileModifyTime(cs, C.longlong(iModifyTime))
	//C.free(unsafe.Pointer(cs))
}
//------------------------------------------------------------