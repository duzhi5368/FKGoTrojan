package config_client

import "testing"

func TestWrite(t *testing.T) {
	t.Log(writeRegistry(&Config{
		ServerIp:  "11.11.11.11",
		CmdPort:   7776,
		TransPort: 8889,
	}))
}

func TestRead(t *testing.T) {
	t.Log(readRegistry())
}
