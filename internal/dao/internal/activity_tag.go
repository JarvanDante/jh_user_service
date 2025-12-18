// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityTagDao is the data access object for the table activity_tag.
type ActivityTagDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ActivityTagColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ActivityTagColumns defines and stores column names for the table activity_tag.
type ActivityTagColumns struct {
	Id         string //
	SiteId     string // 站点ID
	ActivityId string // 活动ID
	Name       string // 标签名称
	Status     string // 状态。1=启用；0=禁用；-1=删除
	Sort       string // 排序。值越小排名越靠前
	CreatedAt  string //
	UpdatedAt  string //
	TagId      string //
}

// activityTagColumns holds the columns for the table activity_tag.
var activityTagColumns = ActivityTagColumns{
	Id:         "id",
	SiteId:     "site_id",
	ActivityId: "activity_id",
	Name:       "name",
	Status:     "status",
	Sort:       "sort",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	TagId:      "tag_id",
}

// NewActivityTagDao creates and returns a new DAO object for table data access.
func NewActivityTagDao(handlers ...gdb.ModelHandler) *ActivityTagDao {
	return &ActivityTagDao{
		group:    "default",
		table:    "activity_tag",
		columns:  activityTagColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ActivityTagDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ActivityTagDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ActivityTagDao) Columns() ActivityTagColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ActivityTagDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ActivityTagDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ActivityTagDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
