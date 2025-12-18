// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AgentLevel is the golang structure of table agent_level for DAO operations like Where/Data.
type AgentLevel struct {
	g.Meta              `orm:"table:agent_level, do:true"`
	Id                  any         //
	SiteId              any         //
	Name                any         //
	Status              any         //
	AgentRegisterBonus  any         // 代理注册奖励金
	SwitchSharing       any         // 分成模式开关。1=开启；0=关闭
	SwitchRebate        any         // 返点模式开关。1=开启；0=关闭
	PercentDeductBonus  any         // 红利扣除比例。以小数记录百分比
	PercentDeductRebate any         // 返水扣除比例。以小数记录百分比
	CreatedAt           *gtime.Time //
	UpdatedAt           *gtime.Time //
	SharingOptions      any         // 分成模式条件列表。JSON格式。month_amount_min：团队当月盈利最小，month_amount_max：团队当月盈利最大，valid_user_count：有效人数，valid_bet_count：有效投注，percent_sports：分成比例 - 体育，percent_lottery：分成比例 - 彩票，percent_live：分成比例 - 真人，percent_egame：分成比例 - 电子游戏。
	RebateOptions       any         // 返点模式条件列表。JSON格式。valid_user_count：有效人数，valid_bet_count：有效投注，percent_sports：分成比例 - 体育，percent_lottery：分成比例 - 彩票，percent_live：分成比例 - 真人，percent_egame：分成比例 - 电子游戏。
}
