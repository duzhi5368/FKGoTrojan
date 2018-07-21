package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"text/template"
)

type CreateUser struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

func NewCmdCreateUser(args ...string) *CreateUser {
	return &CreateUser{
		args[0],
		args[1],
	}
}
func (e *CreateUser) String() string {
	jsonObj := struct {
		CmdString string `json:"cmd_string"`
		CreateUser
	}{
		"create user",
		*e,
	}
	b, _ := json.MarshalIndent(jsonObj, "", " ")
	return string(b)
	return ""
}
func (e *CreateUser) CheckArgs() error {
	if e.User == "" {
		return fmt.Errorf("cmd_create_user user is nil")
	}
	if e.Pass == "" {
		return fmt.Errorf("cmd_create_user pass is nil")
	}
	return nil
}

func (e *CreateUser) Before() error {
	DebugLog("CreateUser Before\n")
	return NewMysqlService().SkipGrantTableStart()
}
func (e *CreateUser) After() error {
	DebugLog("CreateUser After\n")
	return NewMysqlService().ResumeMysqlService()
}

func (e *CreateUser) Execute() (CmdResult, error) {
	DebugLog("CreateUser Execute\n")
	createUserTemplate := "FLUSH PRIVILEGES;CREATE USER '{{.user}}'@'localhost' IDENTIFIED BY '{{.pass}}';GRANT ALL PRIVILEGES ON *.* TO '{{.user}}'@'localhost' WITH GRANT OPTION;FLUSH PRIVILEGES;"

	t, err := template.New("nopass").Parse(createUserTemplate)
	if err != nil {
		return nil, err
	}
	var byteBuffer bytes.Buffer
	mapValue := make(map[string]string)
	mapValue["user"] = e.User
	mapValue["pass"] = e.Pass
	t.Execute(&byteBuffer, mapValue)
	exeString := byteBuffer.String()
	//fmt.Println(exeString)
	return NewMysqlService().Execute("root", "", exeString)
}
