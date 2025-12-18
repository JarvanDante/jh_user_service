// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AgentMessage is the golang structure of table agent_message for DAO operations like Where/Data.
type AgentMessage struct {
	g.Meta    `orm:"table:agent_message, do:true"`
	Id        any         //
	SiteId    any         // 站点ID
	AdminId   any         // 员工ID
	Title     any         // 消息标题
	Content   any         // 消息内容
	Receiver  any         // 接收者
	Status    any         // 状态。1=正常；0=禁用
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
