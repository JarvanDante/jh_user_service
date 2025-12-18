// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SiteRegister is the golang structure of table site_register for DAO operations like Where/Data.
type SiteRegister struct {
	g.Meta    `orm:"table:site_register, do:true"`
	Id        any         //
	SiteId    any         //
	Type      any         //
	Name      any         // 名称
	FieldName any         // 字段标识
	Display   any         // 是否显示。1=显示；0=不显示
	Required  any         // 是否必填。1=必填；0=选填
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
