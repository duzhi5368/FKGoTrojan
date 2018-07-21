/*
Author: FreeKnight
*/
//------------------------------------------------------------
package server
//------------------------------------------------------------
import (
	"io/ioutil"
	"fmt"
	"net/http"
	"strings"
	"os"
	"path/filepath"
	"path"
)
//------------------------------------------------------------
func ProfileFilesHandler(response http.ResponseWriter, request *http.Request) {
	if isEnabled {
		userName := GetUserName(request)
		if userName != "" {
			url := request.URL.Path
			var newurl = strings.Replace(url, "files", "", -1)
			var fixhtml2 string
			if newurl == "/favicon.ico" {
				// 忽略
			} else {
				urlPath := ProfileDir + newurl
				f, err := os.Open(urlPath)
				if err != nil {
					reportError(response, err)
					return
				}

				defer f.Close()
				fi, err := f.Stat()
				if err != nil {
					reportError(response, err)
					return
				}

				switch mode := fi.Mode(); {
				case mode.IsDir():
					files, err := ioutil.ReadDir(urlPath)
					if err != nil {
						reportError(response, err)
						return
					}

					last, err := filepath.Abs(path.Join(url, ".."))
					if err != nil {
						reportError(response, err)
						return
					}

					output := "<li>[Dir] <a href=\"" + last + "\">..</a>"
					for _, element := range files {
						output += "<li>"
						if element.IsDir() {
							output += "[Dir] "
						} else {
							output += "[File] "
						}
						output += "<a href=\"" + path.Join(url, element.Name()) + "\">" + element.Name() + "</a>"
					}
					var fixhtml = strings.Replace(FilebrowseHTML, "{STATS}", createcountDiv(), -1)
					var fixhtml1 = strings.Replace(fixhtml, "{FILES}", output, -1)
					if UseSSL {
						fixhtml2 = strings.Replace(fixhtml1, "{HOST}", "https://"+request.Host, -1)
					} else {
						fixhtml2 = strings.Replace(fixhtml1, "{HOST}", "http://"+request.Host, -1)
					}
					fmt.Fprintf(response, fixhtml2)
				case mode&os.ModeType == 0:
					content := "application/octet-stream"
					extension := filepath.Ext(urlPath)
					switch extension {
					case ".pdf":
						content = "application/pdf"
					case ".mp3":
						content = "audio/mp3"
					case ".jpg":
						content = "image/jpeg"
					case ".gif":
						content = "image/gif"
					case ".png":
						content = "image/png"
					case ".css":
						content = "text/css"
					case ".html":
						content = "text/html"
					case ".js":
						content = "text/javascript"
					case ".mp4":
						content = "video/mp4"
					case ".sh":
						content = "text/plain"
					case ".txt":
						content = "text/plain"
					case ".xml":
						content = "application/xml"
					}

					text, err := ioutil.ReadFile(urlPath)
					if err != nil {
						reportError(response, err)
						return
					}

					response.Header().Set("Content-Type", content)
					response.Write(text)
				}
			}
		} else {
			fmt.Fprintf(response, LoginHTML)
		}
	}
}
//------------------------------------------------------------