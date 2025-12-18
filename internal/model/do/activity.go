// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Activity is the golang structure of table activity for DAO operations like Where/Data.
type Activity struct {
	g.Meta           `orm:"table:activity, do:true"`
	Id               any         // 活动ID
	SiteId           any         // 站点ID
	ActivityName     any         // 活动名称
	Describe         any         // 活动描述
	PcCover          any         // pc端封面
	WapCover         any         // wap端封面
	StartTime        *gtime.Time // 开始时间
	EndTime          *gtime.Time // 结束时间
	Status           any         // 活动状态。1=开启；0=关闭
	IsRelateGame     any         // 是否关联游戏
	RelateGame       any         // 关联游戏
	TemplateId       any         // 活动模板ID
	ActivityDetails  any         // 活动详情
	ActivityModuleId any         // 活动模块ID
	IsPublish        any         // 发布状态。1=发布；0=未发布
	Sort             any         // 排序
	CreatedAt        *gtime.Time //
	UpdatedAt        *gtime.Time //
	ActivityType     any         // 活动类型
	Name             any         // 名称
	MobileCover      any         // 手机端封面
	Intro            any         //
}
