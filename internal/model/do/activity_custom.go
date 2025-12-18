// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ActivityCustom is the golang structure of table activity_custom for DAO operations like Where/Data.
type ActivityCustom struct {
	g.Meta        `orm:"table:activity_custom, do:true"`
	Id            any         //
	SiteId        any         // 站点ID
	ActivityId    any         // 活动ID
	PcContent     any         // pc端活动内容
	MobileContent any         // 手机端活动内容
	PcLink        any         // pc端活动内容链接
	MobileLink    any         // 手机端活动内容链接
	PcType        any         // 内容类型: 1=pc端编辑框内容 2=pc端内容链接
	MobileType    any         // 内容类型: 1=手机端编辑框内容 2=手机端内容链接
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
}
