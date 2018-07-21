package stream_utils

import (
	"os"
	"path/filepath"
	"testing"
)

var (
	fileName  = filepath.Join(os.Getenv("temp"), "teststream.zip")
	readFile  = "d:/bin/output.txt"
	writeFile = "d:/bin/a/b/c/d/e/f/g/output2.txt"
)

func TestFileToStream(t *testing.T) {
	middleFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		t.Log(err)
		return
	}
	defer middleFile.Close()
	FileToStream(readFile, middleFile)
}
func TestStreamToFile(t *testing.T) {
	middleFile, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		t.Log(err)
		return
	}
	defer middleFile.Close()
	StreamToFile(writeFile, middleFile)
}
