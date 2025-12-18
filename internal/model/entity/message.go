// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Message is the golang structure for table message.
type Message struct {
	Id        uint        `json:"id"        orm:"id"         description:""`
	SiteId    int         `json:"siteId"    orm:"site_id"    description:"站点ID"`
	AdminId   int         `json:"adminId"   orm:"admin_id"   description:"员工ID"`
	Title     string      `json:"title"     orm:"title"      description:"消息标题"`
	Content   string      `json:"content"   orm:"content"    description:"消息内容"`
	Receiver  string      `json:"receiver"  orm:"receiver"   description:"接收者"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
