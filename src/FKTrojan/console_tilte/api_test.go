package console_tilte

import (
	"FKTrojan/common"
	"testing"
)

func Test_NameDeo(t *testing.T) {
	t.Log(common.Obfuscate("GetConsoleTitleW"))
	t.Log(common.Obfuscate("SetConsoleTitleW"))
}

func TestGetConsoleTitle(t *testing.T) {
	t.Log(Get())
}

func TestSetConsoleTitle(t *testing.T) {
	t.Log(Set("搞点中文test"))
	t.Log(Get())
}
