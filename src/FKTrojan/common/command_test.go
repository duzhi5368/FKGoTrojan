package common

import (
	"testing"
	"time"
)

var (
	cmd = Command{
		Time: time.Now(),
		Code: "1x1",
		UID:  "000",
		Parameters: []string{
			"mysql_tools.exe",
			"sql",
			"-u",
			"user",
			"-p",
			"pass",
			"-s",
			"\"show processlist;\"",
		},
	}
)

func TestCommand_Encrypt(t *testing.T) {

	en := cmd.Encrypt()

	cmd_after, err := DecryptCommand(en)
	if err != nil {
		t.Error(err)
	}
	if cmd.Encrypt() != cmd_after.Encrypt() {
		t.Error(cmd)
		t.Error(cmd_after)
	}

	t.Log(cmd_after)
}

func TestCommand_String(t *testing.T) {
	t.Log(cmd.String())
}
func TestParseCommand(t *testing.T) {
	c, err := ParseCommand("Wednesday, 07-Mar-18 19:24:24 CST||000|1x1|mysql_tools.exe|sql|-u|user|-p|pass|-s|\"show processlist;\"")
	if err != nil {
		t.Error(err)
	}
	t.Log(c.String())
}
