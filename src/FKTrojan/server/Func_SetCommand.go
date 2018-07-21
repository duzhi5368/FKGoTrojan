/*
Author: FreeKnight
*/
//------------------------------------------------------------
package server
//------------------------------------------------------------
import (
	"fmt"
	"time"
)
//------------------------------------------------------------
// 向DB刷新Command
func setCommand(cmd string) bool {
	var tmpcmd string

	_ = DBPointer.QueryRow("SELECT command FROM command WHERE id=1").Scan(&tmpcmd)
	if tmpcmd == "" {
		_, err := DBPointer.Exec("INSERT INTO command(id, command, timeanddate) VALUES(?, ?, ?)", 1, "none", "never")
		if err != nil {
			fmt.Println(err)
			return false
		}
	}

	_, err := DBPointer.Exec("UPDATE command SET command='" + cmd + "', timeanddate='" + time.Now().Format(time.RFC850) + "' WHERE id=1")
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
//------------------------------------------------------------