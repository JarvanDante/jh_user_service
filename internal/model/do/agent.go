// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Agent is the golang structure of table agent for DAO operations like Where/Data.
type Agent struct {
	g.Meta              `orm:"table:agent, do:true"`
	Id                  any         //
	SiteId              any         // 站点ID
	Username            any         // 用户名
	Password            any         // 密码
	Status              any         // 审核状态。0=停用；1=启用
	StatusCheck         any         // 审核状态。0=待审核；1=通过；2=未通过
	Level               any         // 代理层级
	Realname            any         // 真实姓名
	Mobile              any         // 手机号
	Email               any         // 电子邮箱
	Qq                  any         // QQ
	Skype               any         // Skype
	BankName            any         // 银行名称
	CardAccount         any         // 银行户名
	CardNo              any         // 卡号
	DepositBankProvince any         // 开户行所在省
	DepositBankCity     any         // 开户行所在市
	DepositBank         any         // 开户行
	RegisterIp          any         //
	RegisterTime        *gtime.Time //
	LastLoginIp         any         //
	LastLoginTime       *gtime.Time //
	CreatedAt           *gtime.Time //
	UpdatedAt           *gtime.Time //
	Reason              any         // 申请理由
	Agentcode           any         //
}
