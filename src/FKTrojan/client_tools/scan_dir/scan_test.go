package main

import (
	. "FKTrojan/common"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"testing"
	"time"
)

func init() {
	fmt.Println("PLEASE NOTE : this test may last for a very long time")
}
func TestScan(t *testing.T) {
	f := func(dir string, depth uint) ([]Item, error) {
		begin := time.Now()
		defer func() {
			cost := time.Since(begin).Seconds()
			t.Logf("scan cost %.3f", cost)
		}()
		d, err := Scan(dir, depth)
		if err != nil {
			t.Error(err)
		}
		t.Logf("dir %s has %d items\n", dir, len(d))
		//t.Log(d[:])
		return d, err
	}
	f("c:/", 1)

	d, err := f("d:/", 1)
	if err == nil {
		for _, i := range d {
			if i.ItemType == DIR_ITEM {
				dirfiles, _ := f(i.FullPath, DEPTH_ALL)
				files, _ := getSubItemByFind(i.FullPath, DEPTH_ALL)
				if len(dirfiles) != len(files)-1 {
					// scan_test.go:41: dir(d:/System Volume Information) len(0) != (4)
					// 测试发现，上面的无法通过，find命令可以获取到，scan函数不能

					t.Errorf("dir(%s) len(%d) != (%d)\n",
						i.FullPath, len(dirfiles), len(files)-1)
					fmt.Printf("dir(%s) Failed\n", i.FullPath)
				} else {
					fmt.Printf("dir(%s) OK\n", i.FullPath)
				}
			}
		}
	}
}

func TestGetSubItemByFind(t *testing.T) {
	f, err := getSubItemByFind("d:/DNSModifier - 副本", 2)
	t.Log(err, f)
}
func TestShellRunCmd(t *testing.T) {
	r, err := shellRunCmd("find /d/git/")
	//r, err := shellRunCmd("pwd")
	//r, err := shellRunCmd("date")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(r)
}

// 以下测试代码，使用的是shell下find命令，安装了git，就会有sh.exe
func shellRunCmd(cmdString string) ([]string, error) {
	//fmt.Println(cmdString)
	shExeFullPath := "D:\\Program Files\\Git\\bin\\sh.exe"
	if !Exist(shExeFullPath) {
		fmt.Printf("%s not exist, please find and reset it\n", shExeFullPath)
		os.Exit(1)
	}
	tmpFile := os.Getenv("temp") + "\\find.sh"
	//fmt.Println(tmpFile)

	d1 := []byte(cmdString)
	err := ioutil.WriteFile(tmpFile, d1, 0644)
	cmd := exec.Command(shExeFullPath, "/tmp/find.sh")
	out, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	defer cmd.Wait()
	scanner := bufio.NewScanner(out)
	ret := make([]string, 0)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return ret, nil
}
func dirToUnix(dir string) string {
	// 路径中的空格(）需要转义才能在unix环境中使用
	var re = regexp.MustCompile(`([\(\) ])`)
	dir = re.ReplaceAllString(dir, `\$1`)

	return "/" + dir[0:1] + dir[2:]
}
func getSubItemByFind(dir string, depth uint) (dirfiles []string, err error) {

	commonCmd := fmt.Sprintf("find %s ", dirToUnix(dir))
	if depth != DEPTH_ALL {
		commonCmd += fmt.Sprintf(" -maxdepth %d ", depth-1)
	}
	Cmd := commonCmd //+ " -type d"

	dirfiles, err = shellRunCmd(Cmd)

	return
}
