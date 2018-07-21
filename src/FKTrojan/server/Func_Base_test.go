package server

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAppendToFile(t *testing.T) {
	AppendToFile(filepath.Join(os.Getenv("temp"), "test.data"), "abc")
	AppendToFile(filepath.Join(os.Getenv("temp"), "test.data"), "123")
	AppendToFile(filepath.Join(os.Getenv("temp"), "test.data"), "\r\n4567")
}
