// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SiteSign is the golang structure for table site_sign.
type SiteSign struct {
	Id        uint        `json:"id"        orm:"id"         description:""`
	SiteId    int         `json:"siteId"    orm:"site_id"    description:"站点ID"`
	Name      string      `json:"name"      orm:"name"       description:"活动名称"`
	StartTime *gtime.Time `json:"startTime" orm:"start_time" description:"开始时间"`
	EndTime   *gtime.Time `json:"endTime"   orm:"end_time"   description:"结束时间"`
	Status    int         `json:"status"    orm:"status"     description:"活动状态。1=开启；0=关闭"`
	UserGrade string      `json:"userGrade" orm:"user_grade" description:"会员等级。以,的形式隔开"`
	UserLevel string      `json:"userLevel" orm:"user_level" description:"会员层级。以,的形式隔开"`
	Platform  int         `json:"platform"  orm:"platform"   description:"优惠终端。0=所有；1=网站；2=手机"`
	Remark    string      `json:"remark"    orm:"remark"     description:"活动描述"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
