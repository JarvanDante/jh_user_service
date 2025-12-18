// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserDailyReport is the golang structure for table user_daily_report.
type UserDailyReport struct {
	Id                   uint64      `json:"id"                   orm:"id"                      description:""`
	SiteId               int         `json:"siteId"               orm:"site_id"                 description:"站点ID"`
	UserId               int         `json:"userId"               orm:"user_id"                 description:"会员ID"`
	AgentId              int         `json:"agentId"              orm:"agent_id"                description:"代理ID"`
	Username             string      `json:"username"             orm:"username"                description:"用户名"`
	ReportDate           *gtime.Time `json:"reportDate"           orm:"report_date"             description:"报表日期"`
	BetCount             int         `json:"betCount"             orm:"bet_count"               description:"注单数量"`
	BetAmount            float64     `json:"betAmount"            orm:"bet_amount"              description:"投注总额"`
	ValidBetAmount       float64     `json:"validBetAmount"       orm:"valid_bet_amount"        description:"有效投注金额（投注时间）"`
	ValidBetAmountSettle float64     `json:"validBetAmountSettle" orm:"valid_bet_amount_settle" description:"有效投注金额（结算时间）"`
	WinOrLose            float64     `json:"winOrLose"            orm:"win_or_lose"             description:"输赢结果"`
	RechargeAmount       float64     `json:"rechargeAmount"       orm:"recharge_amount"         description:"充值金额"`
	WithdrawAmount       float64     `json:"withdrawAmount"       orm:"withdraw_amount"         description:"提现金额"`
	BonusAmount          float64     `json:"bonusAmount"          orm:"bonus_amount"            description:"红利金额"`
	RebateAmount         float64     `json:"rebateAmount"         orm:"rebate_amount"           description:"返水金额"`
	FeeAmount            float64     `json:"feeAmount"            orm:"fee_amount"              description:"手续费"`
	CreatedAt            *gtime.Time `json:"createdAt"            orm:"created_at"              description:""`
	UpdatedAt            *gtime.Time `json:"updatedAt"            orm:"updated_at"              description:""`
}
