// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RiskForewarnConfigDao is the data access object for the table risk_forewarn_config.
type RiskForewarnConfigDao struct {
	table    string                    // table is the underlying table name of the DAO.
	group    string                    // group is the database configuration group name of the current DAO.
	columns  RiskForewarnConfigColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler        // handlers for customized model modification.
}

// RiskForewarnConfigColumns defines and stores column names for the table risk_forewarn_config.
type RiskForewarnConfigColumns struct {
	Id                    string //
	SiteId                string // 站点ID
	AccountToGame         string // 账户转账到游戏单笔金额预警, 0=不提示 其他=提示
	GameToAccount         string // 游戏转账到账户单笔金额预警, 0=不提示 其他=提示
	WinMoney              string // 会员单笔游戏赢得金额预警,0=不提示 其他=提示
	BetAmount             string // 会员单笔游戏投注金额预警,0=不提示 其他=提示
	AlterBankCard         string // 会员银行卡号被修改预警，0=不提示；1=提示
	LoginAreaDifference   string // 本次和上次登录异样提示，0=不提示；1=提示
	SameipRegisterTime    string // 同一IP单位时间内注册限制，0=不限时间 其他=时间，单位：分钟
	SameipRegisterCount   string // 同一IP单位时间内注册次数限制，0=不限 其他=提示
	CreatedAt             string //
	UpdatedAt             string //
	IsAccountToGame       string // 是否开启会员账户-游戏预警
	IsGameToAccount       string // 是否开启会员游戏-账户预警
	IsWinMoney            string // 是否开启游戏输赢预警
	IsBetAmount           string // 是否开启游戏投注预警
	IsAlterBankCard       string // 是否开启修改银行卡预警
	IsLoginAreaDifference string // 是否开启登录地区不同预警
	IsSameipRegisterCount string // 是否开启相同IP注册次数预警
}

// riskForewarnConfigColumns holds the columns for the table risk_forewarn_config.
var riskForewarnConfigColumns = RiskForewarnConfigColumns{
	Id:                    "id",
	SiteId:                "site_id",
	AccountToGame:         "account_to_game",
	GameToAccount:         "game_to_account",
	WinMoney:              "win_money",
	BetAmount:             "bet_amount",
	AlterBankCard:         "alter_bank_card",
	LoginAreaDifference:   "login_area_difference",
	SameipRegisterTime:    "sameip_register_time",
	SameipRegisterCount:   "sameip_register_count",
	CreatedAt:             "created_at",
	UpdatedAt:             "updated_at",
	IsAccountToGame:       "is_account_to_game",
	IsGameToAccount:       "is_game_to_account",
	IsWinMoney:            "is_win_money",
	IsBetAmount:           "is_bet_amount",
	IsAlterBankCard:       "is_alter_bank_card",
	IsLoginAreaDifference: "is_login_area_difference",
	IsSameipRegisterCount: "is_sameip_register_count",
}

// NewRiskForewarnConfigDao creates and returns a new DAO object for table data access.
func NewRiskForewarnConfigDao(handlers ...gdb.ModelHandler) *RiskForewarnConfigDao {
	return &RiskForewarnConfigDao{
		group:    "default",
		table:    "risk_forewarn_config",
		columns:  riskForewarnConfigColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RiskForewarnConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RiskForewarnConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RiskForewarnConfigDao) Columns() RiskForewarnConfigColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RiskForewarnConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RiskForewarnConfigDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *RiskForewarnConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
