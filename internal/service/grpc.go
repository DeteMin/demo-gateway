package service

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/cihub/seelog"

	"demoapp/internal/config"
	"demoapp/proto"
)

// todo 定义服务的实现类
type TranslateService struct {
}

// todo 继承gRPC定义的接口
func (s *TranslateService) Translate(ctx context.Context, in *proto.Request) (*proto.Response, error) {
	return (&_handler{}).Translate(ctx, in)
}

// todo 由于gRPC服务注册的处理对象全局唯一，推荐把处理类绑定到一个新的对象上
type _handler struct {
}

// todo 真实的业务逻辑处理
func (h *_handler) Translate(ctx context.Context, in *proto.Request) (*proto.Response, error) {
	sCtx, _ := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(config.ServerConfig.HandleTimeout))

	// 模拟业务逻辑进行延时 0~3
	d := rand.Int31n(3000) + 1
	t := time.NewTicker(time.Duration(d) * time.Millisecond)
	seelog.Infof("got request wait ms %d", d)

	select {
	case <-ctx.Done():
		seelog.Info("client cancel")
		return nil, errors.New("client cancel")
	case <-sCtx.Done():
		seelog.Info("server timeout")
		return nil, errors.New("server timeout")
	case <-t.C:
		seelog.Info("run success")
		break
	}

	return &proto.Response{Retcode: "000000", Retdesc: "success", Result: "mock"}, nil
}
