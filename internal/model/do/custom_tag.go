// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CustomTag is the golang structure of table custom_tag for DAO operations like Where/Data.
type CustomTag struct {
	g.Meta     `orm:"table:custom_tag, do:true"`
	Id         any         //
	SiteId     any         //
	Type       any         // 类型 1=充值 2=提现
	Status     any         // 1=正常，0=禁用
	TagName    any         // 标签名称
	TagContent any         // 标签内容
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
}
