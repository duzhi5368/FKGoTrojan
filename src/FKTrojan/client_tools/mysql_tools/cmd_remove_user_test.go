package main

import "testing"

func TestRemoveUser_Execute(t *testing.T) {
	createUser := NewCmdCreateUser("testRemove", "testRemove")
	err := createUser.Before()
	if err != nil {
		t.Error(err)
		return
	}
	_, err = createUser.Execute()
	if err != nil {
		t.Error(err)
		return
	}
	err = createUser.After()
	if err != nil {
		t.Error(err)
		return
	}
	ret, err := NewCmdRunSql("testRemove", "testRemove", "show databases").Execute()
	t.Log(ret, err)

	removeUser := NewCmdRemoveUser("testRemove", "testRemove", "testRemove")

	removeUser.Execute()
}
