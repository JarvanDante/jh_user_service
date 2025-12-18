// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SiteSign is the golang structure of table site_sign for DAO operations like Where/Data.
type SiteSign struct {
	g.Meta    `orm:"table:site_sign, do:true"`
	Id        any         //
	SiteId    any         // 站点ID
	Name      any         // 活动名称
	StartTime *gtime.Time // 开始时间
	EndTime   *gtime.Time // 结束时间
	Status    any         // 活动状态。1=开启；0=关闭
	UserGrade any         // 会员等级。以,的形式隔开
	UserLevel any         // 会员层级。以,的形式隔开
	Platform  any         // 优惠终端。0=所有；1=网站；2=手机
	Remark    any         // 活动描述
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
