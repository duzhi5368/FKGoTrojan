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
// 输出IP
func IpHandler(response http.ResponseWriter, request *http.Request) {
	if isEnabled {
		ip := strings.Split(request.RemoteAddr, ":")[0]
		fmt.Fprintf(response, ip)
	}
}
//------------------------------------------------------------