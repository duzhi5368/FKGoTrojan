package antivirus_blocker

import (
	"FKTrojan/common"
	"testing"
	"time"
)

func TestExecute(t *testing.T) {
	t.Log(Execute(func() (string, error) {
		stdout, _, err := common.RunExe("ipconfig/all")
		time.Sleep(time.Second * 2)
		return stdout, err
	}))
}
