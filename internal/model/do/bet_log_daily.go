// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BetLogDaily is the golang structure of table bet_log_daily for DAO operations like Where/Data.
type BetLogDaily struct {
	g.Meta               `orm:"table:bet_log_daily, do:true"`
	Id                   any         //
	SiteId               any         // 站点ID
	BetDate              *gtime.Time // 投注日期
	GameId               any         // 游戏ID
	AgentId              any         // 代理ID
	GameType             any         // 游戏类型
	UserId               any         // 会员ID
	Username             any         // 会员用户名
	BetAmount            any         // 投注金额
	BetCount             any         // 注单数量
	ValidBetAmount       any         // 有效投注金额有效投注金额（投注时间）
	ValidBetAmountSettle any         // 有效投注金额（结算时间）
	WinOrLose            any         // 输赢结果
	CreatedAt            *gtime.Time //
	UpdatedAt            *gtime.Time //
}
