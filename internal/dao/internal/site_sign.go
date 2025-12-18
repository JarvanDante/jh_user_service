// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SiteSignDao is the data access object for the table site_sign.
type SiteSignDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SiteSignColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SiteSignColumns defines and stores column names for the table site_sign.
type SiteSignColumns struct {
	Id        string //
	SiteId    string // 站点ID
	Name      string // 活动名称
	StartTime string // 开始时间
	EndTime   string // 结束时间
	Status    string // 活动状态。1=开启；0=关闭
	UserGrade string // 会员等级。以,的形式隔开
	UserLevel string // 会员层级。以,的形式隔开
	Platform  string // 优惠终端。0=所有；1=网站；2=手机
	Remark    string // 活动描述
	CreatedAt string //
	UpdatedAt string //
}

// siteSignColumns holds the columns for the table site_sign.
var siteSignColumns = SiteSignColumns{
	Id:        "id",
	SiteId:    "site_id",
	Name:      "name",
	StartTime: "start_time",
	EndTime:   "end_time",
	Status:    "status",
	UserGrade: "user_grade",
	UserLevel: "user_level",
	Platform:  "platform",
	Remark:    "remark",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSiteSignDao creates and returns a new DAO object for table data access.
func NewSiteSignDao(handlers ...gdb.ModelHandler) *SiteSignDao {
	return &SiteSignDao{
		group:    "default",
		table:    "site_sign",
		columns:  siteSignColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SiteSignDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SiteSignDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SiteSignDao) Columns() SiteSignColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SiteSignDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SiteSignDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SiteSignDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
