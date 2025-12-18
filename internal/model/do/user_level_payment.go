// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLevelPayment is the golang structure of table user_level_payment for DAO operations like Where/Data.
type UserLevelPayment struct {
	g.Meta           `orm:"table:user_level_payment, do:true"`
	Id               any         //
	SiteId           any         // 站点ID
	UserLevelId      any         // 会员层级ID
	Gateway          any         // 支付类型
	PaymentAccountId any         // 支付接口ID
	CreatedAt        *gtime.Time //
	UpdatedAt        *gtime.Time //
}
