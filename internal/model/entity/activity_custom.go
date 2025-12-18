// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ActivityCustom is the golang structure for table activity_custom.
type ActivityCustom struct {
	Id            uint        `json:"id"            orm:"id"             description:""`
	SiteId        int         `json:"siteId"        orm:"site_id"        description:"站点ID"`
	ActivityId    int         `json:"activityId"    orm:"activity_id"    description:"活动ID"`
	PcContent     string      `json:"pcContent"     orm:"pc_content"     description:"pc端活动内容"`
	MobileContent string      `json:"mobileContent" orm:"mobile_content" description:"手机端活动内容"`
	PcLink        string      `json:"pcLink"        orm:"pc_link"        description:"pc端活动内容链接"`
	MobileLink    string      `json:"mobileLink"    orm:"mobile_link"    description:"手机端活动内容链接"`
	PcType        int         `json:"pcType"        orm:"pc_type"        description:"内容类型: 1=pc端编辑框内容 2=pc端内容链接"`
	MobileType    int         `json:"mobileType"    orm:"mobile_type"    description:"内容类型: 1=手机端编辑框内容 2=手机端内容链接"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:""`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:""`
}
