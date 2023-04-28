package server

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/cihub/seelog"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"demoapp/internal/config"
	"demoapp/proto"
)

var gS *http.Server
var cancel context.CancelFunc

func StartGateWay() error {
	ctx := context.Background()

	ctx, cancel = context.WithCancel(ctx)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor)}

	err := proto.RegisterTranslateServiceHandlerFromEndpoint(ctx, mux, "localhost:8111", opts)
	if err != nil {
		return err
	}

	lis, err := net.Listen("tcp", config.ServerConfig.HTTPServiceBind)
	if err != nil {
		return err
	}
	gS = &http.Server{Handler: mux}

	go func() {
		err := gS.Serve(lis)
		if err != nil {
			_ = seelog.Errorf("gateway got error %v", err)
		}
	}()

	return nil
}

func StopGateWay() {
	ctx, c := context.WithTimeout(context.Background(), 30*time.Second)
	err := gS.Shutdown(ctx)
	if err != nil {
		_ = seelog.Errorf("stop gateway failed %v", err)
	}
	c()
	cancel()
}
