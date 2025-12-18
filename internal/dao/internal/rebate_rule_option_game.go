// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RebateRuleOptionGameDao is the data access object for the table rebate_rule_option_game.
type RebateRuleOptionGameDao struct {
	table    string                      // table is the underlying table name of the DAO.
	group    string                      // group is the database configuration group name of the current DAO.
	columns  RebateRuleOptionGameColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler          // handlers for customized model modification.
}

// RebateRuleOptionGameColumns defines and stores column names for the table rebate_rule_option_game.
type RebateRuleOptionGameColumns struct {
	Id           string //
	SiteId       string // 站点ID
	RuleId       string // 返水规则ID
	RuleOptionId string // 活动ID
	GameId       string // 游戏ID
	Percent      string // 返水比例
	CreatedAt    string //
	UpdatedAt    string //
}

// rebateRuleOptionGameColumns holds the columns for the table rebate_rule_option_game.
var rebateRuleOptionGameColumns = RebateRuleOptionGameColumns{
	Id:           "id",
	SiteId:       "site_id",
	RuleId:       "rule_id",
	RuleOptionId: "rule_option_id",
	GameId:       "game_id",
	Percent:      "percent",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewRebateRuleOptionGameDao creates and returns a new DAO object for table data access.
func NewRebateRuleOptionGameDao(handlers ...gdb.ModelHandler) *RebateRuleOptionGameDao {
	return &RebateRuleOptionGameDao{
		group:    "default",
		table:    "rebate_rule_option_game",
		columns:  rebateRuleOptionGameColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RebateRuleOptionGameDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RebateRuleOptionGameDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RebateRuleOptionGameDao) Columns() RebateRuleOptionGameColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RebateRuleOptionGameDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RebateRuleOptionGameDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *RebateRuleOptionGameDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
