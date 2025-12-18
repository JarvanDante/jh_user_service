// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SitePoints is the golang structure for table site_points.
type SitePoints struct {
	Id                   uint        `json:"id"                   orm:"id"                     description:""`
	SiteId               int         `json:"siteId"               orm:"site_id"                description:""`
	SwitchSitePoints     int         `json:"switchSitePoints"     orm:"switch_site_points"     description:""`
	SwitchRechargePoints int         `json:"switchRechargePoints" orm:"switch_recharge_points" description:""`
	EachRechargeAmount   float64     `json:"eachRechargeAmount"   orm:"each_recharge_amount"   description:""`
	EachRechargePoints   int         `json:"eachRechargePoints"   orm:"each_recharge_points"   description:""`
	SwitchBettingPoints  int         `json:"switchBettingPoints"  orm:"switch_betting_points"  description:""`
	EachBettingAmount    float64     `json:"eachBettingAmount"    orm:"each_betting_amount"    description:""`
	EachBettingPoints    int         `json:"eachBettingPoints"    orm:"each_betting_points"    description:""`
	MaxDailyPoints       int         `json:"maxDailyPoints"       orm:"max_daily_points"       description:""`
	CreatedAt            *gtime.Time `json:"createdAt"            orm:"created_at"             description:""`
	UpdatedAt            *gtime.Time `json:"updatedAt"            orm:"updated_at"             description:""`
}
