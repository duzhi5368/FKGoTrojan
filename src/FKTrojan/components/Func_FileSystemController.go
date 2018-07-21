/*
Author: FreeKnight
便捷的文件管理类
 */
//------------------------------------------------------------
package components
//------------------------------------------------------------
import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"os/exec"
	"syscall"
	"time"
	"strconv"
	"log"
	"io/ioutil"
)
//------------------------------------------------------------
// 检查一个文件是否存在
func checkIsFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
//------------------------------------------------------------
// 创建一个文件夹
func createDir(dirPath string, fileMode os.FileMode) bool {
	err := os.MkdirAll(dirPath, fileMode)
	if err != nil {
		return false
	}
	return true
}
//------------------------------------------------------------
// 重命名一个文件
func renameFile(pathFile string, name string) error {
	err := os.Rename(pathFile, name)
	if err != nil {
		return err
	}
	return nil
}
//------------------------------------------------------------
// 创建一个文件
func createFile(pathFile string) error {
	file, err := os.Create(pathFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
//------------------------------------------------------------
// 创建文件并写入
func createFileAndWriteData(fileName string, writeData []byte) error {
	fileHandle, err := os.Create(fileName)

	if err != nil {
		return err
	}
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()
	writer.Write(writeData)
	writer.Flush()
	return nil
}
//------------------------------------------------------------
// 拷贝文件到指定目录
func copyFileToDirectory(pathSourceFile string, pathDestFile string) error {
	FKDebugLog("Copy " + pathSourceFile + " to " + pathDestFile)
	sourceFile, err := os.Open(pathSourceFile)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	if checkIsFileExist(pathDestFile) {
		deleteFile(pathDestFile)
	}
	destFile, err := os.Create(pathDestFile)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	err = destFile.Sync()
	if err != nil {
		return err
	}

	sourceFileInfo, err := sourceFile.Stat()
	if err != nil {
		return err
	}

	destFileInfo, err := destFile.Stat()
	if err != nil {
		return err
	}

	if sourceFileInfo.Size() == destFileInfo.Size() {
	} else {
		err = errors.New("Bad copy file")
		return err
	}
	FKDebugLog("Copy successed!")
	return nil
}
//------------------------------------------------------------
// 删除一个文件
func deleteFile(nameFile string) error {
	err := os.Remove(nameFile)
	return err
}
//------------------------------------------------------------
// 删除一个文件夹（包括其中的所有文件）
func removeDirWithContent(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	err = os.RemoveAll(dir)
	if err != nil {
		return err
	}
	return nil
}
//------------------------------------------------------------
// 隐藏一个文件
func hideFile(file string) {
	runThirdExe("attrib +S +H " + file) //attrib -s -h -r /s /d
}
//------------------------------------------------------------
// 获取文件修改时间 返回unix时间戳
func getFileModifyTime(path string) int64 {
	f, err := os.Open(path)
	if err != nil {
		log.Println("open file error")
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Println("stat fileinfo error")
		return time.Now().Unix()
	}

	return fi.ModTime().Unix()
}
//------------------------------------------------------------
func modifyFile(filepath, content string, isAddToTail bool)(string, error){
	if isAddToTail{
		f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return "", err
		}

		defer f.Close()

		if _, err = f.WriteString(content); err != nil {
			return "", err
		}
	} else{
		f, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return "", err
		}

		defer f.Close()

		contentSlice := []byte(content)
		err = ioutil.WriteFile(filepath, contentSlice, 0644)
		if err != nil {
			return "", err
		}
	}

	dat, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}
//------------------------------------------------------------
// 安全修改一个文件
func safeModifyFile(filePath, content, modifyType string)(string, error){
	if !checkIsFileExist(filePath) {
		return "", errors.New("File doesn't exist.")
	}
	// 解析修改类型
	isAddToTail := strings.Contains(modifyType, "A")	// 是否在尾部添加。否则全文替换，是则添加于尾部
	isModifyDate := strings.Contains(modifyType, "D")	// 是否保持本文件修改时间不变（用以伪装）
	isAntiSafeDog := strings.Contains(modifyType, "S")	// 是否针对安全狗做额外行为

	// 麻醉狗
	cmd := exec.Command("winanvir.exe", "start")
	if isAntiSafeDog{
		// 执行新进程
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		err := cmd.Start()
		if err != nil {
			return "", err
		}

		// 最小化自身窗口
		CMinSizeOfAnvir();

		// 执行自动化填充
		time.Sleep(2 * time.Second)
		CAutoDoAnvir_SetEditText()
		time.Sleep(2 * time.Second)
		CAutoDoAnvir_SelectFristListItem()
		time.Sleep(2 * time.Second)
		CAutoDoAnvir_ClickBlock()
		time.Sleep(2 * time.Second)
	}

	var lModifyTime int64
	lModifyTime = 0
	if isModifyDate{
		lModifyTime := getFileModifyTime(filePath)
		s := strconv.FormatInt(lModifyTime, 10)
		log.Println(s)
	}

	// 修改文件内容
	info, err := modifyFile(filePath, content, isAddToTail)

	// 修改文件最后修改时间
	if isModifyDate && (lModifyTime != 0){
		time.Sleep(1 * time.Second)
		CWriteFileModifyTime(filePath,lModifyTime)
		time.Sleep(1 * time.Second)
	}

	// 唤醒狗
	if isAntiSafeDog {
		CAutoDoAnvir_ClickBlock();
		time.Sleep(2 * time.Second)

		// 杀掉Anvir进程
		done := make(chan error, 1)
		go func() {
			done <- cmd.Wait()
		}()
		select {
		case <-time.After(12 * time.Second):
			if err := cmd.Process.Kill(); err != nil {
				log.Fatal("failed to kill: ", err)
			}
			log.Println("process killed as timeout reached")
		case err := <-done:
			if err != nil {
				log.Printf("process done with error = %v", err)
			} else {
				log.Print("process done gracefully without error")
			}
		}

		// 强杀
		killThirdExe("winanvir.exe")
	}

	return info, err
}