package service

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/xushuhui/kratos-microservice-layout/internal/biz"
	"github.com/xushuhui/kratos-microservice-layout/internal/conf"
)

type GreeterJob struct {
	uc *biz.GreeterUsecase
	ck *conf.Server
}

// NewGreeterService new a greeter service.
func NewGreeterJob(uc *biz.GreeterUsecase, ck *conf.Server) *GreeterJob {
	return &GreeterJob{uc: uc, ck: ck}
}

func (s *GreeterJob) Topic() string {
	return s.ck.Kafka.Topic.Greeter
}

func (s *GreeterJob) Handle(msg *kafka.Message) error {
	fmt.Printf("receive message key %s value %s\n", string(msg.Key), string(msg.Value))

	return nil
}
