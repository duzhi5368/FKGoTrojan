package registry_crypto

import "testing"

func TestSet(t *testing.T) {
	t.Log(Set(UIDKEY, "123-345"))
}

func TestGet(t *testing.T) {
	t.Log(Get(UIDKEY))
}

func TestDel(t *testing.T) {
	t.Log(Del(UIDKEY))
}

func TestList(t *testing.T) {
	l, e := List()
	t.Logf("%+v,%v", l, e)
}
