// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RebateRuleOptionGame is the golang structure for table rebate_rule_option_game.
type RebateRuleOptionGame struct {
	Id           uint        `json:"id"           orm:"id"             description:""`
	SiteId       int         `json:"siteId"       orm:"site_id"        description:"站点ID"`
	RuleId       int         `json:"ruleId"       orm:"rule_id"        description:"返水规则ID"`
	RuleOptionId int         `json:"ruleOptionId" orm:"rule_option_id" description:"活动ID"`
	GameId       int         `json:"gameId"       orm:"game_id"        description:"游戏ID"`
	Percent      float64     `json:"percent"      orm:"percent"        description:"返水比例"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"     description:""`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"     description:""`
}
