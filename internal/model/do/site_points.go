// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SitePoints is the golang structure of table site_points for DAO operations like Where/Data.
type SitePoints struct {
	g.Meta               `orm:"table:site_points, do:true"`
	Id                   any         //
	SiteId               any         //
	SwitchSitePoints     any         //
	SwitchRechargePoints any         //
	EachRechargeAmount   any         //
	EachRechargePoints   any         //
	SwitchBettingPoints  any         //
	EachBettingAmount    any         //
	EachBettingPoints    any         //
	MaxDailyPoints       any         //
	CreatedAt            *gtime.Time //
	UpdatedAt            *gtime.Time //
}
