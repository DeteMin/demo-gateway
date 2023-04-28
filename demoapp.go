package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/cihub/seelog"

	"demoapp/internal/server"
)

func main() {
	defer func() { seelog.Flush() }()

	err := server.StartGRPC()
	if err != nil {
		_ = seelog.Errorf("start gRPC error %v", err)
		return
	}

	err = server.StartGateWay()
	if err != nil {
		_ = seelog.Errorf("start gateway error %v", err)
		return
	}

	err = server.StartMetrics()
	if err != nil {
		_ = seelog.Errorf("start metrics error %v", err)
		return
	}

	seelog.Infof("start success.")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		seelog.Infof("get a signal %s try to kill server.", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			server.StopGateWay()
			server.StopGRPC()
			server.StopMetrics()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
