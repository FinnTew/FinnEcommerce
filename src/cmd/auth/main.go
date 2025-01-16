package main

import (
	"context"
	authk "github.com/FinnTew/FinnEcommerce/src/internal/client/kitex_gen/auth"
	"github.com/FinnTew/FinnEcommerce/src/internal/client/rpc/auth"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
)

func main() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Fatal(err)
	}
	cli, err := auth.NewRPCClient("auth", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}

	res, err := cli.DeliverTokenByRPC(
		context.TODO(),
		&authk.DeliverTokenReq{
			UserId: 1,
		},
	)

	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)

}
