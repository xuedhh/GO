// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Enterprise is the golang structure of table enterprise for DAO operations like Where/Data.
type Enterprise struct {
	g.Meta         `orm:"table:enterprise, do:true"`
	Id             interface{} // Id
	EnterpriseName interface{} // 企业名称
	EnterpriseDesc interface{} // 企业描述
	CreateAt       *gtime.Time // Create Time
	UpdateAt       *gtime.Time // Update Time
	DeleteAt       *gtime.Time // 删除时间
}
