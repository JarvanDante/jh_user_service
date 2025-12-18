// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MessageLog is the golang structure for table message_log.
type MessageLog struct {
	Id          uint        `json:"id"          orm:"id"           description:""`
	SiteId      int         `json:"siteId"      orm:"site_id"      description:"站点ID"`
	MessageId   int         `json:"messageId"   orm:"message_id"   description:"消息ID"`
	MessageType int         `json:"messageType" orm:"message_type" description:"发送类型。1=会员等级，2=会员层级，2=会员状态"`
	UserGrade   int         `json:"userGrade"   orm:"user_grade"   description:"会员等级"`
	UserLevel   int         `json:"userLevel"   orm:"user_level"   description:"会员层级"`
	UserStatus  int         `json:"userStatus"  orm:"user_status"  description:"会员状态.0=离线；1=在线"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:""`
}
