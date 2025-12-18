// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserLevelPaymentDao is the data access object for the table user_level_payment.
type UserLevelPaymentDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  UserLevelPaymentColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// UserLevelPaymentColumns defines and stores column names for the table user_level_payment.
type UserLevelPaymentColumns struct {
	Id               string //
	SiteId           string // 站点ID
	UserLevelId      string // 会员层级ID
	Gateway          string // 支付类型
	PaymentAccountId string // 支付接口ID
	CreatedAt        string //
	UpdatedAt        string //
}

// userLevelPaymentColumns holds the columns for the table user_level_payment.
var userLevelPaymentColumns = UserLevelPaymentColumns{
	Id:               "id",
	SiteId:           "site_id",
	UserLevelId:      "user_level_id",
	Gateway:          "gateway",
	PaymentAccountId: "payment_account_id",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// NewUserLevelPaymentDao creates and returns a new DAO object for table data access.
func NewUserLevelPaymentDao(handlers ...gdb.ModelHandler) *UserLevelPaymentDao {
	return &UserLevelPaymentDao{
		group:    "default",
		table:    "user_level_payment",
		columns:  userLevelPaymentColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserLevelPaymentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserLevelPaymentDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserLevelPaymentDao) Columns() UserLevelPaymentColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserLevelPaymentDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserLevelPaymentDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserLevelPaymentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
