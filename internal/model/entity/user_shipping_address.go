// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserShippingAddress is the golang structure for table user_shipping_address.
type UserShippingAddress struct {
	Id        uint        `json:"id"        orm:"id"         description:""`
	SiteId    int         `json:"siteId"    orm:"site_id"    description:"站点ID"`
	UserId    int         `json:"userId"    orm:"user_id"    description:"用户ID"`
	Province  string      `json:"province"  orm:"province"   description:"省份"`
	City      string      `json:"city"      orm:"city"       description:"城市"`
	Town      string      `json:"town"      orm:"town"       description:"区镇"`
	Address   string      `json:"address"   orm:"address"    description:"详细地址"`
	Addressee string      `json:"addressee" orm:"addressee"  description:"收件人"`
	Mobile    string      `json:"mobile"    orm:"mobile"     description:"手机号"`
	Phone     string      `json:"phone"     orm:"phone"      description:"固定电话"`
	IsDefault int         `json:"isDefault" orm:"is_default" description:"是否是默认收货地址。1=是；0=否"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
