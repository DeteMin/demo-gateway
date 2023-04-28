package server

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/cihub/seelog"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"demoapp/internal/config"
)

var mS *http.Server

func StartMetrics() error {
	lis, err := net.Listen("tcp", config.ServerConfig.HTTPMetricsBind)
	if err != nil {
		return err
	}
	http.Handle("/metrics", promhttp.Handler())
	mS = &http.Server{Handler: promhttp.Handler()}

	go func() {
		err := mS.Serve(lis)
		if err != nil {
			_ = seelog.Errorf("metrics got error %v.", err)
		}
	}()

	return nil
}

func StopMetrics() {
	seelog.Infof("try to kill metrics server.")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	err := mS.Shutdown(ctx)
	if err != nil {
		_ = seelog.Errorf("stop metrics failed %v", err)
	}
	cancel()
}
