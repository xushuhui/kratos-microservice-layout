package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/xushuhui/kratos-microservice-layout/internal/conf"
	"github.com/xushuhui/kratos-microservice-layout/internal/service"

	kafka "github.com/xushuhui/kratos-kafka"
)

func NewKafkaServer(config *conf.Server, gj *service.GreeterJob) (*kafka.Server) {
	// init consumers
	var consumers []kafka.Consumer
	topics := []string{config.Kafka.Topic.Greeter}
	srvConsumer, err := kafka.NewGroupConsumer(config.Kafka.Addr, config.Kafka.Group, topics)
	if err != nil {
		log.Fatal(err)
		return  nil
	}
	consumers = append(consumers, srvConsumer)
	return kafka.NewServer(
		kafka.Consumers(consumers),
		kafka.Handlers(gj))
}
