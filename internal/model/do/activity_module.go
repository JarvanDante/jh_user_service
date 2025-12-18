// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ActivityModule is the golang structure of table activity_module for DAO operations like Where/Data.
type ActivityModule struct {
	g.Meta     `orm:"table:activity_module, do:true"`
	Id         any         //
	SiteId     any         // 站点ID
	ActivityId any         // 活动ID
	ModuleType any         // 活动模块类型。1=充值；2=大转盘；3=开宝箱；4=砸金蛋；5=抢红包
	ModuleId   any         // 活动模块ID
	ModuleName any         // 活动模块名称
	StartTime  *gtime.Time // 开始时间
	EndTime    *gtime.Time // 结束时间
	Status     any         // 活动状态。1=开启；0=关闭
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
}
