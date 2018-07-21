package connect

import (
	"FKTrojan/common"
	"testing"
)

func TestGenerateCFBK(t *testing.T) {
	s := common.Obfuscate("In the beginning God created the heavens and the earth.")
	t.Log(common.Obfuscate(common.Base64Encode(s)))
}
func TestGetCFBK(t *testing.T) {
	t.Log(GetCFBK())
}
