// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserLevelDao is the data access object for the table user_level.
type UserLevelDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UserLevelColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UserLevelColumns defines and stores column names for the table user_level.
type UserLevelColumns struct {
	Id                 string //
	SiteId             string // 站点ID
	Name               string // 层级名称
	IsRebate           string // 是否返水.1=返水；0=不返水
	RebateRuleId       string // 返水规则ID
	DailyWithdrawTimes string // 单日提款次数上限
	LoginUrl           string // 专用登录网址
	Status             string // 状态.1=可用；0=禁用
	CreatedAt          string //
	UpdatedAt          string //
}

// userLevelColumns holds the columns for the table user_level.
var userLevelColumns = UserLevelColumns{
	Id:                 "id",
	SiteId:             "site_id",
	Name:               "name",
	IsRebate:           "is_rebate",
	RebateRuleId:       "rebate_rule_id",
	DailyWithdrawTimes: "daily_withdraw_times",
	LoginUrl:           "login_url",
	Status:             "status",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// NewUserLevelDao creates and returns a new DAO object for table data access.
func NewUserLevelDao(handlers ...gdb.ModelHandler) *UserLevelDao {
	return &UserLevelDao{
		group:    "default",
		table:    "user_level",
		columns:  userLevelColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserLevelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserLevelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserLevelDao) Columns() UserLevelColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserLevelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserLevelDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserLevelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
