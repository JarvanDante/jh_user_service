// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BetLogDaily is the golang structure for table bet_log_daily.
type BetLogDaily struct {
	Id                   uint        `json:"id"                   orm:"id"                      description:""`
	SiteId               uint        `json:"siteId"               orm:"site_id"                 description:"站点ID"`
	BetDate              *gtime.Time `json:"betDate"              orm:"bet_date"                description:"投注日期"`
	GameId               uint        `json:"gameId"               orm:"game_id"                 description:"游戏ID"`
	AgentId              int         `json:"agentId"              orm:"agent_id"                description:"代理ID"`
	GameType             uint        `json:"gameType"             orm:"game_type"               description:"游戏类型"`
	UserId               uint        `json:"userId"               orm:"user_id"                 description:"会员ID"`
	Username             string      `json:"username"             orm:"username"                description:"会员用户名"`
	BetAmount            float64     `json:"betAmount"            orm:"bet_amount"              description:"投注金额"`
	BetCount             uint        `json:"betCount"             orm:"bet_count"               description:"注单数量"`
	ValidBetAmount       float64     `json:"validBetAmount"       orm:"valid_bet_amount"        description:"有效投注金额有效投注金额（投注时间）"`
	ValidBetAmountSettle float64     `json:"validBetAmountSettle" orm:"valid_bet_amount_settle" description:"有效投注金额（结算时间）"`
	WinOrLose            float64     `json:"winOrLose"            orm:"win_or_lose"             description:"输赢结果"`
	CreatedAt            *gtime.Time `json:"createdAt"            orm:"created_at"              description:""`
	UpdatedAt            *gtime.Time `json:"updatedAt"            orm:"updated_at"              description:""`
}
