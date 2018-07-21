package client_singleton

import (
	"net"

	"github.com/Microsoft/go-winio"
)

func pipeListen(name string) (net.Listener, error) {
	return winio.ListenPipe(`\\.\pipe\`+name, nil)
}
