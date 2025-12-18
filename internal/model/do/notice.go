// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Notice is the golang structure of table notice for DAO operations like Where/Data.
type Notice struct {
	g.Meta      `orm:"table:notice, do:true"`
	Id          any         //
	SiteId      any         //
	Type        any         //
	Title       any         //
	Url         any         //
	StartTime   *gtime.Time //
	ExpiredTime *gtime.Time //
	Sort        any         //
	Status      any         //
	Platform    any         //
	Content     any         //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
