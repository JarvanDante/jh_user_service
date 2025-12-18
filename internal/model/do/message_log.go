// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MessageLog is the golang structure of table message_log for DAO operations like Where/Data.
type MessageLog struct {
	g.Meta      `orm:"table:message_log, do:true"`
	Id          any         //
	SiteId      any         // 站点ID
	MessageId   any         // 消息ID
	MessageType any         // 发送类型。1=会员等级，2=会员层级，2=会员状态
	UserGrade   any         // 会员等级
	UserLevel   any         // 会员层级
	UserStatus  any         // 会员状态.0=离线；1=在线
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
