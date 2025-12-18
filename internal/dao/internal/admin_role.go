// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminRoleDao is the data access object for the table admin_role.
type AdminRoleDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AdminRoleColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AdminRoleColumns defines and stores column names for the table admin_role.
type AdminRoleColumns struct {
	Id          string //
	SiteId      string //
	Name        string // 角色名称
	Status      string // 状态。1=可用；0=禁用
	CreatedAt   string //
	UpdatedAt   string //
	Permissions string // 权限列表。格式：权限id以,隔开
}

// adminRoleColumns holds the columns for the table admin_role.
var adminRoleColumns = AdminRoleColumns{
	Id:          "id",
	SiteId:      "site_id",
	Name:        "name",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	Permissions: "permissions",
}

// NewAdminRoleDao creates and returns a new DAO object for table data access.
func NewAdminRoleDao(handlers ...gdb.ModelHandler) *AdminRoleDao {
	return &AdminRoleDao{
		group:    "default",
		table:    "admin_role",
		columns:  adminRoleColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AdminRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AdminRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AdminRoleDao) Columns() AdminRoleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AdminRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AdminRoleDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AdminRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
