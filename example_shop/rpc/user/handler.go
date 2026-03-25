package main

import (
	"context"
	"example_shop/kitex_gen/example/shop/user"
	"example_shop/rpc/user/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	resp = new(user.RegisterResp)
	if len(req.UserName) == 0 || len(req.Password) == 0 {
		return resp, nil
	}

	registerService := service.NewRegister(ctx)
	err = registerService.Register(req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	// TODO: Your code here...
	return
}
