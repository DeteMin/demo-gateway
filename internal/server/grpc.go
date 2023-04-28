package server

import (
	"net"

	"github.com/cihub/seelog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/grpc-ecosystem/go-grpc-prometheus"

	"demoapp/internal/config"
	"demoapp/internal/service"
	"demoapp/proto"
)

var t *grpc.Server

func StartGRPC() error {
	lis, err := net.Listen("tcp", config.ServerConfig.GRPCServiceBind)
	if err != nil {
		return err
	}

	t = grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	// todo 改成你自己的服务
	s := &service.TranslateService{}

	proto.RegisterTranslateServiceServer(t, s)
	reflection.Register(t)
	grpc_prometheus.Register(t)
	if config.ServerConfig.EnableHandlingTime {
		grpc_prometheus.EnableHandlingTimeHistogram()
	}
	go func() {
		err := t.Serve(lis)
		if err != nil {
			_ = seelog.Errorf("start gRPC server failed %v.", err)
		}
	}()

	return nil
}

func StopGRPC() {
	seelog.Infof("try to stop grpc server.")
	t.GracefulStop()
}
