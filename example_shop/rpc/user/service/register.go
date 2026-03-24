package service

import (
	"context"
	"example_shop/kitex_gen/example/shop/user"
	"example_shop/rpc/user/dal/db"
)

type RegisterService struct {
	ctx context.Context
}

func NewRegister(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

func (s *RegisterService) Register(req *user.RegisterReq) error {
	return db.Register(s.ctx, []*db.User{{
		UserName: req.UserName,
		Password: req.Password,
	}})
}
