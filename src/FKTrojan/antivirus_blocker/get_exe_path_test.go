package antivirus_blocker

import "testing"

func TestGetExePath(t *testing.T) {
	t.Log(getExePath("ingress.exe"))
	t.Log(getExePath("iexplore.exe"))
}

func TestRemoveExt(t *testing.T) {
	t.Log(removeExt("a.exe"))
	t.Log(removeExt("a"))
	t.Log(removeExt("abc.exe.exe"))
}
