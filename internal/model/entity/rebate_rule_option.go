// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RebateRuleOption is the golang structure for table rebate_rule_option.
type RebateRuleOption struct {
	Id               uint        `json:"id"               orm:"id"                  description:""`
	SiteId           int         `json:"siteId"           orm:"site_id"             description:""`
	RuleId           int         `json:"ruleId"           orm:"rule_id"             description:"返水规则ID"`
	DailyValidBetMin float64     `json:"dailyValidBetMin" orm:"daily_valid_bet_min" description:"单日有效投注范围最小金额"`
	DailyValidBetMax float64     `json:"dailyValidBetMax" orm:"daily_valid_bet_max" description:"单日有效投注范围最大金额"`
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"          description:""`
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"          description:""`
}
