package common

import (
	"os"
	"testing"
)

func TestRunningCount(t *testing.T) {
	t.Log(RunningCount(os.Args[0]))
}

func TestKillPathExe(t *testing.T) {
	t.Log(KillPathExe("C:\\Windows\\system32\\sppmgr.exe"))
}
