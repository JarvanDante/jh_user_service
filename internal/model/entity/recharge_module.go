// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RechargeModule is the golang structure for table recharge_module.
type RechargeModule struct {
	Id                     uint        `json:"id"                     orm:"id"                        description:""`
	SiteId                 int         `json:"siteId"                 orm:"site_id"                   description:"站点ID"`
	ActivityId             int         `json:"activityId"             orm:"activity_id"               description:"活动id"`
	Name                   string      `json:"name"                   orm:"name"                      description:"活动名称"`
	StartTime              *gtime.Time `json:"startTime"              orm:"start_time"                description:"开始时间"`
	EndTime                *gtime.Time `json:"endTime"                orm:"end_time"                  description:"结束时间"`
	Status                 int         `json:"status"                 orm:"status"                    description:"活动状态。1=开启；0=关闭"`
	UserGradeIds           string      `json:"userGradeIds"           orm:"user_grade_ids"            description:"会员等级。以,的形式隔开"`
	UserLevelIds           string      `json:"userLevelIds"           orm:"user_level_ids"            description:"会员层级。以,的形式隔开"`
	Platform               int         `json:"platform"               orm:"platform"                  description:"优惠终端。0=所有；1=网站；2=手机"`
	PaymentGateways        string      `json:"paymentGateways"        orm:"payment_gateways"          description:"在线支付网关。以,的形式隔开"`
	TransferTypes          string      `json:"transferTypes"          orm:"transfer_types"            description:"转账汇款类型。以,的形式隔开"`
	UserType               int         `json:"userType"               orm:"user_type"                 description:"面向用户类型。0=所有用户，1=新用户；2=老用户"`
	ConditionRule          int         `json:"conditionRule"          orm:"condition_rule"            description:"充值条件。1=当日单笔充值，2=本周单笔充值；3=本月单笔充值；4=会员首次充值"`
	ConditionCompare       int         `json:"conditionCompare"       orm:"condition_compare"         description:"充值条件。1=大于等于，2=等于"`
	ConditionValue         float64     `json:"conditionValue"         orm:"condition_value"           description:"充值条件。值"`
	BonusType              int         `json:"bonusType"              orm:"bonus_type"                description:"奖励金额类型。1=奖励实际金额；0=奖励金额比例"`
	BonusAmount            float64     `json:"bonusAmount"            orm:"bonus_amount"              description:"奖励金额"`
	BonusPercent           float64     `json:"bonusPercent"           orm:"bonus_percent"             description:"奖励金额比例"`
	BonusMax               float64     `json:"bonusMax"               orm:"bonus_max"                 description:"最高奖励金额.0=不限制"`
	ApplyFrequency         int         `json:"applyFrequency"         orm:"apply_frequency"           description:"申请频率。0=不限；1=每日；2=每周；3=每月"`
	ApplyTimes             int         `json:"applyTimes"             orm:"apply_times"               description:"每人申请次数"`
	WithdrawNeedWaterTimes int         `json:"withdrawNeedWaterTimes" orm:"withdraw_need_water_times" description:"提现流水要求多少倍流水"`
	CreatedAt              *gtime.Time `json:"createdAt"              orm:"created_at"                description:""`
	UpdatedAt              *gtime.Time `json:"updatedAt"              orm:"updated_at"                description:""`
}
