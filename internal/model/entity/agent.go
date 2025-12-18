// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Agent is the golang structure for table agent.
type Agent struct {
	Id                  uint        `json:"id"                  orm:"id"                    description:""`
	SiteId              int         `json:"siteId"              orm:"site_id"               description:"站点ID"`
	Username            string      `json:"username"            orm:"username"              description:"用户名"`
	Password            string      `json:"password"            orm:"password"              description:"密码"`
	Status              int         `json:"status"              orm:"status"                description:"审核状态。0=停用；1=启用"`
	StatusCheck         int         `json:"statusCheck"         orm:"status_check"          description:"审核状态。0=待审核；1=通过；2=未通过"`
	Level               int         `json:"level"               orm:"level"                 description:"代理层级"`
	Realname            string      `json:"realname"            orm:"realname"              description:"真实姓名"`
	Mobile              string      `json:"mobile"              orm:"mobile"                description:"手机号"`
	Email               string      `json:"email"               orm:"email"                 description:"电子邮箱"`
	Qq                  string      `json:"qq"                  orm:"qq"                    description:"QQ"`
	Skype               string      `json:"skype"               orm:"skype"                 description:"Skype"`
	BankName            string      `json:"bankName"            orm:"bank_name"             description:"银行名称"`
	CardAccount         string      `json:"cardAccount"         orm:"card_account"          description:"银行户名"`
	CardNo              string      `json:"cardNo"              orm:"card_no"               description:"卡号"`
	DepositBankProvince string      `json:"depositBankProvince" orm:"deposit_bank_province" description:"开户行所在省"`
	DepositBankCity     string      `json:"depositBankCity"     orm:"deposit_bank_city"     description:"开户行所在市"`
	DepositBank         string      `json:"depositBank"         orm:"deposit_bank"          description:"开户行"`
	RegisterIp          string      `json:"registerIp"          orm:"register_ip"           description:""`
	RegisterTime        *gtime.Time `json:"registerTime"        orm:"register_time"         description:""`
	LastLoginIp         string      `json:"lastLoginIp"         orm:"last_login_ip"         description:""`
	LastLoginTime       *gtime.Time `json:"lastLoginTime"       orm:"last_login_time"       description:""`
	CreatedAt           *gtime.Time `json:"createdAt"           orm:"created_at"            description:""`
	UpdatedAt           *gtime.Time `json:"updatedAt"           orm:"updated_at"            description:""`
	Reason              string      `json:"reason"              orm:"reason"                description:"申请理由"`
	Agentcode           string      `json:"agentcode"           orm:"agentcode"             description:""`
}
