package main

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

type ServiceInfo struct {
	Name          string
	RunCmd        string
	CurrentStatus svc.State
}
type CmdControl int

const (
	CMD_START CmdControl = iota
	CMD_STOP
	CMD_KILL
)

func ListServices() ([]ServiceInfo, error) {
	serviceNames, _, err := ExecuteWindowsCmd("sc query state= all | findstr SERVICE_NAME")

	if err != nil {
		return nil, err
	}
	m, err := mgr.Connect()
	if err != nil {
		return nil, err
	}
	defer m.Disconnect()

	services := make([]ServiceInfo, 0)
	for _, serviceName := range serviceNames {
		name := strings.Replace(serviceName, "SERVICE_NAME: ", "", 1)
		name = strings.Trim(name, " ")
		func() {
			s, err := m.OpenService(name)
			if err != nil {
				return
			}
			defer s.Close()
			c, err := s.Config()
			if err != nil {
				return
			}
			q, err := s.Query()
			if err != nil {
				return
			}
			services = append(services, ServiceInfo{name, c.BinaryPathName, q.State})
		}()

	}
	return services, nil
}
func (s *ServiceInfo) control(name string, c CmdControl) error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()
	service, err := m.OpenService(name)
	if err != nil {
		return err
	}
	defer service.Close()
	switch c {
	case CMD_STOP:
		_, err = service.Control(svc.Stop)
		break
	case CMD_START:
		err = service.Start()
		break
	default:
		err = fmt.Errorf("ServiceInfo.Control(%d) unknown cmd", c)
		break
	}
	return err
}

func (s *ServiceInfo) Start() error {
	status, err := s.GetStatus()
	if err != nil {
		return err
	}
	if status == svc.Running || status == svc.StartPending {
		return nil
	}
	return s.control(s.Name, CMD_START)
}

func (s *ServiceInfo) Stop() error {

	status, err := s.GetStatus()
	if err != nil {
		return err
	}
	if status == svc.Stopped || status == svc.StopPending {
		return nil
	}
	return s.control(s.Name, CMD_STOP)
}
func (s *ServiceInfo) GetRunCmdExe() (string, error) {
	exeFullPath := ""
	// 执行exe路径中有空格，会将exe引起来
	if s.RunCmd[0:1] == "\"" {
		i := strings.Index(s.RunCmd[1:], "\"")
		exeFullPath = s.RunCmd[1 : i+1]
	} else {
		// 执行exe没有空格，不需要引号
		i := strings.Index(s.RunCmd[0:], " ")
		if i == -1 {
			// 无参数
			exeFullPath = s.RunCmd
		} else {
			// 有参数
			exeFullPath = s.RunCmd[0 : i+1]
		}
	}
	return exeFullPath, nil
}
func (s *ServiceInfo) Kill() error {
	exePath, err := s.GetRunCmdExe()
	if err != nil {
		return err
	}
	exeName := path.Base(filepath.ToSlash(exePath))
	_, _, err = ExecuteWindowsCmd(fmt.Sprintf("taskkill /im %s /f", exeName))
	return err
}

func (s *ServiceInfo) GetStatus() (svc.State, error) {
	m, err := mgr.Connect()
	if err != nil {
		return svc.Stopped, err
	}
	defer m.Disconnect()
	service, err := m.OpenService(s.Name)
	if err != nil {
		return svc.Stopped, err
	}
	defer service.Close()
	status, err := service.Query()
	if err != nil {
		return svc.Stopped, err
	}
	return status.State, nil
}
