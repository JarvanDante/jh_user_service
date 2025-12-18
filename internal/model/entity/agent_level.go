// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AgentLevel is the golang structure for table agent_level.
type AgentLevel struct {
	Id                  uint        `json:"id"                  orm:"id"                    description:""`
	SiteId              int         `json:"siteId"              orm:"site_id"               description:""`
	Name                string      `json:"name"                orm:"name"                  description:""`
	Status              int         `json:"status"              orm:"status"                description:""`
	AgentRegisterBonus  float64     `json:"agentRegisterBonus"  orm:"agent_register_bonus"  description:"代理注册奖励金"`
	SwitchSharing       int         `json:"switchSharing"       orm:"switch_sharing"        description:"分成模式开关。1=开启；0=关闭"`
	SwitchRebate        int         `json:"switchRebate"        orm:"switch_rebate"         description:"返点模式开关。1=开启；0=关闭"`
	PercentDeductBonus  float64     `json:"percentDeductBonus"  orm:"percent_deduct_bonus"  description:"红利扣除比例。以小数记录百分比"`
	PercentDeductRebate float64     `json:"percentDeductRebate" orm:"percent_deduct_rebate" description:"返水扣除比例。以小数记录百分比"`
	CreatedAt           *gtime.Time `json:"createdAt"           orm:"created_at"            description:""`
	UpdatedAt           *gtime.Time `json:"updatedAt"           orm:"updated_at"            description:""`
	SharingOptions      string      `json:"sharingOptions"      orm:"sharing_options"       description:"分成模式条件列表。JSON格式。month_amount_min：团队当月盈利最小，month_amount_max：团队当月盈利最大，valid_user_count：有效人数，valid_bet_count：有效投注，percent_sports：分成比例 - 体育，percent_lottery：分成比例 - 彩票，percent_live：分成比例 - 真人，percent_egame：分成比例 - 电子游戏。"`
	RebateOptions       string      `json:"rebateOptions"       orm:"rebate_options"        description:"返点模式条件列表。JSON格式。valid_user_count：有效人数，valid_bet_count：有效投注，percent_sports：分成比例 - 体育，percent_lottery：分成比例 - 彩票，percent_live：分成比例 - 真人，percent_egame：分成比例 - 电子游戏。"`
}
