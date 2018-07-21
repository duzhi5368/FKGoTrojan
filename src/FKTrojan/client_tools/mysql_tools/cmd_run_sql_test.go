package main

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestRunSql_Execute(t *testing.T) {
	sqls := []string{
		"select * from panel.accounts",
		"use panel ; show create table accounts",
		"select * from mysql.user",
		"show databases",
	}
	for _, sql := range sqls {
		r := NewCmdRunSql("myuser1", "mypass1", sql)
		ret, err := r.Execute()
		if err != nil {
			t.Errorf("sql [%s] get error %v", sql, err)
		}
		json, _ := json.MarshalIndent(ret, " ", " ")
		t.Log(sql)
		t.Log(string(json))
	}
}

func TestSplitSkipNull(t *testing.T) {
	source := "a\tb\t\t"
	ret := strings.Split(source, "\t")
	t.Log(len(ret))
}

func TestRunSql_String(t *testing.T) {

	r := NewCmdRunSql("myuser1", "mypass1", "show databases")
	t.Log(r.String())
}
