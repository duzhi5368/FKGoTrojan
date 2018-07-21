/*
Author: FreeKnight
反向代理功能
*/
//------------------------------------------------------------
package components
//------------------------------------------------------------
import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
)
//------------------------------------------------------------
const (
	DEFAULT_SERVER_TIMEOUT = 30	// 默认反向代理超时时间
)
//------------------------------------------------------------
type (
	BackendServer struct {
		Proxy *httputil.ReverseProxy
		Url   *url.URL
	}
)
//------------------------------------------------------------
var (
	port           string
	backends       string
	backendServers []*BackendServer
)
//------------------------------------------------------------
// 處理請求消息
func handle(w http.ResponseWriter, req *http.Request) {
	backendServer, err := getBackendServer()
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	FKDebugLog("Proxying request for " + req.URL.String() + " to backend server with address: " + backendServer.Url.String())
	// 啥都不幹，直接轉發給後端服務器
	backendServer.Proxy.ServeHTTP(w, req)
}
//------------------------------------------------------------
// 獲取隨機後端服務器
func getBackendServer() (*BackendServer, error) {
	if len(backendServers) == 0 {
		return nil, fmt.Errorf("No backend servers available :(")
	}

	return backendServers[rand.Intn(len(backendServers))], nil
}
//------------------------------------------------------------
// 解析backends	例如"127.0.0.1:8080"
func parseBackends() {
	splitBackends := strings.Split(backends, ",")

	for _, backend := range splitBackends {
		backend = strings.Trim(backend, " ")

		match, _ := regexp.MatchString("^(?:https?:)?//", backend)
		if match == false {
			backend = "http://" + backend
		}

		backendUrl, err := url.Parse(backend)
		if err != nil || len(backend) == 0 {
			continue
		}

		backendServer := &BackendServer{
			Proxy: httputil.NewSingleHostReverseProxy(backendUrl),
			Url:   backendUrl,
		}

		backendServers = append(backendServers, backendServer)
	}
}
//------------------------------------------------------------
// 启动代理服务器
func startProxServer(){
	mux := http.NewServeMux()

	server := &http.Server{}

	server.Addr = ":" + port
	server.Handler = mux
	server.ReadTimeout = time.Duration(DEFAULT_SERVER_TIMEOUT) * time.Second
	server.WriteTimeout = time.Duration(DEFAULT_SERVER_TIMEOUT) * time.Second

	mux.Handle("/", http.HandlerFunc(handle))	// 註冊處理器

	FKDebugLog("Proxy Server running on port " + port)

	go server.ListenAndServe()
}
//------------------------------------------------------------
// 開啟反向代理服務器
func startReverseProxy(myport, yourbackends string)(string, error) {
	// 先添加自身到防火墙
	addFileToFirewall(myExeName, os.Args[0])
	// 记录信息
	port = myport
	backends = yourbackends

	parseBackends()		// 解析backends
	startProxServer()	// 启动代理服务器

	return "Start reverse proxy successed.", nil
}
//------------------------------------------------------------