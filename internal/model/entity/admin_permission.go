// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminPermission is the golang structure for table admin_permission.
type AdminPermission struct {
	Id          uint        `json:"id"          orm:"id"          description:"权限ID"`
	ParentId    int         `json:"parentId"    orm:"parent_id"   description:"父权限ID"`
	Type        int         `json:"type"        orm:"type"        description:"权限类型。1=菜单；2=操作权限"`
	Name        string      `json:"name"        orm:"name"        description:"权限名称"`
	BackendUrl  string      `json:"backendUrl"  orm:"backend_url" description:"后台URL"`
	FrontendUrl string      `json:"frontendUrl" orm:"frontend_url" description:"前台URL"`
	Status      int         `json:"status"      orm:"status"      description:"状态。1=启用；0=禁用"`
	Sort        int         `json:"sort"        orm:"sort"        description:"排序"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"更新时间"`
}
