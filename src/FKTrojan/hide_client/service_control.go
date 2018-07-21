package hide_client

import (
	"fmt"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/eventlog"
	"golang.org/x/sys/windows/svc/mgr"
)

func (si *ServiceInfo) installService() error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()
	s, err := m.OpenService(si.Name)
	if err == nil {
		s.Close()
		return fmt.Errorf("service %s already exists", si.Name)
	}
	s, err = m.CreateService(si.Name, si.Path, mgr.Config{DisplayName: si.DisplayName, Description: si.Desc, StartType: mgr.StartAutomatic}, si.Args...)
	if err != nil {
		return err
	}
	defer s.Close()
	err = eventlog.InstallAsEventCreate(si.Name, eventlog.Error|eventlog.Warning|eventlog.Info)
	if err != nil {
		s.Delete()
		return fmt.Errorf("SetupEventLogSource() failed: %s", err)
	}
	return nil
}
func (si *ServiceInfo) startService() error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()
	s, err := m.OpenService(si.Name)
	if err != nil {
		return err
	}
	defer s.Close()
	return s.Start(si.Args...)
}
func (si *ServiceInfo) stopService() error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()
	s, err := m.OpenService(si.Name)
	if err != nil {
		return nil
	}
	defer s.Close()
	_, err = s.Control(svc.Stop)
	return err
}

func (si *ServiceInfo) status() (svc.State, error) {
	m, err := mgr.Connect()
	if err != nil {
		return svc.State(0), err
	}
	defer m.Disconnect()
	s, err := m.OpenService(si.Name)
	if err != nil {
		return svc.State(0), err
	}
	defer s.Close()
	status, err := s.Query()
	if err != nil {
		return svc.State(0), err
	}
	return status.State, nil
}
func (si *ServiceInfo) removeService() error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()
	s, err := m.OpenService(si.Name)
	if err != nil {
		return fmt.Errorf("service %s is not installed", si.Name)
	}
	defer s.Close()
	err = s.Delete()
	if err != nil {
		return err
	}
	err = eventlog.Remove(si.Name)
	if err != nil {
		return fmt.Errorf("RemoveEventLogSource() failed: %s", err)
	}
	return nil
}
