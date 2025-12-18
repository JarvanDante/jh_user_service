// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SiteDenyDao is the data access object for the table site_deny.
type SiteDenyDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SiteDenyColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SiteDenyColumns defines and stores column names for the table site_deny.
type SiteDenyColumns struct {
	Id        string //
	SiteId    string //
	Type      string // 屏蔽类型。1=ip；2=地址
	Ip        string // IP地址
	Address   string // 地区名称
	CreatedAt string //
	UpdatedAt string //
}

// siteDenyColumns holds the columns for the table site_deny.
var siteDenyColumns = SiteDenyColumns{
	Id:        "id",
	SiteId:    "site_id",
	Type:      "type",
	Ip:        "ip",
	Address:   "address",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSiteDenyDao creates and returns a new DAO object for table data access.
func NewSiteDenyDao(handlers ...gdb.ModelHandler) *SiteDenyDao {
	return &SiteDenyDao{
		group:    "default",
		table:    "site_deny",
		columns:  siteDenyColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SiteDenyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SiteDenyDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SiteDenyDao) Columns() SiteDenyColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SiteDenyDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SiteDenyDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SiteDenyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
