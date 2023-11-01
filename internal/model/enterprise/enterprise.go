package enterprise

import (
	"github.com/gogf/gf/v2/os/gtime"
	"user/internal/model/entity"
)

type CreateEnterpriseInput struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type CreateEnterpriseOutput struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type DeleteEnterpriseInput struct {
	Id int `json:"id"`
}

type DeleteEnterpriseOutput struct {
	Id int
	OK bool
}

type UpdateEnterpriseInput struct {
	Id   int    `json:"id"`
	Name string `json:"enterprise_name"`
	Desc string `json:"enterprise_desc"`
}

type UpdateEnterpriseOutput struct {
	Id int
}

type SelectEnterpriseInput struct {
	Id int
}

type SelectEnterpriseOutput struct {
	Id             int         `json:"id"`
	EnterpriseName string      `json:"enterprise_name"`
	EnterpriseDesc string      `json:"enterprise_desc"`
	CreateAt       *gtime.Time `json:"create_at"`
	UpdateAt       *gtime.Time `json:"update_at"`
}

type SelectEnterpriseAllInput struct{}

type SelectEnterpriseAllOutput struct {
	Enterprise []*entity.Enterprise `json:"enterprise"`
	//Id             int         `json:"id"`
	//EnterpriseName string      `json:"enterprise_name"`
	//EnterpriseDesc string      `json:"enterprise_desc"`
	//CreateAt       *gtime.Time `json:"create_at"`
	//UpdateAt       *gtime.Time `json:"update_at"`
}

type SelectEnterprisePageInput struct {
	Limit  int
	Offset int
}

type SelectEnterprisePageOutput struct {
	Limit      int                  `json:"limit"`
	Offset     int                  `json:"offset"`
	Total      int                  `json:"total"`
	Enterprise []*entity.Enterprise `json:"enterprise"`
}
