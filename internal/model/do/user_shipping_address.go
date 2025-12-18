// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserShippingAddress is the golang structure of table user_shipping_address for DAO operations like Where/Data.
type UserShippingAddress struct {
	g.Meta    `orm:"table:user_shipping_address, do:true"`
	Id        any         //
	SiteId    any         // 站点ID
	UserId    any         // 用户ID
	Province  any         // 省份
	City      any         // 城市
	Town      any         // 区镇
	Address   any         // 详细地址
	Addressee any         // 收件人
	Mobile    any         // 手机号
	Phone     any         // 固定电话
	IsDefault any         // 是否是默认收货地址。1=是；0=否
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
