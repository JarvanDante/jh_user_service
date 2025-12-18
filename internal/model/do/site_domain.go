// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SiteDomain is the golang structure of table site_domain for DAO operations like Where/Data.
type SiteDomain struct {
	g.Meta    `orm:"table:site_domain, do:true"`
	Id        any         //
	SiteId    any         //
	Domain    any         //
	Status    any         //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
