// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SignLog is the golang structure of table sign_log for DAO operations like Where/Data.
type SignLog struct {
	g.Meta    `orm:"table:sign_log, do:true"`
	Id        any         //
	SiteId    any         // 站点ID
	UserId    any         // 会员ID
	SignDate  *gtime.Time // 签到日期
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
