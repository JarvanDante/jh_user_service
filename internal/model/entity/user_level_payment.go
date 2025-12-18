// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLevelPayment is the golang structure for table user_level_payment.
type UserLevelPayment struct {
	Id               uint        `json:"id"               orm:"id"                 description:""`
	SiteId           int         `json:"siteId"           orm:"site_id"            description:"站点ID"`
	UserLevelId      int         `json:"userLevelId"      orm:"user_level_id"      description:"会员层级ID"`
	Gateway          int         `json:"gateway"          orm:"gateway"            description:"支付类型"`
	PaymentAccountId int         `json:"paymentAccountId" orm:"payment_account_id" description:"支付接口ID"`
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"         description:""`
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"         description:""`
}
