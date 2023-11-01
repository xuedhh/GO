package user

import (
	"context"
	v1 "user/api/user/v1"
)

type IUser interface {
	Users(ctx context.Context, req *v1.UserReq) (res *v1.UserRes, err error)
	SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error)
	UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error)
}
