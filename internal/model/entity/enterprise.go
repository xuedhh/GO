// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Enterprise is the golang structure for table enterprise.
type Enterprise struct {
	Id             uint        `json:"id"             description:"Id"`
	EnterpriseName string      `json:"enterpriseName" description:"企业名称"`
	EnterpriseDesc string      `json:"enterpriseDesc" description:"企业描述"`
	CreateAt       *gtime.Time `json:"createAt"       description:"Create Time"`
	UpdateAt       *gtime.Time `json:"updateAt"       description:"Update Time"`
	DeleteAt       *gtime.Time `json:"deleteAt"       description:"删除时间"`
}
