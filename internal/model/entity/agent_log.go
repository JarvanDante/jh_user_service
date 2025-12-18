// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AgentLog is the golang structure for table agent_log.
type AgentLog struct {
	Id        uint        `json:"id"        orm:"id"         description:""`
	SiteId    int         `json:"siteId"    orm:"site_id"    description:"站点ID"`
	AgentId   int         `json:"agentId"   orm:"agent_id"   description:"代理ID"`
	Username  string      `json:"username"  orm:"username"   description:"代理用户名"`
	Ip        string      `json:"ip"        orm:"ip"         description:"ip"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
}
