// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RiskForewarnLog is the golang structure for table risk_forewarn_log.
type RiskForewarnLog struct {
	Id        uint64      `json:"id"        orm:"id"         description:""`
	SiteId    int         `json:"siteId"    orm:"site_id"    description:"站点ID"`
	UserId    int         `json:"userId"    orm:"user_id"    description:"会员ID"`
	Username  string      `json:"username"  orm:"username"   description:"用户名"`
	Content   string      `json:"content"   orm:"content"    description:"预警提示内容"`
	Ip        string      `json:"ip"        orm:"ip"         description:"用户IP"`
	Device    int         `json:"device"    orm:"device"     description:"终端。1=电脑；2=手机；3=平板"`
	Status    int         `json:"status"    orm:"status"     description:""`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
