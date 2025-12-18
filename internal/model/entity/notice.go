// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Notice is the golang structure for table notice.
type Notice struct {
	Id          uint        `json:"id"          orm:"id"           description:""`
	SiteId      int         `json:"siteId"      orm:"site_id"      description:""`
	Type        int         `json:"type"        orm:"type"         description:""`
	Title       string      `json:"title"       orm:"title"        description:""`
	Url         string      `json:"url"         orm:"url"          description:""`
	StartTime   *gtime.Time `json:"startTime"   orm:"start_time"   description:""`
	ExpiredTime *gtime.Time `json:"expiredTime" orm:"expired_time" description:""`
	Sort        int         `json:"sort"        orm:"sort"         description:""`
	Status      int         `json:"status"      orm:"status"       description:""`
	Platform    int         `json:"platform"    orm:"platform"     description:""`
	Content     string      `json:"content"     orm:"content"      description:""`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:""`
}
