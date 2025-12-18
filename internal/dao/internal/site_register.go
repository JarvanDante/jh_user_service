// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SiteRegisterDao is the data access object for the table site_register.
type SiteRegisterDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  SiteRegisterColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// SiteRegisterColumns defines and stores column names for the table site_register.
type SiteRegisterColumns struct {
	Id        string //
	SiteId    string //
	Type      string //
	Name      string // 名称
	FieldName string // 字段标识
	Display   string // 是否显示。1=显示；0=不显示
	Required  string // 是否必填。1=必填；0=选填
	CreatedAt string //
	UpdatedAt string //
}

// siteRegisterColumns holds the columns for the table site_register.
var siteRegisterColumns = SiteRegisterColumns{
	Id:        "id",
	SiteId:    "site_id",
	Type:      "type",
	Name:      "name",
	FieldName: "field_name",
	Display:   "display",
	Required:  "required",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSiteRegisterDao creates and returns a new DAO object for table data access.
func NewSiteRegisterDao(handlers ...gdb.ModelHandler) *SiteRegisterDao {
	return &SiteRegisterDao{
		group:    "default",
		table:    "site_register",
		columns:  siteRegisterColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SiteRegisterDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SiteRegisterDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SiteRegisterDao) Columns() SiteRegisterColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SiteRegisterDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SiteRegisterDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SiteRegisterDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
