/*
Author: FreeKnight
*/
//------------------------------------------------------------
package server
//------------------------------------------------------------
import (
	"net/http"
	"strings"
	"fmt"
)
//------------------------------------------------------------
func IndexHandler(response http.ResponseWriter, request *http.Request) {
	if isEnabled {
		ip := strings.Split(request.RemoteAddr, ":")[0]
		FKLog("Index visited by " + ip)
		fmt.Fprintf(response, "404 page not found")
	}
}
//------------------------------------------------------------