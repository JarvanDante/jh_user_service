// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AgentMessageLog is the golang structure for table agent_message_log.
type AgentMessageLog struct {
	Id             uint        `json:"id"             orm:"id"               description:""`
	SiteId         int         `json:"siteId"         orm:"site_id"          description:"站点ID"`
	AgentMessageId int         `json:"agentMessageId" orm:"agent_message_id" description:"代理消息id"`
	AgentId        int         `json:"agentId"        orm:"agent_id"         description:"代理id"`
	IsRead         int         `json:"isRead"         orm:"is_read"          description:"是否读过"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:""`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       description:""`
}
