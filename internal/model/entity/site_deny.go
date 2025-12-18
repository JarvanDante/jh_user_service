// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SiteDeny is the golang structure for table site_deny.
type SiteDeny struct {
	Id        uint        `json:"id"        orm:"id"         description:""`
	SiteId    int         `json:"siteId"    orm:"site_id"    description:""`
	Type      int         `json:"type"      orm:"type"       description:"屏蔽类型。1=ip；2=地址"`
	Ip        string      `json:"ip"        orm:"ip"         description:"IP地址"`
	Address   string      `json:"address"   orm:"address"    description:"地区名称"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
