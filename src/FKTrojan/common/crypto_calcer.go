/*
Author: FreeKnight
加解密
*/
//------------------------------------------------------------
package common

//------------------------------------------------------------
import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"io"
	"os"
)

//------------------------------------------------------------
// MD5
func Md5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

//------------------------------------------------------------
// MD5 hash一个文件
func Md5HashFile(filePath string) []byte {
	var result []byte
	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return nil
	}
	return hash.Sum(result)
}

// MD5 hash一个文件
func Md5HashStringFile(filePath string) string {
	var result []byte
	file, err := os.Open(filePath)
	if err != nil {
		return ""
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return ""
	}
	return hex.EncodeToString(hash.Sum(result))
}

//------------------------------------------------------------
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

//------------------------------------------------------------
func Base64Decode(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(data)
}

//------------------------------------------------------------
// 反模糊
func Deobfuscate(Data string) string {
	var ClearText string
	for i := 0; i < len(Data); i++ {
		ClearText += string(int(Data[i]) - 1)
	}
	return ClearText
}

//------------------------------------------------------------
// 快速的模糊处理
func Obfuscate(Data string) string {
	var ObfuscateText string
	for i := 0; i < len(Data); i++ {
		ObfuscateText += string(int(Data[i]) + 1)
	}
	return ObfuscateText
}

//------------------------------------------------------------
