package components

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var (
	configPath = filepath.Join(os.Getenv("temp"), "config_client.json")
)

func init() {
	var config Config
	config.ServerAddress = "http://localhost:7777/"

	bb, err := json.MarshalIndent(config, " ", "  ")
	if err != nil {
		fmt.Println(err)
	}
	ioutil.WriteFile(configPath, bb, 0644)
}
func TestLoadConfig(t *testing.T) {
	loadConfig(configPath)
}
