// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminRole is the golang structure for table admin_role.
type AdminRole struct {
	Id          uint        `json:"id"          orm:"id"          description:""`
	SiteId      int         `json:"siteId"      orm:"site_id"     description:""`
	Name        string      `json:"name"        orm:"name"        description:"角色名称"`
	Status      int         `json:"status"      orm:"status"      description:"状态。1=可用；0=禁用"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:""`
	Permissions string      `json:"permissions" orm:"permissions" description:"权限列表。格式：权限id以,隔开"`
}
