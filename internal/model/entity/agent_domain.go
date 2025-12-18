// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AgentDomain is the golang structure for table agent_domain.
type AgentDomain struct {
	Id        uint        `json:"id"        orm:"id"         description:""`
	SiteId    int         `json:"siteId"    orm:"site_id"    description:"站点ID"`
	AgentId   int         `json:"agentId"   orm:"agent_id"   description:"代理ID"`
	Domain    string      `json:"domain"    orm:"domain"     description:"域名"`
	Status    int         `json:"status"    orm:"status"     description:"状态。1=正常；0=禁用"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
