// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SitePointsDao is the data access object for the table site_points.
type SitePointsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SitePointsColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SitePointsColumns defines and stores column names for the table site_points.
type SitePointsColumns struct {
	Id                   string //
	SiteId               string //
	SwitchSitePoints     string //
	SwitchRechargePoints string //
	EachRechargeAmount   string //
	EachRechargePoints   string //
	SwitchBettingPoints  string //
	EachBettingAmount    string //
	EachBettingPoints    string //
	MaxDailyPoints       string //
	CreatedAt            string //
	UpdatedAt            string //
}

// sitePointsColumns holds the columns for the table site_points.
var sitePointsColumns = SitePointsColumns{
	Id:                   "id",
	SiteId:               "site_id",
	SwitchSitePoints:     "switch_site_points",
	SwitchRechargePoints: "switch_recharge_points",
	EachRechargeAmount:   "each_recharge_amount",
	EachRechargePoints:   "each_recharge_points",
	SwitchBettingPoints:  "switch_betting_points",
	EachBettingAmount:    "each_betting_amount",
	EachBettingPoints:    "each_betting_points",
	MaxDailyPoints:       "max_daily_points",
	CreatedAt:            "created_at",
	UpdatedAt:            "updated_at",
}

// NewSitePointsDao creates and returns a new DAO object for table data access.
func NewSitePointsDao(handlers ...gdb.ModelHandler) *SitePointsDao {
	return &SitePointsDao{
		group:    "default",
		table:    "site_points",
		columns:  sitePointsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SitePointsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SitePointsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SitePointsDao) Columns() SitePointsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SitePointsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SitePointsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SitePointsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
