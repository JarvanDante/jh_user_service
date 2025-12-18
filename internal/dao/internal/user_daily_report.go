// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDailyReportDao is the data access object for the table user_daily_report.
type UserDailyReportDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  UserDailyReportColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// UserDailyReportColumns defines and stores column names for the table user_daily_report.
type UserDailyReportColumns struct {
	Id                   string //
	SiteId               string // 站点ID
	UserId               string // 会员ID
	AgentId              string // 代理ID
	Username             string // 用户名
	ReportDate           string // 报表日期
	BetCount             string // 注单数量
	BetAmount            string // 投注总额
	ValidBetAmount       string // 有效投注金额（投注时间）
	ValidBetAmountSettle string // 有效投注金额（结算时间）
	WinOrLose            string // 输赢结果
	RechargeAmount       string // 充值金额
	WithdrawAmount       string // 提现金额
	BonusAmount          string // 红利金额
	RebateAmount         string // 返水金额
	FeeAmount            string // 手续费
	CreatedAt            string //
	UpdatedAt            string //
}

// userDailyReportColumns holds the columns for the table user_daily_report.
var userDailyReportColumns = UserDailyReportColumns{
	Id:                   "id",
	SiteId:               "site_id",
	UserId:               "user_id",
	AgentId:              "agent_id",
	Username:             "username",
	ReportDate:           "report_date",
	BetCount:             "bet_count",
	BetAmount:            "bet_amount",
	ValidBetAmount:       "valid_bet_amount",
	ValidBetAmountSettle: "valid_bet_amount_settle",
	WinOrLose:            "win_or_lose",
	RechargeAmount:       "recharge_amount",
	WithdrawAmount:       "withdraw_amount",
	BonusAmount:          "bonus_amount",
	RebateAmount:         "rebate_amount",
	FeeAmount:            "fee_amount",
	CreatedAt:            "created_at",
	UpdatedAt:            "updated_at",
}

// NewUserDailyReportDao creates and returns a new DAO object for table data access.
func NewUserDailyReportDao(handlers ...gdb.ModelHandler) *UserDailyReportDao {
	return &UserDailyReportDao{
		group:    "default",
		table:    "user_daily_report",
		columns:  userDailyReportColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserDailyReportDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserDailyReportDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserDailyReportDao) Columns() UserDailyReportColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserDailyReportDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserDailyReportDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserDailyReportDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
