// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Ad is the golang structure of table ad for DAO operations like Where/Data.
type Ad struct {
	g.Meta      `orm:"table:ad, do:true"`
	Id          any         //
	SiteId      any         //
	Type        any         //
	Name        any         //
	Image       any         //
	Url         any         //
	StartTime   *gtime.Time //
	ExpiredTime *gtime.Time //
	Sort        any         //
	Status      any         //
	BeforeLogin any         //
	Position    any         // banner图位置
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
