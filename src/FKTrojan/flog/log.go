package Flog

import (
	"FKTrojan/common"
	"os"
	"path/filepath"

	"fmt"

	"github.com/qiniu/log"
)

var (
	Flog *log.Logger
)

func init() {
	curr := common.CurrentBinaryDir()
	logDir := filepath.Join(curr, "log")
	os.MkdirAll(logDir, 666)
	logFilePath := filepath.Join(logDir, "log.txt")
	//fmt.Println(logFilePath)
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open log file %s error %v", logFilePath, err)
		os.Exit(1)
	}
	Flog = log.New(logFile, "", log.Ldefault)

	if Flog == nil {
		fmt.Printf("log init error")
		os.Exit(1)
	}
}
