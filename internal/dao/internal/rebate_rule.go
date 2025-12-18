// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RebateRuleDao is the data access object for the table rebate_rule.
type RebateRuleDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  RebateRuleColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// RebateRuleColumns defines and stores column names for the table rebate_rule.
type RebateRuleColumns struct {
	Id        string //
	SiteId    string //
	Name      string //
	Status    string //
	CreatedAt string //
	UpdatedAt string //
}

// rebateRuleColumns holds the columns for the table rebate_rule.
var rebateRuleColumns = RebateRuleColumns{
	Id:        "id",
	SiteId:    "site_id",
	Name:      "name",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewRebateRuleDao creates and returns a new DAO object for table data access.
func NewRebateRuleDao(handlers ...gdb.ModelHandler) *RebateRuleDao {
	return &RebateRuleDao{
		group:    "default",
		table:    "rebate_rule",
		columns:  rebateRuleColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RebateRuleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RebateRuleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RebateRuleDao) Columns() RebateRuleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RebateRuleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RebateRuleDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *RebateRuleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
