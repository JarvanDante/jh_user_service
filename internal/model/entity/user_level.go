// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLevel is the golang structure for table user_level.
type UserLevel struct {
	Id                 uint        `json:"id"                 orm:"id"                   description:""`
	SiteId             int         `json:"siteId"             orm:"site_id"              description:"站点ID"`
	Name               string      `json:"name"               orm:"name"                 description:"层级名称"`
	IsRebate           int         `json:"isRebate"           orm:"is_rebate"            description:"是否返水.1=返水；0=不返水"`
	RebateRuleId       int         `json:"rebateRuleId"       orm:"rebate_rule_id"       description:"返水规则ID"`
	DailyWithdrawTimes int         `json:"dailyWithdrawTimes" orm:"daily_withdraw_times" description:"单日提款次数上限"`
	LoginUrl           string      `json:"loginUrl"           orm:"login_url"            description:"专用登录网址"`
	Status             int         `json:"status"             orm:"status"               description:"状态.1=可用；0=禁用"`
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"           description:""`
	UpdatedAt          *gtime.Time `json:"updatedAt"          orm:"updated_at"           description:""`
}
