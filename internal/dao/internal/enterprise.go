// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EnterpriseDao is the data access object for table enterprise.
type EnterpriseDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns EnterpriseColumns // columns contains all the column names of Table for convenient usage.
}

// EnterpriseColumns defines and stores column names for table enterprise.
type EnterpriseColumns struct {
	Id             string // Id
	EnterpriseName string // 企业名称
	EnterpriseDesc string // 企业描述
	CreateAt       string // Create Time
	UpdateAt       string // Update Time
	DeleteAt       string // 删除时间
}

// enterpriseColumns holds the columns for table enterprise.
var enterpriseColumns = EnterpriseColumns{
	Id:             "id",
	EnterpriseName: "enterpriseName",
	EnterpriseDesc: "enterpriseDesc",
	CreateAt:       "create_at",
	UpdateAt:       "update_at",
	DeleteAt:       "delete_at",
}

// NewEnterpriseDao creates and returns a new DAO object for table data access.
func NewEnterpriseDao() *EnterpriseDao {
	return &EnterpriseDao{
		group:   "default",
		table:   "enterprise",
		columns: enterpriseColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EnterpriseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EnterpriseDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EnterpriseDao) Columns() EnterpriseColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EnterpriseDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EnterpriseDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EnterpriseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
