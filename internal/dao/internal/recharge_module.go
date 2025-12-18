// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RechargeModuleDao is the data access object for the table recharge_module.
type RechargeModuleDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  RechargeModuleColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// RechargeModuleColumns defines and stores column names for the table recharge_module.
type RechargeModuleColumns struct {
	Id                     string //
	SiteId                 string // 站点ID
	ActivityId             string // 活动id
	Name                   string // 活动名称
	StartTime              string // 开始时间
	EndTime                string // 结束时间
	Status                 string // 活动状态。1=开启；0=关闭
	UserGradeIds           string // 会员等级。以,的形式隔开
	UserLevelIds           string // 会员层级。以,的形式隔开
	Platform               string // 优惠终端。0=所有；1=网站；2=手机
	PaymentGateways        string // 在线支付网关。以,的形式隔开
	TransferTypes          string // 转账汇款类型。以,的形式隔开
	UserType               string // 面向用户类型。0=所有用户，1=新用户；2=老用户
	ConditionRule          string // 充值条件。1=当日单笔充值，2=本周单笔充值；3=本月单笔充值；4=会员首次充值
	ConditionCompare       string // 充值条件。1=大于等于，2=等于
	ConditionValue         string // 充值条件。值
	BonusType              string // 奖励金额类型。1=奖励实际金额；0=奖励金额比例
	BonusAmount            string // 奖励金额
	BonusPercent           string // 奖励金额比例
	BonusMax               string // 最高奖励金额.0=不限制
	ApplyFrequency         string // 申请频率。0=不限；1=每日；2=每周；3=每月
	ApplyTimes             string // 每人申请次数
	WithdrawNeedWaterTimes string // 提现流水要求多少倍流水
	CreatedAt              string //
	UpdatedAt              string //
}

// rechargeModuleColumns holds the columns for the table recharge_module.
var rechargeModuleColumns = RechargeModuleColumns{
	Id:                     "id",
	SiteId:                 "site_id",
	ActivityId:             "activity_id",
	Name:                   "name",
	StartTime:              "start_time",
	EndTime:                "end_time",
	Status:                 "status",
	UserGradeIds:           "user_grade_ids",
	UserLevelIds:           "user_level_ids",
	Platform:               "platform",
	PaymentGateways:        "payment_gateways",
	TransferTypes:          "transfer_types",
	UserType:               "user_type",
	ConditionRule:          "condition_rule",
	ConditionCompare:       "condition_compare",
	ConditionValue:         "condition_value",
	BonusType:              "bonus_type",
	BonusAmount:            "bonus_amount",
	BonusPercent:           "bonus_percent",
	BonusMax:               "bonus_max",
	ApplyFrequency:         "apply_frequency",
	ApplyTimes:             "apply_times",
	WithdrawNeedWaterTimes: "withdraw_need_water_times",
	CreatedAt:              "created_at",
	UpdatedAt:              "updated_at",
}

// NewRechargeModuleDao creates and returns a new DAO object for table data access.
func NewRechargeModuleDao(handlers ...gdb.ModelHandler) *RechargeModuleDao {
	return &RechargeModuleDao{
		group:    "default",
		table:    "recharge_module",
		columns:  rechargeModuleColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RechargeModuleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RechargeModuleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RechargeModuleDao) Columns() RechargeModuleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RechargeModuleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RechargeModuleDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *RechargeModuleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
