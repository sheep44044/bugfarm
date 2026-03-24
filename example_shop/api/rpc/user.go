package rpc

import (
	"context"
	"example_shop/kitex_gen/example/shop/user"
	"example_shop/kitex_gen/example/shop/user/userservice"
	"log"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func InitUserRpcClient() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	c, err := userservice.NewClient("example.shop.user", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	userClient = c
}

func Register(ctx context.Context, req *user.RegisterReq) error {
	resp, err := userClient.Register(ctx, req)
	if err != nil {
		log.Println(err)
	}

	if resp != nil {
		log.Println(resp)
	}

	return nil
}
