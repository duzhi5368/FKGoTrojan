/*
Author: FreeKnight
*/
//------------------------------------------------------------
package server
//------------------------------------------------------------
import (
	"net/http"
	"database/sql"
	"strings"
	"fmt"
)
//------------------------------------------------------------
func PurgeHandler(response http.ResponseWriter, request *http.Request) {
	userName := GetUserName(request)
	if userName != "" {
		request.ParseForm()
		GUID := request.Form.Get("guid")
		var tmpguid string

		err := DBPointer.QueryRow("SELECT guid FROM clients WHERE guid=?", GUID).Scan(&tmpguid)
		if err == sql.ErrNoRows {
			// 不存在……
			r := strings.NewReplacer("{STATS}", createcountDiv(), "{ERROR}", "No entry by GUID found")
			result := r.Replace(ErrorHTML)
			fmt.Fprintf(response, result)
		} else {
			err1 := DBPointer.QueryRow("DELETE FROM clients WHERE guid=?", GUID)
			if err1 != nil {
				r := strings.NewReplacer("{STATS}", createcountDiv(), "{ERROR}", "Database Error")
				result := r.Replace(ErrorHTML)
				fmt.Fprintf(response, result)
			} else {
				//Files will be kept
				r := strings.NewReplacer("{STATS}", createcountDiv(), "{MESSAGE}", "Client Purged from Database!")
				result := r.Replace(SuccessHTML)
				fmt.Fprintf(response, result)
			}
		}
	} else {
		fmt.Fprintf(response, LoginHTML)
	}
}
//------------------------------------------------------------