package main

import (
	"encoding/json"
	"fmt"
)

type RunSql struct {
	User string `json:"user"`
	Pass string `json:"pass"`
	Sql  string `json:"sql"`
}

func NewCmdRunSql(args ...string) *RunSql {
	return &RunSql{
		args[0],
		args[1],
		args[2],
	}
}
func (e *RunSql) String() string {
	jsonObj := struct {
		CmdString string `json:"cmd_string"`
		RunSql
	}{
		"run_sql",
		*e,
	}
	b, _ := json.MarshalIndent(jsonObj, "", " ")
	return string(b)
}
func (e *RunSql) CheckArgs() error {
	DebugLog("RunSql CheckArgs\n")
	if e.User == "" {
		return fmt.Errorf("cmd_run_sql user is nil")
	}
	if e.Pass == "" {
		return fmt.Errorf("cmd_run_sql pass is nil")
	}
	if e.Sql == "" {
		return fmt.Errorf("cmd_run_sql sql is nil")
	}
	return nil
}

func (e *RunSql) Before() error {
	DebugLog("RunSql Before\n")
	return nil
}
func (e *RunSql) After() error {
	DebugLog("RunSql After\n")
	return nil
}

func (e *RunSql) Execute() (CmdResult, error) {
	DebugLog("RunSql Execute\n")
	return NewMysqlService().Execute(e.User, e.Pass, e.Sql)
}
