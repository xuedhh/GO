package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"user/internal/model/user"
)

type UserReq struct {
	g.Meta `path:"/user" tags:"user" method:"get" summary:"You first hello api"`
}
type UserRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type SignInReq struct {
	g.Meta    `path:"/signIn" tags:"user" method:"post" summary:"You first hello api"`
	Username  string `v:"required|length:1,20"`
	Password  string `v:"required|length:1,20"`
	Password2 string `v:"required|length:1,20|same:Password"`
}

type SignInRes struct {
	user.SignUpOutput
}

// 修改用户名

type UpdateUserReq struct {
	g.Meta   `path:"/update_username" tags:"username" method:"post" summary:"You first hello api"`
	Username string `v:"required|length:1,20"`
	Id       int    `v:"required|length:1,20"`
}

type UpdateUserRes struct {
	UpdateUserInfo user.UpdateUserOutput `json:"update_user_info"`
}
