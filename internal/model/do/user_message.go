// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserMessage is the golang structure of table user_message for DAO operations like Where/Data.
type UserMessage struct {
	g.Meta    `orm:"table:user_message, do:true"`
	Id        any         //
	SiteId    any         // 站点ID
	UserId    any         // 用户ID
	MessageId any         // 消息ID
	Status    any         // 状态。1=已读；0=未读
	DeleteAt  any         // 是否删除。0=未删除；其他为删除时间戳
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
