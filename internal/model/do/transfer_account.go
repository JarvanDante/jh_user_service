// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TransferAccount is the golang structure of table transfer_account for DAO operations like Where/Data.
type TransferAccount struct {
	g.Meta      `orm:"table:transfer_account, do:true"`
	Id          any         //
	SiteId      any         // 站点ID
	BankType    any         // 银行类型
	BankName    any         // 银行名称
	Name        any         // 转账接口名称
	BankUrl     any         // 银行链接
	Qrcode      any         // 二维码图片地址
	CardAccount any         // 银行户名或者第三方收款人
	CardNo      any         // 银行卡号或者第三方账号
	DepositBank any         // 开户行
	EachMin     any         // 单笔最低。默认为1元
	EachMax     any         // 单笔最高。默认为5000元
	DailyMax    any         // 单日上限。默认为0。0=不限
	TodayCount  any         // 今日入款次数
	TodayAmount any         // 今日转账总额
	Status      any         // 状态。1=可用；0=禁用
	Sort        any         // 排序。值越小排名越靠前
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	Remark      any         //
}
