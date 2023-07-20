package service

import (
	"context"

	"github.com/xushuhui/kratos-microservice-layout/api"
	"github.com/xushuhui/kratos-microservice-layout/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	api.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &api.HelloReply{Message: "Hello " + g.Hello}, nil
}
