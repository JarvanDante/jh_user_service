// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CustomTag is the golang structure for table custom_tag.
type CustomTag struct {
	Id         uint        `json:"id"         orm:"id"          description:""`
	SiteId     int         `json:"siteId"     orm:"site_id"     description:""`
	Type       int         `json:"type"       orm:"type"        description:"类型 1=充值 2=提现"`
	Status     int         `json:"status"     orm:"status"      description:"1=正常，0=禁用"`
	TagName    string      `json:"tagName"    orm:"tag_name"    description:"标签名称"`
	TagContent string      `json:"tagContent" orm:"tag_content" description:"标签内容"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:""`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:""`
}
