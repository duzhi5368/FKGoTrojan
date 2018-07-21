package main

import (
	"testing"
	"time"

	"fmt"

	"golang.org/x/sys/windows/svc"
)

func TestListServices(t *testing.T) {
	l, err := ListServices()
	t.Log(l, err)
}

type handleFunc func(*ServiceInfo) (bool, error)

func RunAllService(f handleFunc) (err error) {
	l, err := ListServices()
	if err != nil {
		return
	}
	for _, v := range l {
		isContinue, e := f(&v)
		if !isContinue {
			err = e
			break
		}
	}
	return
}
func TestServiceInfo_Stop(t *testing.T) {
	err := RunAllService(func(v *ServiceInfo) (isContinue bool, err error) {
		if v.Name == "MySQL57" {
			err = v.Stop()
			isContinue = false
			if err != nil {
				return
			}
			time.Sleep(time.Second)

			state, e := v.GetStatus()
			if err != nil {
				err = e
				return
			}
			if state != svc.Stopped && state != svc.StopPending {
				err = fmt.Errorf("status %d to right after Stop", state)
			}
		} else {
			isContinue = true
		}
		return
	})

	if err != nil {
		t.Error(err)
	}
}

func TestServiceInfo_Start(t *testing.T) {
	err := RunAllService(func(v *ServiceInfo) (isContinue bool, err error) {
		if v.Name == "MySQL57" {
			err = v.Start()
			isContinue = false
			if err != nil {
				t.Error(err)
				return
			}
			time.Sleep(time.Second)

			state, e := v.GetStatus()
			if err != nil {
				t.Error(e)
				err = e
				return
			}
			if state != svc.Running && state != svc.StartPending {
				t.Errorf("status %d to right after Start", state)
			}
		} else {
			isContinue = true
		}
		return
	})
	if err != nil {
		t.Error(err)
	}
}
func TestServiceInfo_Kill(t *testing.T) {
	err := RunAllService(func(v *ServiceInfo) (isContinue bool, err error) {
		if v.Name == "MySQL57" {
			err = v.Kill()
			isContinue = false
			if err != nil {
				return
			}
			time.Sleep(time.Second)

			state, e := v.GetStatus()
			if err != nil {
				err = e
				return
			}
			if state != svc.StopPending && state != svc.Stopped {
				err = fmt.Errorf("status %d to right after Kill", state)
			}
		} else {
			isContinue = true
		}
		return
	})
	if err != nil {
		t.Error(err)
	}
}

func TestServiceInfo_GetRunCmdExe(t *testing.T) {
	RunAllService(func(info *ServiceInfo) (bool, error) {
		t.Log(info.Name)
		n, err := info.GetRunCmdExe()
		if err != nil {
			t.Error(err)
		}
		t.Log(n)
		return true, nil
	})
}
