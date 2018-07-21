package file_crypto

import "testing"

var (
	enc = FileInfo{
		DecryptPath: "d:/bin/command_tools.exe",
		EncryptPath: "d:/bin/encrypt/output.data",
	}
	dec = FileInfo{
		DecryptPath: "d:/bin/encrypt/output.txt",
		EncryptPath: "d:/bin/encrypt/output.data",
	}
)

func TestFileInfo_EncryptFile(t *testing.T) {
	t.Log(enc.encryptFile())
}
func TestFileInfo_DecryptFile(t *testing.T) {
	t.Log(dec.decryptFile())
}
func TestMd5sum(t *testing.T) {
	t.Log(enc.md5sum())
}
func TestFileInfo_Encrypt(t *testing.T) {
	t.Log(enc.Encrypt())
	t.Log(enc)
}
func TestFindFile(t *testing.T) {
	t.Log(FindFile("b9d070847e9f24d1b6afc083e5b0e29f"))
}
