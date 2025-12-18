// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminRole is the golang structure of table admin_role for DAO operations like Where/Data.
type AdminRole struct {
	g.Meta      `orm:"table:admin_role, do:true"`
	Id          any         //
	SiteId      any         //
	Name        any         // 角色名称
	Status      any         // 状态。1=可用；0=禁用
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	Permissions any         // 权限列表。格式：权限id以,隔开
}
