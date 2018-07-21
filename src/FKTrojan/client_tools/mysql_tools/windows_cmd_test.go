package main

import "testing"

func TestExecuteWindowsCmd(t *testing.T) {
	o, e, err := ExecuteWindowsCmd("sc query | findstr SERVICE_NAME")
	t.Log(o, e, err)
}
