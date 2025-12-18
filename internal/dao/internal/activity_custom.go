// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityCustomDao is the data access object for the table activity_custom.
type ActivityCustomDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  ActivityCustomColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// ActivityCustomColumns defines and stores column names for the table activity_custom.
type ActivityCustomColumns struct {
	Id            string //
	SiteId        string // 站点ID
	ActivityId    string // 活动ID
	PcContent     string // pc端活动内容
	MobileContent string // 手机端活动内容
	PcLink        string // pc端活动内容链接
	MobileLink    string // 手机端活动内容链接
	PcType        string // 内容类型: 1=pc端编辑框内容 2=pc端内容链接
	MobileType    string // 内容类型: 1=手机端编辑框内容 2=手机端内容链接
	CreatedAt     string //
	UpdatedAt     string //
}

// activityCustomColumns holds the columns for the table activity_custom.
var activityCustomColumns = ActivityCustomColumns{
	Id:            "id",
	SiteId:        "site_id",
	ActivityId:    "activity_id",
	PcContent:     "pc_content",
	MobileContent: "mobile_content",
	PcLink:        "pc_link",
	MobileLink:    "mobile_link",
	PcType:        "pc_type",
	MobileType:    "mobile_type",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewActivityCustomDao creates and returns a new DAO object for table data access.
func NewActivityCustomDao(handlers ...gdb.ModelHandler) *ActivityCustomDao {
	return &ActivityCustomDao{
		group:    "default",
		table:    "activity_custom",
		columns:  activityCustomColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ActivityCustomDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ActivityCustomDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ActivityCustomDao) Columns() ActivityCustomColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ActivityCustomDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ActivityCustomDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ActivityCustomDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
