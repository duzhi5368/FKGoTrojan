/*
Author: FreeKnight
*/
//------------------------------------------------------------
package server
//------------------------------------------------------------
import "net/http"
//------------------------------------------------------------
func LogoutHandler(response http.ResponseWriter, request *http.Request) {
	if isEnabled {
		ClearSession(response)
		http.Redirect(response, request, "/", 302)
	}
}
//------------------------------------------------------------