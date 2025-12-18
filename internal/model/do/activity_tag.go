// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ActivityTag is the golang structure of table activity_tag for DAO operations like Where/Data.
type ActivityTag struct {
	g.Meta     `orm:"table:activity_tag, do:true"`
	Id         any         //
	SiteId     any         // 站点ID
	ActivityId any         // 活动ID
	Name       any         // 标签名称
	Status     any         // 状态。1=启用；0=禁用；-1=删除
	Sort       any         // 排序。值越小排名越靠前
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
	TagId      any         //
}
