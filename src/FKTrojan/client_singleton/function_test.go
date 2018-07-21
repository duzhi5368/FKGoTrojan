package client_singleton

import "testing"

func TestMasterRegisterTwiceNormal(t *testing.T) {
	u, err := MasterRegister()
	if err != nil {
		t.Log(err)
		return
	}
	err = Unregister(u)
	if err != nil {
		t.Log(err)
		return
	}
	u, err = MasterRegister()
	if err != nil {
		t.Log(err)
		return
	}
	err = Unregister(u)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log("ok")
}
func TestMasterRegisterTwiceOpen(t *testing.T) {
	u, err := MasterRegister()
	if err != nil {
		t.Log(err)
		return
	}

	u, err = MasterRegister()
	if err != nil {
		t.Log(err)
		return
	}
	err = Unregister(u)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log("ok")
}
func TestMasterUnregister(t *testing.T) {
	err := Unregister(nil)
	t.Log(err)
}

func TestMasterIsRunning(t *testing.T) {
	t.Log(MasterIsRunning())
}

func TestSlaveIsRunning(t *testing.T) {
	t.Log(SlaveIsRunning())
}
