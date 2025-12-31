// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminPermission is the golang structure of table admin_permission for DAO operations like Where/Data.
type AdminPermission struct {
	g.Meta      `orm:"table:admin_permission, do:true"`
	Id          interface{} // 权限ID
	ParentId    interface{} // 父权限ID
	Type        interface{} // 权限类型。1=菜单；2=操作权限
	Name        interface{} // 权限名称
	BackendUrl  interface{} // 后台URL
	FrontendUrl interface{} // 前台URL
	Status      interface{} // 状态。1=启用；0=禁用
	Sort        interface{} // 排序
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
