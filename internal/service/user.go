package service

import (
	"context"
	"user/internal/model/user"
)

type IUser interface {
	SignIn(ctx context.Context, input user.SignInInput) (output user.SignUpOutput, err error)
	UpdateUser(ctx context.Context, input user.UpdateUserInput) (output user.UpdateUserOutput, err error)
}

var localUser IUser

func User() IUser {
	if localUser == nil {
		panic("user service not init")
	}
	return localUser
}

func RegisterUser(userService IUser) {
	localUser = userService
}
