// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RebateRuleOption is the golang structure of table rebate_rule_option for DAO operations like Where/Data.
type RebateRuleOption struct {
	g.Meta           `orm:"table:rebate_rule_option, do:true"`
	Id               any         //
	SiteId           any         //
	RuleId           any         // 返水规则ID
	DailyValidBetMin any         // 单日有效投注范围最小金额
	DailyValidBetMax any         // 单日有效投注范围最大金额
	CreatedAt        *gtime.Time //
	UpdatedAt        *gtime.Time //
}
