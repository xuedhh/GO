package user

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"user/api/user"
	v1 "user/api/user/v1"
	user2 "user/internal/model/user"
	"user/internal/service"
)

type ControllerV1 struct{}

func (c *ControllerV1) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error) {
	updateUser, err := service.User().UpdateUser(ctx, user2.UpdateUserInput{
		Name: req.Username,
		Id:   req.Id,
	})
	fmt.Println(updateUser)
	if err != nil {
		return
	}
	if updateUser.OK == false {
		return nil, err
	}
	res = &v1.UpdateUserRes{
		UpdateUserInfo: updateUser,
	}
	return res, nil
	// 判断 updateUser 是否为空
	//res = &v1.UpdateUserRes{
	//	UpdateUserInfo: updateUser,
	//}
	//return res, nil
}

func User() user.IUser {
	return &ControllerV1{}
}

func (c *ControllerV1) Users(ctx context.Context, req *v1.UserReq) (res *v1.UserRes, err error) {
	//TODO implement me
	g.RequestFromCtx(ctx).Response.Writeln("我是用户")
	return
}

func (c *ControllerV1) SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error) {
	output, err := service.User().SignIn(ctx, user2.SignInInput{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return
	}
	signOutput := user2.SignUpOutput{
		Id:   output.Id,
		Name: output.Name,
		Info: output.Info,
	}
	res = &v1.SignInRes{
		signOutput,
	}
	return res, nil
}
