// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BetLogDailyDao is the data access object for the table bet_log_daily.
type BetLogDailyDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  BetLogDailyColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// BetLogDailyColumns defines and stores column names for the table bet_log_daily.
type BetLogDailyColumns struct {
	Id                   string //
	SiteId               string // 站点ID
	BetDate              string // 投注日期
	GameId               string // 游戏ID
	AgentId              string // 代理ID
	GameType             string // 游戏类型
	UserId               string // 会员ID
	Username             string // 会员用户名
	BetAmount            string // 投注金额
	BetCount             string // 注单数量
	ValidBetAmount       string // 有效投注金额有效投注金额（投注时间）
	ValidBetAmountSettle string // 有效投注金额（结算时间）
	WinOrLose            string // 输赢结果
	CreatedAt            string //
	UpdatedAt            string //
}

// betLogDailyColumns holds the columns for the table bet_log_daily.
var betLogDailyColumns = BetLogDailyColumns{
	Id:                   "id",
	SiteId:               "site_id",
	BetDate:              "bet_date",
	GameId:               "game_id",
	AgentId:              "agent_id",
	GameType:             "game_type",
	UserId:               "user_id",
	Username:             "username",
	BetAmount:            "bet_amount",
	BetCount:             "bet_count",
	ValidBetAmount:       "valid_bet_amount",
	ValidBetAmountSettle: "valid_bet_amount_settle",
	WinOrLose:            "win_or_lose",
	CreatedAt:            "created_at",
	UpdatedAt:            "updated_at",
}

// NewBetLogDailyDao creates and returns a new DAO object for table data access.
func NewBetLogDailyDao(handlers ...gdb.ModelHandler) *BetLogDailyDao {
	return &BetLogDailyDao{
		group:    "default",
		table:    "bet_log_daily",
		columns:  betLogDailyColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BetLogDailyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BetLogDailyDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BetLogDailyDao) Columns() BetLogDailyColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BetLogDailyDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BetLogDailyDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BetLogDailyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
