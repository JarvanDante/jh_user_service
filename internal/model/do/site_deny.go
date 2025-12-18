// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SiteDeny is the golang structure of table site_deny for DAO operations like Where/Data.
type SiteDeny struct {
	g.Meta    `orm:"table:site_deny, do:true"`
	Id        any         //
	SiteId    any         //
	Type      any         // 屏蔽类型。1=ip；2=地址
	Ip        any         // IP地址
	Address   any         // 地区名称
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
