package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/kardianos/service"
)

var (
	name        = flag.String("sn", "", "安装服务名称，默认空")
	displayName = flag.String("dsn", "", "显示服务名称,默认空")
	description = flag.String("des", "", "服务描述信息,默认空")
	command     = flag.String("c", "run", "执行的命令参数（run,install,uninstall),默认run")
	port        = flag.Int("p", 80, "内置API端口号，默认80")
)

func main() {
	flag.Parse()
	serviceConfig := &service.Config{
		Name:        *name,
		DisplayName: *displayName,
		Description: *description,
	}
	// 构建服务对象
	prog := &Program{}
	s, err := service.New(prog, serviceConfig)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	// 用于记录系统日志
	logger, err := s.Logger(nil)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	switch *command {
	case "run":
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
	case "install":
		err = s.Install()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("安装成功")
	case "uninstall":
		err = s.Uninstall()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("卸载成功")
	default:
		fmt.Println("命令参数有误")
	}
}

type Program struct{}

func (p *Program) Start(s service.Service) error {
	log.Println("开始服务")
	go p.run()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	log.Println("停止服务")
	return nil
}

func (p *Program) run() {
	// 此处编写具体的服务代码
}
