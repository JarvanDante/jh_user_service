// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminCustomFieldDao is the data access object for the table admin_custom_field.
type AdminCustomFieldDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  AdminCustomFieldColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// AdminCustomFieldColumns defines and stores column names for the table admin_custom_field.
type AdminCustomFieldColumns struct {
	Id        string //
	SiteId    string // 站点ID
	AdminId   string // 员工ID
	Page      string // 页面。1=会员列表；默认=1
	Fields    string // 字段列表。以,隔开
	CreatedAt string //
	UpdatedAt string //
}

// adminCustomFieldColumns holds the columns for the table admin_custom_field.
var adminCustomFieldColumns = AdminCustomFieldColumns{
	Id:        "id",
	SiteId:    "site_id",
	AdminId:   "admin_id",
	Page:      "page",
	Fields:    "fields",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewAdminCustomFieldDao creates and returns a new DAO object for table data access.
func NewAdminCustomFieldDao(handlers ...gdb.ModelHandler) *AdminCustomFieldDao {
	return &AdminCustomFieldDao{
		group:    "default",
		table:    "admin_custom_field",
		columns:  adminCustomFieldColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AdminCustomFieldDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AdminCustomFieldDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AdminCustomFieldDao) Columns() AdminCustomFieldColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AdminCustomFieldDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AdminCustomFieldDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AdminCustomFieldDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
