package dao

import (
	"testing"
	"time"
)

var (
	cmd = Command{
		Time: time.Now(),
		Code: CMD_RUN_EXE,
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

func TestCommand_String(t *testing.T) {
	t.Log(cmd.String())
}
func TestCommand_StandardCmd(t *testing.T) {
	t.Log(cmd.StandardCmd())
}
