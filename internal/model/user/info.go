package user

import "github.com/gogf/gf/v2/os/gtime"

type InfoInput struct {
	Id int `json:"id"`
}

type InfoOutput struct {
	Id       int         `json:"id"`
	CreateAt *gtime.Time `json:"create_at"`
}
