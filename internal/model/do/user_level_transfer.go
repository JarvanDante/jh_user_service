// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLevelTransfer is the golang structure of table user_level_transfer for DAO operations like Where/Data.
type UserLevelTransfer struct {
	g.Meta            `orm:"table:user_level_transfer, do:true"`
	Id                any         //
	SiteId            any         // 站点ID
	UserLevelId       any         // 会员层级ID
	BankType          any         // 转账类型。1=网银转账；2=微信；3=支付宝
	TransferAccountId any         // 转账接口ID
	CreatedAt         *gtime.Time //
	UpdatedAt         *gtime.Time //
}
