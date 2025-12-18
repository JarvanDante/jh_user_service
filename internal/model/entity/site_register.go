// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SiteRegister is the golang structure for table site_register.
type SiteRegister struct {
	Id        uint        `json:"id"        orm:"id"         description:""`
	SiteId    int         `json:"siteId"    orm:"site_id"    description:""`
	Type      int         `json:"type"      orm:"type"       description:""`
	Name      string      `json:"name"      orm:"name"       description:"名称"`
	FieldName string      `json:"fieldName" orm:"field_name" description:"字段标识"`
	Display   int         `json:"display"   orm:"display"    description:"是否显示。1=显示；0=不显示"`
	Required  int         `json:"required"  orm:"required"   description:"是否必填。1=必填；0=选填"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
