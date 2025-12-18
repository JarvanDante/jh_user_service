// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityModuleDao is the data access object for the table activity_module.
type ActivityModuleDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  ActivityModuleColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// ActivityModuleColumns defines and stores column names for the table activity_module.
type ActivityModuleColumns struct {
	Id         string //
	SiteId     string // 站点ID
	ActivityId string // 活动ID
	ModuleType string // 活动模块类型。1=充值；2=大转盘；3=开宝箱；4=砸金蛋；5=抢红包
	ModuleId   string // 活动模块ID
	ModuleName string // 活动模块名称
	StartTime  string // 开始时间
	EndTime    string // 结束时间
	Status     string // 活动状态。1=开启；0=关闭
	CreatedAt  string //
	UpdatedAt  string //
}

// activityModuleColumns holds the columns for the table activity_module.
var activityModuleColumns = ActivityModuleColumns{
	Id:         "id",
	SiteId:     "site_id",
	ActivityId: "activity_id",
	ModuleType: "module_type",
	ModuleId:   "module_id",
	ModuleName: "module_name",
	StartTime:  "start_time",
	EndTime:    "end_time",
	Status:     "status",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewActivityModuleDao creates and returns a new DAO object for table data access.
func NewActivityModuleDao(handlers ...gdb.ModelHandler) *ActivityModuleDao {
	return &ActivityModuleDao{
		group:    "default",
		table:    "activity_module",
		columns:  activityModuleColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ActivityModuleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ActivityModuleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ActivityModuleDao) Columns() ActivityModuleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ActivityModuleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ActivityModuleDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ActivityModuleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
