/*
Author: FreeKnight
*/
//------------------------------------------------------------
package server
//------------------------------------------------------------
import (
	"net/http"
	"fmt"
)
//------------------------------------------------------------
func UpdateHandler(response http.ResponseWriter, request *http.Request) {
	if isEnabled {
		request.ParseForm()
		fmt.Fprintf(response, "ok")
	}
}
//------------------------------------------------------------