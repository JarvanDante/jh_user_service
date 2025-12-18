// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserDailyReport is the golang structure of table user_daily_report for DAO operations like Where/Data.
type UserDailyReport struct {
	g.Meta               `orm:"table:user_daily_report, do:true"`
	Id                   any         //
	SiteId               any         // 站点ID
	UserId               any         // 会员ID
	AgentId              any         // 代理ID
	Username             any         // 用户名
	ReportDate           *gtime.Time // 报表日期
	BetCount             any         // 注单数量
	BetAmount            any         // 投注总额
	ValidBetAmount       any         // 有效投注金额（投注时间）
	ValidBetAmountSettle any         // 有效投注金额（结算时间）
	WinOrLose            any         // 输赢结果
	RechargeAmount       any         // 充值金额
	WithdrawAmount       any         // 提现金额
	BonusAmount          any         // 红利金额
	RebateAmount         any         // 返水金额
	FeeAmount            any         // 手续费
	CreatedAt            *gtime.Time //
	UpdatedAt            *gtime.Time //
}
