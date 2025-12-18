// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminCustomField is the golang structure for table admin_custom_field.
type AdminCustomField struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	SiteId    int         `json:"siteId"    orm:"site_id"    description:"站点ID"`
	AdminId   int         `json:"adminId"   orm:"admin_id"   description:"员工ID"`
	Page      int         `json:"page"      orm:"page"       description:"页面。1=会员列表；默认=1"`
	Fields    string      `json:"fields"    orm:"fields"     description:"字段列表。以,隔开"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
