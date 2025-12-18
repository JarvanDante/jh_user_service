// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RiskForewarnConfig is the golang structure of table risk_forewarn_config for DAO operations like Where/Data.
type RiskForewarnConfig struct {
	g.Meta                `orm:"table:risk_forewarn_config, do:true"`
	Id                    any         //
	SiteId                any         // 站点ID
	AccountToGame         any         // 账户转账到游戏单笔金额预警, 0=不提示 其他=提示
	GameToAccount         any         // 游戏转账到账户单笔金额预警, 0=不提示 其他=提示
	WinMoney              any         // 会员单笔游戏赢得金额预警,0=不提示 其他=提示
	BetAmount             any         // 会员单笔游戏投注金额预警,0=不提示 其他=提示
	AlterBankCard         any         // 会员银行卡号被修改预警，0=不提示；1=提示
	LoginAreaDifference   any         // 本次和上次登录异样提示，0=不提示；1=提示
	SameipRegisterTime    any         // 同一IP单位时间内注册限制，0=不限时间 其他=时间，单位：分钟
	SameipRegisterCount   any         // 同一IP单位时间内注册次数限制，0=不限 其他=提示
	CreatedAt             *gtime.Time //
	UpdatedAt             *gtime.Time //
	IsAccountToGame       any         // 是否开启会员账户-游戏预警
	IsGameToAccount       any         // 是否开启会员游戏-账户预警
	IsWinMoney            any         // 是否开启游戏输赢预警
	IsBetAmount           any         // 是否开启游戏投注预警
	IsAlterBankCard       any         // 是否开启修改银行卡预警
	IsLoginAreaDifference any         // 是否开启登录地区不同预警
	IsSameipRegisterCount any         // 是否开启相同IP注册次数预警
}
