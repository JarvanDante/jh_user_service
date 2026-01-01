// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Site is the golang structure of table site for DAO operations like Where/Data.
type Site struct {
	g.Meta    `orm:"table:site, do:true"`
	Id        any         //
	Code      any         // 站点标识
	Name      any         // 站点名称
	Status    any         // 状态。1=正常；0=禁用
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
