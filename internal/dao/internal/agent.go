// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AgentDao is the data access object for the table agent.
type AgentDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AgentColumns       // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AgentColumns defines and stores column names for the table agent.
type AgentColumns struct {
	Id                  string //
	SiteId              string // 站点ID
	Username            string // 用户名
	Password            string // 密码
	Status              string // 审核状态。0=停用；1=启用
	StatusCheck         string // 审核状态。0=待审核；1=通过；2=未通过
	Level               string // 代理层级
	Realname            string // 真实姓名
	Mobile              string // 手机号
	Email               string // 电子邮箱
	Qq                  string // QQ
	Skype               string // Skype
	BankName            string // 银行名称
	CardAccount         string // 银行户名
	CardNo              string // 卡号
	DepositBankProvince string // 开户行所在省
	DepositBankCity     string // 开户行所在市
	DepositBank         string // 开户行
	RegisterIp          string //
	RegisterTime        string //
	LastLoginIp         string //
	LastLoginTime       string //
	CreatedAt           string //
	UpdatedAt           string //
	Reason              string // 申请理由
	Agentcode           string //
}

// agentColumns holds the columns for the table agent.
var agentColumns = AgentColumns{
	Id:                  "id",
	SiteId:              "site_id",
	Username:            "username",
	Password:            "password",
	Status:              "status",
	StatusCheck:         "status_check",
	Level:               "level",
	Realname:            "realname",
	Mobile:              "mobile",
	Email:               "email",
	Qq:                  "qq",
	Skype:               "skype",
	BankName:            "bank_name",
	CardAccount:         "card_account",
	CardNo:              "card_no",
	DepositBankProvince: "deposit_bank_province",
	DepositBankCity:     "deposit_bank_city",
	DepositBank:         "deposit_bank",
	RegisterIp:          "register_ip",
	RegisterTime:        "register_time",
	LastLoginIp:         "last_login_ip",
	LastLoginTime:       "last_login_time",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
	Reason:              "reason",
	Agentcode:           "agentcode",
}

// NewAgentDao creates and returns a new DAO object for table data access.
func NewAgentDao(handlers ...gdb.ModelHandler) *AgentDao {
	return &AgentDao{
		group:    "default",
		table:    "agent",
		columns:  agentColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AgentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AgentDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AgentDao) Columns() AgentColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AgentDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AgentDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AgentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
