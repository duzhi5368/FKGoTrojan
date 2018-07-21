package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
	"time"
)

type MysqlService struct{}

func NewMysqlService() *MysqlService {
	return &MysqlService{}
}
func (m *MysqlService) findMysqlService() (*ServiceInfo, error) {
	list, err := ListServices()
	if err != nil {
		return nil, err
	}
	for _, service := range list {
		if strings.Contains(strings.ToUpper(service.Name), "MYSQL") {
			e, _ := service.GetRunCmdExe()
			if strings.Contains(strings.ToUpper(e), "MYSQLD.EXE") {
				return &service, nil
			}
		}
	}
	return nil, fmt.Errorf("not found mysqld")
}
func (m *MysqlService) SkipGrantTableStart() error {
	service, err := m.findMysqlService()
	if err != nil {
		return err
	}
	service.Stop()
	time.Sleep(time.Millisecond * 1000)
	service.Kill()
	runCmd := strings.Replace(service.RunCmd, " "+service.Name, " --skip-grant-tables", 1)
	runCmd = "start /b \"mysql\" " + runCmd
	//fmt.Println(runCmd)
	tmpFile := os.Getenv("temp") + "\\mysql.bat"
	//fmt.Println(tmpFile)

	d1 := []byte(runCmd)
	err = ioutil.WriteFile(tmpFile, d1, 0644)
	_, stdErr, err := ExecuteWindowsCmd(tmpFile)
	DebugLog(fmt.Sprintf("%+v", stdErr))
	DebugLog(fmt.Sprintf("%+v", runCmd))
	os.Remove(tmpFile)
	return err
}

func (m *MysqlService) ResumeMysqlService() error {
	service, err := m.findMysqlService()
	if err != nil {
		return err
	}
	err = service.Kill()
	if err != nil {
		return err
	}
	err = service.Start()
	return err
}
func (m *MysqlService) Execute(user, pass, sql string) (CmdResult, error) {
	nopassTemplate := `"{{.bin}}" -u{{.user}} -e "{{.sql}};"`
	passTemplate := `"{{.bin}}" -u{{.user}} -p{{.pass}} -e "{{.sql}};"`
	service, err := m.findMysqlService()
	if err != nil {
		return nil, err
	}
	mysqldExe, err := service.GetRunCmdExe()
	if err != nil {
		return nil, err
	}
	mysqlExe := strings.Replace(mysqldExe, "mysqld.exe", "mysql.exe", -1)
	exeString := ""
	mapValue := make(map[string]string)
	mapValue["bin"] = mysqlExe
	mapValue["user"] = user
	mapValue["sql"] = sql
	var t *template.Template
	if pass != "" {
		t, err = template.New("pass").Parse(passTemplate)
		if err != nil {
			return nil, err
		}
		mapValue["pass"] = pass
	} else {
		t, err = template.New("nopass").Parse(nopassTemplate)
		if err != nil {
			return nil, err
		}
	}
	var byteBuffer bytes.Buffer
	err = t.Execute(&byteBuffer, mapValue)
	if err != nil {
		return nil, err
	}
	exeString = byteBuffer.String()
	resultFile := os.Getenv("temp") + "\\result.txt"
	exeString = "@echo off \r\n" + exeString + " > " + resultFile
	DebugLog(exeString)
	tmpFile := os.Getenv("temp") + "\\mysql.bat"
	DebugLog(tmpFile)

	d1 := []byte(exeString)
	err = ioutil.WriteFile(tmpFile, d1, 0644)
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile)
	defer os.Remove(resultFile)
	stdOut, stdErr, err := ExecuteWindowsCmd(tmpFile)
	if err != nil {
		return nil, err
	}
	DebugLog(fmt.Sprintf("mysql_execute stdout : %+v\n", stdOut))
	DebugLog(fmt.Sprintf("mysql execute stderr : %+v\n", stdErr))
	file, err := os.Open(resultFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var headers []string
	ret := make([]map[string]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(headers) == 0 {
			headers = strings.Split(line, "\t")
		} else {
			values := strings.Split(line, "\t")
			record := make(map[string]string)
			for i, v := range values {
				record[headers[i]] = v
			}
			ret = append(ret, record)
		}
	}
	ret = append(ret, map[string]string{
		"stdout": fmt.Sprintf("%+v", stdOut),
		"stderr": fmt.Sprintf("%+v", stdErr),
	})
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return ret, nil
}
