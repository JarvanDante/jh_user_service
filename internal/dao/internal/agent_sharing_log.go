// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AgentSharingLogDao is the data access object for the table agent_sharing_log.
type AgentSharingLogDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  AgentSharingLogColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// AgentSharingLogColumns defines and stores column names for the table agent_sharing_log.
type AgentSharingLogColumns struct {
	Id                    string //
	SiteId                string // 站点ID
	StartDate             string // 开始日期
	EndDate               string // 结束日期
	AgentId               string // 代理ID
	AgentUsername         string // 代理账号
	UserCount             string // 有效会员数
	RechargeAmount        string // 充值总计
	WithdrawAmount        string // 提现总计
	BonusAmount           string // 红利总计
	RebateAmount          string // 返水总计
	ValidBetAmount        string // 有效投注总计
	WinOrLose             string // 输赢总计
	FeeAmount             string // 手续费总计
	AdministrationExpense string // 行政费用
	CalculateCommission   string // 程序计算出的代理佣金
	Commission            string // 发放佣金
	Status                string // 状态。1=未结算；2=结算成功；3=结算失败；4=未盈利
	AdminId               string // 后台员工ID
	CreatedAt             string //
	UpdatedAt             string //
	CommissionData        string // 返佣金额计算公式数据
}

// agentSharingLogColumns holds the columns for the table agent_sharing_log.
var agentSharingLogColumns = AgentSharingLogColumns{
	Id:                    "id",
	SiteId:                "site_id",
	StartDate:             "start_date",
	EndDate:               "end_date",
	AgentId:               "agent_id",
	AgentUsername:         "agent_username",
	UserCount:             "user_count",
	RechargeAmount:        "recharge_amount",
	WithdrawAmount:        "withdraw_amount",
	BonusAmount:           "bonus_amount",
	RebateAmount:          "rebate_amount",
	ValidBetAmount:        "valid_bet_amount",
	WinOrLose:             "win_or_lose",
	FeeAmount:             "fee_amount",
	AdministrationExpense: "administration_expense",
	CalculateCommission:   "calculate_commission",
	Commission:            "commission",
	Status:                "status",
	AdminId:               "admin_id",
	CreatedAt:             "created_at",
	UpdatedAt:             "updated_at",
	CommissionData:        "commission_data",
}

// NewAgentSharingLogDao creates and returns a new DAO object for table data access.
func NewAgentSharingLogDao(handlers ...gdb.ModelHandler) *AgentSharingLogDao {
	return &AgentSharingLogDao{
		group:    "default",
		table:    "agent_sharing_log",
		columns:  agentSharingLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AgentSharingLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AgentSharingLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AgentSharingLogDao) Columns() AgentSharingLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AgentSharingLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AgentSharingLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AgentSharingLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
