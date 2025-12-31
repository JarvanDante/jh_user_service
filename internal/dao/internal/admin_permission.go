// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminPermissionDao is the data access object for table admin_permission.
type AdminPermissionDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns AdminPermissionColumns // columns contains all the column names of Table for convenient usage.
}

// AdminPermissionColumns defines and stores column names for table admin_permission.
type AdminPermissionColumns struct {
	Id          string // 权限ID
	ParentId    string // 父权限ID
	Type        string // 权限类型。1=菜单；2=操作权限
	Name        string // 权限名称
	BackendUrl  string // 后台URL
	FrontendUrl string // 前台URL
	Status      string // 状态。1=启用；0=禁用
	Sort        string // 排序
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// adminPermissionColumns holds the columns for table admin_permission.
var adminPermissionColumns = AdminPermissionColumns{
	Id:          "id",
	ParentId:    "parent_id",
	Type:        "type",
	Name:        "name",
	BackendUrl:  "backend_url",
	FrontendUrl: "frontend_url",
	Status:      "status",
	Sort:        "sort",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewAdminPermissionDao creates and returns a new DAO object for table data access.
func NewAdminPermissionDao() *AdminPermissionDao {
	return &AdminPermissionDao{
		group:   "default",
		table:   "admin_permission",
		columns: adminPermissionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminPermissionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminPermissionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminPermissionDao) Columns() AdminPermissionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminPermissionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminPermissionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}
