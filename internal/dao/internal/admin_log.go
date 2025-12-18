// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminLogDao is the data access object for the table admin_log.
type AdminLogDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AdminLogColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AdminLogColumns defines and stores column names for the table admin_log.
type AdminLogColumns struct {
	Id            string //
	SiteId        string //
	AdminId       string //
	AdminUsername string //
	Ip            string //
	CreatedAt     string //
	Remark        string //
}

// adminLogColumns holds the columns for the table admin_log.
var adminLogColumns = AdminLogColumns{
	Id:            "id",
	SiteId:        "site_id",
	AdminId:       "admin_id",
	AdminUsername: "admin_username",
	Ip:            "ip",
	CreatedAt:     "created_at",
	Remark:        "remark",
}

// NewAdminLogDao creates and returns a new DAO object for table data access.
func NewAdminLogDao(handlers ...gdb.ModelHandler) *AdminLogDao {
	return &AdminLogDao{
		group:    "default",
		table:    "admin_log",
		columns:  adminLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AdminLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AdminLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AdminLogDao) Columns() AdminLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AdminLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AdminLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AdminLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
