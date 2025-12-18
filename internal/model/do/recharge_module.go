// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RechargeModule is the golang structure of table recharge_module for DAO operations like Where/Data.
type RechargeModule struct {
	g.Meta                 `orm:"table:recharge_module, do:true"`
	Id                     any         //
	SiteId                 any         // 站点ID
	ActivityId             any         // 活动id
	Name                   any         // 活动名称
	StartTime              *gtime.Time // 开始时间
	EndTime                *gtime.Time // 结束时间
	Status                 any         // 活动状态。1=开启；0=关闭
	UserGradeIds           any         // 会员等级。以,的形式隔开
	UserLevelIds           any         // 会员层级。以,的形式隔开
	Platform               any         // 优惠终端。0=所有；1=网站；2=手机
	PaymentGateways        any         // 在线支付网关。以,的形式隔开
	TransferTypes          any         // 转账汇款类型。以,的形式隔开
	UserType               any         // 面向用户类型。0=所有用户，1=新用户；2=老用户
	ConditionRule          any         // 充值条件。1=当日单笔充值，2=本周单笔充值；3=本月单笔充值；4=会员首次充值
	ConditionCompare       any         // 充值条件。1=大于等于，2=等于
	ConditionValue         any         // 充值条件。值
	BonusType              any         // 奖励金额类型。1=奖励实际金额；0=奖励金额比例
	BonusAmount            any         // 奖励金额
	BonusPercent           any         // 奖励金额比例
	BonusMax               any         // 最高奖励金额.0=不限制
	ApplyFrequency         any         // 申请频率。0=不限；1=每日；2=每周；3=每月
	ApplyTimes             any         // 每人申请次数
	WithdrawNeedWaterTimes any         // 提现流水要求多少倍流水
	CreatedAt              *gtime.Time //
	UpdatedAt              *gtime.Time //
}
