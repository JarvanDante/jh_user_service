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
	Id          interface{} // 角色ID
	SiteId      interface{} // 站点ID
	Name        interface{} // 角色名称
	Status      interface{} // 状态。1=启用；0=禁用
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	Permissions interface{} // 权限列表。格式：权限id以,隔开
}
