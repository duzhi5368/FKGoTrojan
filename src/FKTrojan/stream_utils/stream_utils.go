package stream_utils

import (
	"FKTrojan/common"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/*
负责传输
*/

// 文件传输，首先gzip压缩
// 场景：
//     本系统大多数传输的文件是文本，经测试文本文件的gzip能够压缩至原文件的10%-20%
//     另外一个是golang的exe，压缩至30%-40%
// 总之：消耗部分cpu性能，换取网络传输速度

// 本地server client测试结果  文件 10MB的exe文件
/* useGzip == true:
     FileToStream cost 6.960
     StreamToFile cost 7.011
   useGzip == false
     FileToStream cost 0.339
     StreamToFile cost 0.385
明显 本地测试不用压缩要快很多，但在网络情况复杂，网速较慢的环境gzip仍是可选项，将其配置化，
是否使用压缩选项由服务器协商客户端决定，防止出现：服务器使用压缩传输，客户端不压缩，或者反之
*/
func FileToStream(fileName string, w io.Writer, useGzip bool) error {
	/*beginTime := time.Now()
	defer func() {
		Flog.Printf("FileToStream cost %.3f\n", time.Since(beginTime).Seconds())
	}()*/
	fileStream, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer fileStream.Close()
	if useGzip {
		g := gzip.NewWriter(w)
		if err != nil {
			return err
		}
		defer g.Close()
		_, err = io.Copy(g, fileStream)
	} else {
		_, err = io.Copy(w, fileStream)
	}

	return err
}
func StreamToFile(fileName string, r io.Reader, useGzip bool) error {
	/*beginTime := time.Now()
	defer func() {
		Flog.Printf("StreamToFile cost %.3f\n", time.Since(beginTime).Seconds())
	}()
	*/
	baseDir := filepath.Dir(fileName)
	err := os.MkdirAll(baseDir, 666)
	if err != nil {
		return err
	}
	// 对存在的文件不覆盖，防止出现服务端的文件被强制覆盖掉
	// 如果客户端需要覆盖某文件，可以结合标准命令实现
	if common.PathExist(fileName) {
		timeStr := time.Now().Format("20060102-15-04-05")
		newName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
		fileName = newName + timeStr + filepath.Ext(fileName)
	}
	fileStream, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer fileStream.Close()
	if useGzip {
		g, err := gzip.NewReader(r)
		if err != nil {
			return err
		}
		defer g.Close()
		_, err = io.Copy(fileStream, g)
	} else {
		_, err = io.Copy(fileStream, r)
	}
	return err
}
