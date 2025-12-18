// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MessageLogDao is the data access object for the table message_log.
type MessageLogDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  MessageLogColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// MessageLogColumns defines and stores column names for the table message_log.
type MessageLogColumns struct {
	Id          string //
	SiteId      string // 站点ID
	MessageId   string // 消息ID
	MessageType string // 发送类型。1=会员等级，2=会员层级，2=会员状态
	UserGrade   string // 会员等级
	UserLevel   string // 会员层级
	UserStatus  string // 会员状态.0=离线；1=在线
	CreatedAt   string //
	UpdatedAt   string //
}

// messageLogColumns holds the columns for the table message_log.
var messageLogColumns = MessageLogColumns{
	Id:          "id",
	SiteId:      "site_id",
	MessageId:   "message_id",
	MessageType: "message_type",
	UserGrade:   "user_grade",
	UserLevel:   "user_level",
	UserStatus:  "user_status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewMessageLogDao creates and returns a new DAO object for table data access.
func NewMessageLogDao(handlers ...gdb.ModelHandler) *MessageLogDao {
	return &MessageLogDao{
		group:    "default",
		table:    "message_log",
		columns:  messageLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MessageLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MessageLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MessageLogDao) Columns() MessageLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MessageLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MessageLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MessageLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
