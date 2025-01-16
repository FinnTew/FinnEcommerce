package client

import (
	"github.com/FinnTew/FinnEcommerce/src/internal/client/rpc/auth"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/user/conf"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	AuthClient auth.RPCClient
	once       sync.Once
)

func Init() {
	initClient()
}

func initClient() {
	once.Do(func() {
		// AuthClient init
		r, _ := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
		AuthClient, _ = auth.NewRPCClient("auth", client.WithResolver(r))
	})
}
