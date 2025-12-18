// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AgentDomain is the golang structure of table agent_domain for DAO operations like Where/Data.
type AgentDomain struct {
	g.Meta    `orm:"table:agent_domain, do:true"`
	Id        any         //
	SiteId    any         // 站点ID
	AgentId   any         // 代理ID
	Domain    any         // 域名
	Status    any         // 状态。1=正常；0=禁用
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
