// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminCustomField is the golang structure of table admin_custom_field for DAO operations like Where/Data.
type AdminCustomField struct {
	g.Meta    `orm:"table:admin_custom_field, do:true"`
	Id        any         //
	SiteId    any         // 站点ID
	AdminId   any         // 员工ID
	Page      any         // 页面。1=会员列表；默认=1
	Fields    any         // 字段列表。以,隔开
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
