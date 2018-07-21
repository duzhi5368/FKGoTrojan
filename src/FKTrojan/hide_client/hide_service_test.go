package hide_client

import (
	"testing"
)

func TestCopyBinary(t *testing.T) {
	t.Log(copyBinary(&ServiceInfos[0]))
}

func TestWriteRegistry(t *testing.T) {
	t.Log(writeRegistry(&ServiceInfos[1]))
}

func TestReadRegistry(t *testing.T) {
	t.Log(readRegistry())
}

func TestInstall(t *testing.T) {
	t.Log(Install())
}

func TestUnInstall(t *testing.T) {
	t.Log(Uninstall())
}
