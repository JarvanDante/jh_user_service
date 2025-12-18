// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AgentLevelDao is the data access object for the table agent_level.
type AgentLevelDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AgentLevelColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AgentLevelColumns defines and stores column names for the table agent_level.
type AgentLevelColumns struct {
	Id                  string //
	SiteId              string //
	Name                string //
	Status              string //
	AgentRegisterBonus  string // 代理注册奖励金
	SwitchSharing       string // 分成模式开关。1=开启；0=关闭
	SwitchRebate        string // 返点模式开关。1=开启；0=关闭
	PercentDeductBonus  string // 红利扣除比例。以小数记录百分比
	PercentDeductRebate string // 返水扣除比例。以小数记录百分比
	CreatedAt           string //
	UpdatedAt           string //
	SharingOptions      string // 分成模式条件列表。JSON格式。month_amount_min：团队当月盈利最小，month_amount_max：团队当月盈利最大，valid_user_count：有效人数，valid_bet_count：有效投注，percent_sports：分成比例 - 体育，percent_lottery：分成比例 - 彩票，percent_live：分成比例 - 真人，percent_egame：分成比例 - 电子游戏。
	RebateOptions       string // 返点模式条件列表。JSON格式。valid_user_count：有效人数，valid_bet_count：有效投注，percent_sports：分成比例 - 体育，percent_lottery：分成比例 - 彩票，percent_live：分成比例 - 真人，percent_egame：分成比例 - 电子游戏。
}

// agentLevelColumns holds the columns for the table agent_level.
var agentLevelColumns = AgentLevelColumns{
	Id:                  "id",
	SiteId:              "site_id",
	Name:                "name",
	Status:              "status",
	AgentRegisterBonus:  "agent_register_bonus",
	SwitchSharing:       "switch_sharing",
	SwitchRebate:        "switch_rebate",
	PercentDeductBonus:  "percent_deduct_bonus",
	PercentDeductRebate: "percent_deduct_rebate",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
	SharingOptions:      "sharing_options",
	RebateOptions:       "rebate_options",
}

// NewAgentLevelDao creates and returns a new DAO object for table data access.
func NewAgentLevelDao(handlers ...gdb.ModelHandler) *AgentLevelDao {
	return &AgentLevelDao{
		group:    "default",
		table:    "agent_level",
		columns:  agentLevelColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AgentLevelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AgentLevelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AgentLevelDao) Columns() AgentLevelColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AgentLevelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AgentLevelDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AgentLevelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
