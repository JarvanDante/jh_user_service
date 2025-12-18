// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RebateRuleOptionGame is the golang structure of table rebate_rule_option_game for DAO operations like Where/Data.
type RebateRuleOptionGame struct {
	g.Meta       `orm:"table:rebate_rule_option_game, do:true"`
	Id           any         //
	SiteId       any         // 站点ID
	RuleId       any         // 返水规则ID
	RuleOptionId any         // 活动ID
	GameId       any         // 游戏ID
	Percent      any         // 返水比例
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}
