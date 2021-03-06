package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var (
	configPath = filepath.Join(os.Getenv("temp"), "config.json")
)

func init() {
	var config Config
	config.MySQLUser = "root"
	config.MySQLPass = "pass"
	config.MySQLName = "name"
	config.MySQLHost = "host"
	bb, err := json.MarshalIndent(config, " ", "  ")
	if err != nil {
		fmt.Println(err)
	}
	ioutil.WriteFile(configPath, bb, 0644)
}
func TestLoadConfig(t *testing.T) {
	loadConfig(configPath)
}
