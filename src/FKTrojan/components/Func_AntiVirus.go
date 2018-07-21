/*
Author: FreeKnight
安全保护 - 反病毒检查程序
*/
//------------------------------------------------------------
package components
//------------------------------------------------------------
import (
	"time"
)
//------------------------------------------------------------
var fkMagicNumber int64 = 0
//------------------------------------------------------------
// 绕过反病毒检查
func bypassAntiVirus() {
	// 无聊的人生要开始了
	if antiVirusBypass{
		allocateUnusefulMemory()
		jump()
	}
}
//------------------------------------------------------------
// 分配无聊的内存
func allocateUnusefulMemory() {
	for i := 0; i < 1000; i++ {
		var Size int = 30000000
		unusefulBuffer1 := make([]byte, Size)
		unusefulBuffer1[0] = 1
		var unusefulBuffer2 [102400000]byte
		unusefulBuffer2[0] = 0
	}
}
//------------------------------------------------------------
// 无聊的跳转和内存修改
func jump() {
	fkMagicNumber++
	hop1()
}

func hop1() {
	fkMagicNumber++
	time.Sleep(time.Duration(randInt(100, 250)) * time.Nanosecond)
	hop2()
}
func hop2() {
	fkMagicNumber++
	time.Sleep(time.Duration(randInt(100, 250)) * time.Nanosecond)
	hop3()
}
func hop3() {
	fkMagicNumber++
	time.Sleep(time.Duration(randInt(100, 250)) * time.Nanosecond)
	hop4()
}
func hop4() {
	fkMagicNumber++
	time.Sleep(time.Duration(randInt(100, 250)) * time.Nanosecond)
	hop5()
}
func hop5() {
	fkMagicNumber++
	time.Sleep(time.Duration(randInt(100, 250)) * time.Nanosecond)
	hop6()
}
func hop6() {
	fkMagicNumber++
	time.Sleep(time.Duration(randInt(100, 250)) * time.Nanosecond)
	hop7()
}
func hop7() {
	fkMagicNumber++
	time.Sleep(time.Duration(randInt(100, 250)) * time.Nanosecond)
	hop8()
}
func hop8() {
	fkMagicNumber++
	time.Sleep(time.Duration(randInt(100, 250)) * time.Nanosecond)
	hop9()
}
func hop9() {
	fkMagicNumber++
	time.Sleep(time.Duration(randInt(100, 250)) * time.Nanosecond)
	hop10()
}
func hop10() {
	fkMagicNumber++
}
