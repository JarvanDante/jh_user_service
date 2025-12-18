// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Activity is the golang structure for table activity.
type Activity struct {
	Id               uint        `json:"id"               orm:"id"                 description:"活动ID"`
	SiteId           int         `json:"siteId"           orm:"site_id"            description:"站点ID"`
	ActivityName     string      `json:"activityName"     orm:"activity_name"      description:"活动名称"`
	Describe         string      `json:"describe"         orm:"describe"           description:"活动描述"`
	PcCover          string      `json:"pcCover"          orm:"pc_cover"           description:"pc端封面"`
	WapCover         string      `json:"wapCover"         orm:"wap_cover"          description:"wap端封面"`
	StartTime        *gtime.Time `json:"startTime"        orm:"start_time"         description:"开始时间"`
	EndTime          *gtime.Time `json:"endTime"          orm:"end_time"           description:"结束时间"`
	Status           int         `json:"status"           orm:"status"             description:"活动状态。1=开启；0=关闭"`
	IsRelateGame     int         `json:"isRelateGame"     orm:"is_relate_game"     description:"是否关联游戏"`
	RelateGame       string      `json:"relateGame"       orm:"relate_game"        description:"关联游戏"`
	TemplateId       int         `json:"templateId"       orm:"template_id"        description:"活动模板ID"`
	ActivityDetails  string      `json:"activityDetails"  orm:"activity_details"   description:"活动详情"`
	ActivityModuleId string      `json:"activityModuleId" orm:"activity_module_id" description:"活动模块ID"`
	IsPublish        int         `json:"isPublish"        orm:"is_publish"         description:"发布状态。1=发布；0=未发布"`
	Sort             int         `json:"sort"             orm:"sort"               description:"排序"`
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"         description:""`
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"         description:""`
	ActivityType     int         `json:"activityType"     orm:"activity_type"      description:"活动类型"`
	Name             string      `json:"name"             orm:"name"               description:"名称"`
	MobileCover      string      `json:"mobileCover"      orm:"mobile_cover"       description:"手机端封面"`
	Intro            string      `json:"intro"            orm:"intro"              description:""`
}
