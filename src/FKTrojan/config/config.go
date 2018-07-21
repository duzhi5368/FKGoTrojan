package config

import (
	"FKTrojan/common"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	MySQLUser              string `json:"mysql_user"`
	MySQLPass              string `json:"mysql_pass"`
	MySQLHost              string `json:"mysql_host"`
	MySQLName              string `json:"mysql_name"`
	BaseDataDir            string `json:"base_data_dir"`
	EnableCompressTransfer bool   `json:"enable_compress_transfer"`
	CmdPort                int    `json:"cmd_port"`
	TransPort              int    `json:"trans_port"`
}

var (
	Conf Config
)

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
	if Conf.CmdPort == 0 || Conf.TransPort == 0 {
		return fmt.Errorf("port not set")
	}
	//fmt.Printf("use %s %+v\n", configPath, config)
	return nil
}

func Load() {
	currentPath := common.CurrentBinaryDir()
	configPath := filepath.Join(currentPath, "config.json")
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
		copy("D:/bin/config.json", configPath)
	}
	err := loadConfig(configPath)
	if err != nil {
		fmt.Printf("load config error %v\n", err)
		os.Exit(1)
	}
}
