// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CustomTagDao is the data access object for the table custom_tag.
type CustomTagDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  CustomTagColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// CustomTagColumns defines and stores column names for the table custom_tag.
type CustomTagColumns struct {
	Id         string //
	SiteId     string //
	Type       string // 类型 1=充值 2=提现
	Status     string // 1=正常，0=禁用
	TagName    string // 标签名称
	TagContent string // 标签内容
	CreatedAt  string //
	UpdatedAt  string //
}

// customTagColumns holds the columns for the table custom_tag.
var customTagColumns = CustomTagColumns{
	Id:         "id",
	SiteId:     "site_id",
	Type:       "type",
	Status:     "status",
	TagName:    "tag_name",
	TagContent: "tag_content",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewCustomTagDao creates and returns a new DAO object for table data access.
func NewCustomTagDao(handlers ...gdb.ModelHandler) *CustomTagDao {
	return &CustomTagDao{
		group:    "default",
		table:    "custom_tag",
		columns:  customTagColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CustomTagDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CustomTagDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CustomTagDao) Columns() CustomTagColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CustomTagDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CustomTagDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *CustomTagDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
