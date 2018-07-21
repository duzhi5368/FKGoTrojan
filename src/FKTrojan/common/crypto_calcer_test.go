package common

import "testing"

func TestDeobfuscate(t *testing.T) {
	t.Log(Deobfuscate("f24bdd7c.b738.53:b.c5:d.7ef2467b41:1"))
}

func TestObfuscate(t *testing.T) {
	t.Log(Base64Decode(Deobfuscate("fzKl[XOzfYC1Y4CieHhjPjKEPmydW3mv[H:4d2ydTX6{eHGtcHWzYGy8NUJyRUmDOkJuSENxND11NEOEMVJ5NkRuRlZ{RlSCPE[DN{Z6gWydZ3:vbH:{eD6mfHVjMDKmcnOzfYC1Y4CieHhjPjKEPmydW3mv[H:4d2ydTX6{eHGtcHWzYGy8NUJyRUmDOkJuSENxND11NEOEMVJ5NkRuRlZ{RlSCPE[DN{Z6gWydNER2Z{N6OkB2NXJ1NnN4OEh4O3SkO3V2O{B5N{h1OnFjgR>>")))
}
