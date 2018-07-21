package main

import (
	"testing"
)

func TestMysqlSevice_FindMysqlService(t *testing.T) {
	s, err := NewMysqlService().findMysqlService()
	t.Log(s, err)
}

func TestMysqlSevice_SkipGrantTableStart(t *testing.T) {
	err := NewMysqlService().SkipGrantTableStart()
	t.Log(err)
}

func TestMysqlSevice_ResumeMysqlService(t *testing.T) {
	err := NewMysqlService().ResumeMysqlService()
	t.Log(err)
}
