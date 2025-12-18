// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RiskForewarnLog is the golang structure of table risk_forewarn_log for DAO operations like Where/Data.
type RiskForewarnLog struct {
	g.Meta    `orm:"table:risk_forewarn_log, do:true"`
	Id        any         //
	SiteId    any         // 站点ID
	UserId    any         // 会员ID
	Username  any         // 用户名
	Content   any         // 预警提示内容
	Ip        any         // 用户IP
	Device    any         // 终端。1=电脑；2=手机；3=平板
	Status    any         //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
