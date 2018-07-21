package common

import (
	"fmt"
	"strings"
	"time"
)

type Command struct {
	Code       string
	CommandID  string
	UID        string
	Time       time.Time
	Parameters []string
}

var (
	SupportCMDCode = map[string]bool{
		"2x6": true,
	}
)

// 密文->command结构
func DecryptCommand(encryptoString string) (*Command, error) {
	time_command := strings.Split(encryptoString, "||")
	if len(time_command) < 2 {
		return nil, fmt.Errorf("%s is not a valid time_command", encryptoString)
	}
	then, err := time.Parse(time.RFC850, time_command[0])
	if err != nil {
		return nil, err
	}
	// 发过来的命令是先模糊再Base64的，于是这里要先解模糊再Base64解码
	decode := Base64Decode(Deobfuscate(time_command[1]))
	uid_command := strings.Split(decode, "|") // 客户端UID参数
	if len(uid_command) < 4 {
		return nil, fmt.Errorf("%s is not a valid uid_command", decode)
	}
	var cmd Command
	cmd.Time = then
	cmd.CommandID = uid_command[0]
	cmd.UID = uid_command[1]
	cmd.Code = uid_command[2]
	cmd.Parameters = uid_command[3:]
	return &cmd, nil
}

// command加密
func (c *Command) Encrypt() string {
	uid_code_para := make([]string, 0)
	uid_code_para = append(uid_code_para, c.CommandID)
	uid_code_para = append(uid_code_para, c.UID)
	uid_code_para = append(uid_code_para, c.Code)
	uid_code_para = append(uid_code_para, c.Parameters...)
	encrypto := fmt.Sprintf("%s||%s",
		c.Time.Format(time.RFC850),
		Obfuscate(Base64Encode(strings.Join(uid_code_para, "|"))))
	return encrypto
}

func (c *Command) String() string {
	uid_code_para := make([]string, 0)
	uid_code_para = append(uid_code_para, c.CommandID)
	uid_code_para = append(uid_code_para, c.UID)
	uid_code_para = append(uid_code_para, c.Code)
	uid_code_para = append(uid_code_para, c.Parameters...)
	str := fmt.Sprintf("%s||%s",
		c.Time.Format(time.RFC850),
		strings.Join(uid_code_para, "|"))
	return str
}
func (c *Command) Hash() string {
	return Md5Hash(c.String())
}

// 数据库明文-> command结构
func ParseCommand(commandString string) (*Command, error) {
	time_command := strings.Split(commandString, "||")
	if len(time_command) < 2 {
		return nil, fmt.Errorf("%s is not a valid time_command", commandString)
	}
	then, err := time.Parse(time.RFC850, time_command[0])
	if err != nil {
		return nil, err
	}
	uid_command := strings.Split(time_command[1], "|") // 客户端UID参数
	if len(uid_command) < 4 {
		return nil, fmt.Errorf("%s is not a valid uid_command", time_command[1])
	}
	var cmd Command
	cmd.Time = then
	cmd.CommandID = uid_command[0]
	cmd.UID = uid_command[1]
	cmd.Code = uid_command[2]
	cmd.Parameters = uid_command[3:]
	return &cmd, nil
}
