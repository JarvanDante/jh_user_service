// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AgentLog is the golang structure of table agent_log for DAO operations like Where/Data.
type AgentLog struct {
	g.Meta    `orm:"table:agent_log, do:true"`
	Id        any         //
	SiteId    any         // 站点ID
	AgentId   any         // 代理ID
	Username  any         // 代理用户名
	Ip        any         // ip
	CreatedAt *gtime.Time // 创建时间
}
