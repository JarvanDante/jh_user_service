// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AgentSharingLog is the golang structure of table agent_sharing_log for DAO operations like Where/Data.
type AgentSharingLog struct {
	g.Meta                `orm:"table:agent_sharing_log, do:true"`
	Id                    any         //
	SiteId                any         // 站点ID
	StartDate             *gtime.Time // 开始日期
	EndDate               *gtime.Time // 结束日期
	AgentId               any         // 代理ID
	AgentUsername         any         // 代理账号
	UserCount             any         // 有效会员数
	RechargeAmount        any         // 充值总计
	WithdrawAmount        any         // 提现总计
	BonusAmount           any         // 红利总计
	RebateAmount          any         // 返水总计
	ValidBetAmount        any         // 有效投注总计
	WinOrLose             any         // 输赢总计
	FeeAmount             any         // 手续费总计
	AdministrationExpense any         // 行政费用
	CalculateCommission   any         // 程序计算出的代理佣金
	Commission            any         // 发放佣金
	Status                any         // 状态。1=未结算；2=结算成功；3=结算失败；4=未盈利
	AdminId               any         // 后台员工ID
	CreatedAt             *gtime.Time //
	UpdatedAt             *gtime.Time //
	CommissionData        any         // 返佣金额计算公式数据
}
