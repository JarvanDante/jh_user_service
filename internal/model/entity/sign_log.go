// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SignLog is the golang structure for table sign_log.
type SignLog struct {
	Id        uint        `json:"id"        orm:"id"         description:""`
	SiteId    int         `json:"siteId"    orm:"site_id"    description:"站点ID"`
	UserId    int         `json:"userId"    orm:"user_id"    description:"会员ID"`
	SignDate  *gtime.Time `json:"signDate"  orm:"sign_date"  description:"签到日期"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
