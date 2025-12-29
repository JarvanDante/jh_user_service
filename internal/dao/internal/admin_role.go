// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminRoleDao is the data access object for table admin_role.
type AdminRoleDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns AdminRoleColumns // columns contains all the column names of Table for convenient usage.
}

// AdminRoleColumns defines and stores column names for table admin_role.
type AdminRoleColumns struct {
	Id          string // 角色ID
	SiteId      string // 站点ID
	Name        string // 角色名称
	Status      string // 状态。1=启用；0=禁用
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	Permissions string // 权限列表。格式：权限id以,隔开
}

// adminRoleColumns holds the columns for table admin_role.
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
func NewAdminRoleDao() *AdminRoleDao {
	return &AdminRoleDao{
		group:   "default",
		table:   "admin_role",
		columns: adminRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminRoleDao) Columns() AdminRoleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}
