// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserBank is the golang structure for table user_bank.
type UserBank struct {
	Id          uint        `json:"id"          orm:"id"           description:""`
	SiteId      int         `json:"siteId"      orm:"site_id"      description:"站点ID"`
	UserId      int         `json:"userId"      orm:"user_id"      description:"会员ID"`
	BankName    string      `json:"bankName"    orm:"bank_name"    description:"银行名称"`
	CardAccount string      `json:"cardAccount" orm:"card_account" description:"银行账户名.这个应该与会员实名一致"`
	CardNo      string      `json:"cardNo"      orm:"card_no"      description:"银行卡号"`
	DepositBank string      `json:"depositBank" orm:"deposit_bank" description:"开户行"`
	IsDefault   int         `json:"isDefault"   orm:"is_default"   description:"默认地址。1=默认，0=非默认"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:""`
}
