// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AgentMessageLog is the golang structure of table agent_message_log for DAO operations like Where/Data.
type AgentMessageLog struct {
	g.Meta         `orm:"table:agent_message_log, do:true"`
	Id             any         //
	SiteId         any         // 站点ID
	AgentMessageId any         // 代理消息id
	AgentId        any         // 代理id
	IsRead         any         // 是否读过
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
}
