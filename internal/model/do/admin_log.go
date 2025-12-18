// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminLog is the golang structure of table admin_log for DAO operations like Where/Data.
type AdminLog struct {
	g.Meta        `orm:"table:admin_log, do:true"`
	Id            any         //
	SiteId        any         //
	AdminId       any         //
	AdminUsername any         //
	Ip            any         //
	CreatedAt     *gtime.Time //
	Remark        any         //
}
