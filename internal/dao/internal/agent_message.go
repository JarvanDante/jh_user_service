// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AgentMessageDao is the data access object for the table agent_message.
type AgentMessageDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  AgentMessageColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// AgentMessageColumns defines and stores column names for the table agent_message.
type AgentMessageColumns struct {
	Id        string //
	SiteId    string // 站点ID
	AdminId   string // 员工ID
	Title     string // 消息标题
	Content   string // 消息内容
	Receiver  string // 接收者
	Status    string // 状态。1=正常；0=禁用
	CreatedAt string //
	UpdatedAt string //
}

// agentMessageColumns holds the columns for the table agent_message.
var agentMessageColumns = AgentMessageColumns{
	Id:        "id",
	SiteId:    "site_id",
	AdminId:   "admin_id",
	Title:     "title",
	Content:   "content",
	Receiver:  "receiver",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewAgentMessageDao creates and returns a new DAO object for table data access.
func NewAgentMessageDao(handlers ...gdb.ModelHandler) *AgentMessageDao {
	return &AgentMessageDao{
		group:    "default",
		table:    "agent_message",
		columns:  agentMessageColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AgentMessageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AgentMessageDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AgentMessageDao) Columns() AgentMessageColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AgentMessageDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AgentMessageDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AgentMessageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
