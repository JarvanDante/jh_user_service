// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserMessage is the golang structure for table user_message.
type UserMessage struct {
	Id        uint        `json:"id"        orm:"id"         description:""`
	SiteId    int         `json:"siteId"    orm:"site_id"    description:"站点ID"`
	UserId    int         `json:"userId"    orm:"user_id"    description:"用户ID"`
	MessageId int         `json:"messageId" orm:"message_id" description:"消息ID"`
	Status    int         `json:"status"    orm:"status"     description:"状态。1=已读；0=未读"`
	DeleteAt  uint        `json:"deleteAt"  orm:"delete_at"  description:"是否删除。0=未删除；其他为删除时间戳"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
