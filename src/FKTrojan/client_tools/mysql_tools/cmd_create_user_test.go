package main

import "testing"

func TestCreateUser_Execute(t *testing.T) {
	//createUser := NewCmdCreateUser("myuser1", "mypass1")
	createUser := NewCmdCreateUser("testuser", "123")
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
}
