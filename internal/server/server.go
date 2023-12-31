package server

import (
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/xushuhui/kratos-microservice-layout/internal/conf"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewDiscovery, NewRegistrar,NewKafkaServer)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(conf.Nacos.Addr, 443),
	}
	cc := constant.ClientConfig{
		NamespaceId: "public",
		TimeoutMs:   5000,
		Username:    conf.Nacos.Username,
		Password:    conf.Nacos.Password,
	}
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ServerConfigs: sc,
			ClientConfig:  &cc,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	r := nacos.New(client)
	return r
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(conf.Nacos.Addr, 443),
	}
	cc := constant.ClientConfig{
		NamespaceId: "public",
		TimeoutMs:   5000,
		Username:    conf.Nacos.Username,
		Password:    conf.Nacos.Password,
	}
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ServerConfigs: sc,
			ClientConfig:  &cc,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	r := nacos.New(client)
	return r
}
