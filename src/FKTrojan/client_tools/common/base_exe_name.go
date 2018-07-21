package common

import (
	"os"
	"path/filepath"
	"strings"
)

func ExeBaseName() string {
	baseName := filepath.Base(os.Args[0])
	return strings.TrimSuffix(baseName, filepath.Ext(baseName))
}
