package database

import (
	"testing"
)

func TestGetNewCmd(t *testing.T) {
	r, err := GetNewCmd("de1fcd20-6661-4304-b525-be283ab39ccd")
	t.Log(r, err)
	//r, err = dbPointer.Exec("update command set last_update = NOW()")
	//t.Log(r, err)
}
