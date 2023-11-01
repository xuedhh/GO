package user

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"user/internal/dao"
	"user/internal/model/do"
	"user/internal/model/entity"
	"user/internal/model/user"
	"user/internal/service"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

func (s *sUser) SignIn(ctx context.Context, input user.SignInInput) (output user.SignUpOutput, err error) {
	var users *entity.User
	dao.User.Ctx(ctx).Where(do.User{
		Nickname: input.Username,
		Password: input.Password,
	}).Scan(&users)

	if err != nil {
		return output, err
	}
	if users == nil {
		return output, gerror.Newf("用户名或密码错误")
	}
	id := users.Id
	output.Id = int(id)
	output.Name = users.Nickname

	infoOutput := user.InfoOutput{
		Id:       int(users.Id),
		CreateAt: users.CreateAt,
	}
	output.Info = infoOutput
	return output, nil
}

func (s *sUser) UpdateUser(ctx context.Context, input user.UpdateUserInput) (output user.UpdateUserOutput, err error) {
	var users *entity.User
	dao.User.Ctx(ctx).Where(do.User{
		Id: input.Id,
	}).Scan(&users)
	if err != nil {
		return output, err
	}
	// 用户不存在
	if users == nil {
		return output, gerror.Newf("用户不存在")
	}
	// 替换用户信息
	//update, err := g.Model("user").Data(g.Map{"Nickname": input.Name}).Where(do.User{Id: input.Id}).Update()
	//fmt.Println(update)
	update, err := dao.User.Ctx(ctx).Data(do.User{
		Nickname: input.Name,
	}).Where(do.User{
		Id: input.Id,
	}).Update()
	fmt.Println(update)
	if err != nil {
		output.OK = false
		output.Message = "修改失败"
		return output, err
	}
	output.OK = true
	output.Message = "修改成功"
	return output, nil

}
