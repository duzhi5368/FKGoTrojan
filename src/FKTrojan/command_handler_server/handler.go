package command_handler_server

import (
	"FKTrojan/common"
	"FKTrojan/dao"
	"FKTrojan/database"
	"FKTrojan/flog"
	"fmt"
)

func ServerDo(cmd *dao.Command, client *dao.Client) error {
	err := database.InsertClientIfNotExist(client)
	if err != nil {
		Flog.Flog.Printf("query client error %v", err)
	}
	standardCmd(cmd)
	return nil
}

func standardCmd(cmd *dao.Command) error {
	if cmd.Code != dao.CMD_RUN_EXE {
		return nil
	}
	file := cmd.Parameters[0]
	if !common.PathExist(file) {
		return fmt.Errorf("file %s not exist", file)
	}
	cmd.Md5sum = common.Md5HashStringFile(file)
	return nil
}
