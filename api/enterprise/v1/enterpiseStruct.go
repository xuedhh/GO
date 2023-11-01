package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"user/internal/model/enterprise"
)

type CreateEnterpriseUserReq struct {
	g.Meta         `path:"/create/enterpriseUser" tags:"enterpriseService" method:"post" summary:""`
	EnterpriseName string `v:"required|length:1,10"`
	EnterpriseDesc string `v:"required|length:1,10"`
}

type CreateEnterpriseUserRes struct {
	Id             int    `json:"id"`
	EnterpriseName string `json:"enterpriseName"`
	EnterpriseDesc string `json:"enterpriseDesc"`
}

type DeleteEnterpriseUserReq struct {
	g.Meta `path:"/delete/enterpriseUser" tags:"enterpriseService" method:"post" summary:""`
	Id     int `v:"required"`
}

type DeleteEnterpriseUserRes struct {
	Id       int  `json:"id"`
	IsDelete bool `json:"is_delete"`
}

type UpdateEnterpriseUserReq struct {
	g.Meta         `path:"/update/enterpriseUser" tags:"enterpriseService" method:"post" summary:""`
	Id             int    `v:"required" json:"id"`
	EnterpriseName string `v:"required|length:1,20" json:"enterprise_name"`
	EnterpriseDesc string `v:"length:1,20" json:"enterprise_desc"`
}

type UpdateEnterpriseUserRes struct {
	Id int `json:"id"`
}

type SelectEnterpriseUserReq struct {
	g.Meta `path:"/select/enterpriseUser" tags:"enterpriseService" method:"post" summary:""`
	Id     int
}

type SelectEnterpriseUserRes struct {
	Id             int         `json:"id"`
	EnterpriseName string      `json:"enterprise_name"`
	EnterpriseDesc string      `json:"enterprise_desc"`
	CreateAt       *gtime.Time `json:"create_at"`
	UpdateAt       *gtime.Time `json:"update_at"`
}

type SelectEnterpriseUserAllReq struct {
	g.Meta `path:"/selectAll/enterpriseUser" tags:"enterpriseService" method:"post" summary:""`
}

type SelectEnterpriseUserAllRes struct {
	enterprise.SelectEnterpriseAllOutput
}

type SelectEnterpriseUserPageReq struct {
	g.Meta `path:"/selectPage/enterpriseUser" tags:"enterpriseService" method:"post" summary:""`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type SelectEnterpriseUserPageRes struct {
	enterprise.SelectEnterprisePageOutput
}
