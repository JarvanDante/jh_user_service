// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLevel is the golang structure of table user_level for DAO operations like Where/Data.
type UserLevel struct {
	g.Meta             `orm:"table:user_level, do:true"`
	Id                 any         //
	SiteId             any         // 站点ID
	Name               any         // 层级名称
	IsRebate           any         // 是否返水.1=返水；0=不返水
	RebateRuleId       any         // 返水规则ID
	DailyWithdrawTimes any         // 单日提款次数上限
	LoginUrl           any         // 专用登录网址
	Status             any         // 状态.1=可用；0=禁用
	CreatedAt          *gtime.Time //
	UpdatedAt          *gtime.Time //
}
