package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"text/template"
)

type RemoveUser struct {
	User       string `json:"user"`
	Pass       string `json:"pass"`
	RemoveUser string `json:"remove_user"`
}

func NewCmdRemoveUser(args ...string) *RemoveUser {
	return &RemoveUser{
		args[0],
		args[1],
		args[2],
	}
}

func (e *RemoveUser) String() string {
	jsonObj := struct {
		CmdString string `json:"cmd_string"`
		RemoveUser
	}{
		"delete user",
		*e,
	}
	b, _ := json.MarshalIndent(jsonObj, "", " ")
	return string(b)
}
func (e *RemoveUser) CheckArgs() error {
	if e.User == "" {
		return fmt.Errorf("cmd_remove_user user is nil")
	}
	if e.Pass == "" {
		return fmt.Errorf("cmd_remove_user pass is nil")
	}
	if e.RemoveUser == "" {
		return fmt.Errorf("cmd_remove_user remove_user is nil")
	}
	return nil
}

func (e *RemoveUser) Before() error {
	//DebugLog("RemoveUser Before\n")
	return nil
}
func (e *RemoveUser) After() error {
	//DebugLog("RemoveUser After\n")
	return nil
}

func (e *RemoveUser) Execute() (CmdResult, error) {
	//DebugLog("RemoveUser Execute\n")
	createUserTemplate := "delete from mysql.user where User='{{.user}}'"

	t, err := template.New("nopass").Parse(createUserTemplate)
	if err != nil {
		return nil, err
	}
	var byteBuffer bytes.Buffer
	mapValue := make(map[string]string)
	mapValue["user"] = e.RemoveUser
	t.Execute(&byteBuffer, mapValue)
	exeString := byteBuffer.String()
	return NewMysqlService().Execute(e.User, e.Pass, exeString)
}
