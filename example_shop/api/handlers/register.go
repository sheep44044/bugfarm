package handlers

import (
	"context"
	"example_shop/api/rpc"
	"example_shop/kitex_gen/example/shop/user"

	"github.com/cloudwego/hertz/pkg/app"
)

func Register(ctx context.Context, c *app.RequestContext) {
	var registerVar UserParam
	if err := c.Bind(&registerVar); err != nil {
		return
	}

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		return
	}

	err := rpc.Register(context.Background(), &user.RegisterReq{
		UserName: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		return
	}

	return
}
