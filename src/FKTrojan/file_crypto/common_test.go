package file_crypto

import (
	. "FKTrojan/common"
	"crypto/rand"
	"fmt"
	"io"
	"testing"
)

func TestGetCFBK(t *testing.T) {
	t.Log(Obfuscate(Base64Encode("Now the earth was formless and empty, darkness was over the surface of the deep, and the Spirit of God was hovering over the waters")))
}

func TestGetIV(t *testing.T) {
	t.Log(Obfuscate(Base64Encode("And God said, \"Let there be light,\" and there was light.")))
}

func newUUID() string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return ""
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("{%X-%X-%X-%X-%X}", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}
func TestGenerateTrustDir(t *testing.T) {
	i := 10
	for i > 0 {
		t.Log("\"C:/Windows/Installer/" + string(newUUID()) + "\",")
		i--
	}
}
