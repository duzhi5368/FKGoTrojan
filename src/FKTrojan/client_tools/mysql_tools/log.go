package main

import "fmt"

var (
	g_debug bool = false
)

func DebugLog(fmtString string, a ...interface{}) {
	if !g_debug {
		return
	}
	fmt.Printf(fmtString, a...)
}
