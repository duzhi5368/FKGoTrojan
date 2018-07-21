package config_client

import (
	"FKTrojan/common"
	"testing"
)

var (
	remoteConfigStr = `
{
   "server_ip": "128.14.143.30",
   "cmd_port":7778,
   "trans_port":7779
}

`
)

func TestGetConfig(t *testing.T) {
	t.Log(common.Obfuscate(common.Base64Encode(remoteConfigStr)))
}
