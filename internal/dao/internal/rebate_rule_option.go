// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RebateRuleOptionDao is the data access object for the table rebate_rule_option.
type RebateRuleOptionDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  RebateRuleOptionColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// RebateRuleOptionColumns defines and stores column names for the table rebate_rule_option.
type RebateRuleOptionColumns struct {
	Id               string //
	SiteId           string //
	RuleId           string // 返水规则ID
	DailyValidBetMin string // 单日有效投注范围最小金额
	DailyValidBetMax string // 单日有效投注范围最大金额
	CreatedAt        string //
	UpdatedAt        string //
}

// rebateRuleOptionColumns holds the columns for the table rebate_rule_option.
var rebateRuleOptionColumns = RebateRuleOptionColumns{
	Id:               "id",
	SiteId:           "site_id",
	RuleId:           "rule_id",
	DailyValidBetMin: "daily_valid_bet_min",
	DailyValidBetMax: "daily_valid_bet_max",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// NewRebateRuleOptionDao creates and returns a new DAO object for table data access.
func NewRebateRuleOptionDao(handlers ...gdb.ModelHandler) *RebateRuleOptionDao {
	return &RebateRuleOptionDao{
		group:    "default",
		table:    "rebate_rule_option",
		columns:  rebateRuleOptionColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RebateRuleOptionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RebateRuleOptionDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RebateRuleOptionDao) Columns() RebateRuleOptionColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RebateRuleOptionDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RebateRuleOptionDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *RebateRuleOptionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
