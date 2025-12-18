// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AgentMessageLogDao is the data access object for the table agent_message_log.
type AgentMessageLogDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  AgentMessageLogColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// AgentMessageLogColumns defines and stores column names for the table agent_message_log.
type AgentMessageLogColumns struct {
	Id             string //
	SiteId         string // 站点ID
	AgentMessageId string // 代理消息id
	AgentId        string // 代理id
	IsRead         string // 是否读过
	CreatedAt      string //
	UpdatedAt      string //
}

// agentMessageLogColumns holds the columns for the table agent_message_log.
var agentMessageLogColumns = AgentMessageLogColumns{
	Id:             "id",
	SiteId:         "site_id",
	AgentMessageId: "agent_message_id",
	AgentId:        "agent_id",
	IsRead:         "is_read",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewAgentMessageLogDao creates and returns a new DAO object for table data access.
func NewAgentMessageLogDao(handlers ...gdb.ModelHandler) *AgentMessageLogDao {
	return &AgentMessageLogDao{
		group:    "default",
		table:    "agent_message_log",
		columns:  agentMessageLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AgentMessageLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AgentMessageLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AgentMessageLogDao) Columns() AgentMessageLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AgentMessageLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AgentMessageLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AgentMessageLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
