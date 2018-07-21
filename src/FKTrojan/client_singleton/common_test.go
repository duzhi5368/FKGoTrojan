package client_singleton

import "testing"

func Test_PipeListen(t *testing.T) {
	t.Log(pipeListen(`\\.\pipe\tmp.sock`))
}
