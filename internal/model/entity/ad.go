// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Ad is the golang structure for table ad.
type Ad struct {
	Id          uint        `json:"id"          orm:"id"           description:""`
	SiteId      int         `json:"siteId"      orm:"site_id"      description:""`
	Type        int         `json:"type"        orm:"type"         description:""`
	Name        string      `json:"name"        orm:"name"         description:""`
	Image       string      `json:"image"       orm:"image"        description:""`
	Url         string      `json:"url"         orm:"url"          description:""`
	StartTime   *gtime.Time `json:"startTime"   orm:"start_time"   description:""`
	ExpiredTime *gtime.Time `json:"expiredTime" orm:"expired_time" description:""`
	Sort        int         `json:"sort"        orm:"sort"         description:""`
	Status      int         `json:"status"      orm:"status"       description:""`
	BeforeLogin int         `json:"beforeLogin" orm:"before_login" description:""`
	Position    int         `json:"position"    orm:"position"     description:"banner图位置"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:""`
}
