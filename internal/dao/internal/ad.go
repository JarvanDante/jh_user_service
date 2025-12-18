// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdDao is the data access object for the table ad.
type AdDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AdColumns          // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AdColumns defines and stores column names for the table ad.
type AdColumns struct {
	Id          string //
	SiteId      string //
	Type        string //
	Name        string //
	Image       string //
	Url         string //
	StartTime   string //
	ExpiredTime string //
	Sort        string //
	Status      string //
	BeforeLogin string //
	Position    string // banner图位置
	CreatedAt   string //
	UpdatedAt   string //
}

// adColumns holds the columns for the table ad.
var adColumns = AdColumns{
	Id:          "id",
	SiteId:      "site_id",
	Type:        "type",
	Name:        "name",
	Image:       "image",
	Url:         "url",
	StartTime:   "start_time",
	ExpiredTime: "expired_time",
	Sort:        "sort",
	Status:      "status",
	BeforeLogin: "before_login",
	Position:    "position",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewAdDao creates and returns a new DAO object for table data access.
func NewAdDao(handlers ...gdb.ModelHandler) *AdDao {
	return &AdDao{
		group:    "default",
		table:    "ad",
		columns:  adColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AdDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AdDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AdDao) Columns() AdColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AdDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AdDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AdDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
