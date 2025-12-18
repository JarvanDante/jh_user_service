// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ActivityTag is the golang structure for table activity_tag.
type ActivityTag struct {
	Id         uint        `json:"id"         orm:"id"          description:""`
	SiteId     int         `json:"siteId"     orm:"site_id"     description:"站点ID"`
	ActivityId int         `json:"activityId" orm:"activity_id" description:"活动ID"`
	Name       string      `json:"name"       orm:"name"        description:"标签名称"`
	Status     int         `json:"status"     orm:"status"      description:"状态。1=启用；0=禁用；-1=删除"`
	Sort       int         `json:"sort"       orm:"sort"        description:"排序。值越小排名越靠前"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:""`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:""`
	TagId      int         `json:"tagId"      orm:"tag_id"      description:""`
}
