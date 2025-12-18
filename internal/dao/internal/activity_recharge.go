// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityRechargeDao is the data access object for the table activity_recharge.
type ActivityRechargeDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  ActivityRechargeColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// ActivityRechargeColumns defines and stores column names for the table activity_recharge.
type ActivityRechargeColumns struct {
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

// activityRechargeColumns holds the columns for the table activity_recharge.
var activityRechargeColumns = ActivityRechargeColumns{
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

// NewActivityRechargeDao creates and returns a new DAO object for table data access.
func NewActivityRechargeDao(handlers ...gdb.ModelHandler) *ActivityRechargeDao {
	return &ActivityRechargeDao{
		group:    "default",
		table:    "activity_recharge",
		columns:  activityRechargeColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ActivityRechargeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ActivityRechargeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ActivityRechargeDao) Columns() ActivityRechargeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ActivityRechargeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ActivityRechargeDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ActivityRechargeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
