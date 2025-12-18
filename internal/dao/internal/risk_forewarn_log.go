// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RiskForewarnLogDao is the data access object for the table risk_forewarn_log.
type RiskForewarnLogDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  RiskForewarnLogColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// RiskForewarnLogColumns defines and stores column names for the table risk_forewarn_log.
type RiskForewarnLogColumns struct {
	Id        string //
	SiteId    string // 站点ID
	UserId    string // 会员ID
	Username  string // 用户名
	Content   string // 预警提示内容
	Ip        string // 用户IP
	Device    string // 终端。1=电脑；2=手机；3=平板
	Status    string //
	CreatedAt string //
	UpdatedAt string //
}

// riskForewarnLogColumns holds the columns for the table risk_forewarn_log.
var riskForewarnLogColumns = RiskForewarnLogColumns{
	Id:        "id",
	SiteId:    "site_id",
	UserId:    "user_id",
	Username:  "username",
	Content:   "content",
	Ip:        "ip",
	Device:    "device",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewRiskForewarnLogDao creates and returns a new DAO object for table data access.
func NewRiskForewarnLogDao(handlers ...gdb.ModelHandler) *RiskForewarnLogDao {
	return &RiskForewarnLogDao{
		group:    "default",
		table:    "risk_forewarn_log",
		columns:  riskForewarnLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RiskForewarnLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RiskForewarnLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RiskForewarnLogDao) Columns() RiskForewarnLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RiskForewarnLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RiskForewarnLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *RiskForewarnLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
