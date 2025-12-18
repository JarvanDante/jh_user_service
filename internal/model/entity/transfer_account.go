// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TransferAccount is the golang structure for table transfer_account.
type TransferAccount struct {
	Id          uint        `json:"id"          orm:"id"           description:""`
	SiteId      int         `json:"siteId"      orm:"site_id"      description:"站点ID"`
	BankType    int         `json:"bankType"    orm:"bank_type"    description:"银行类型"`
	BankName    string      `json:"bankName"    orm:"bank_name"    description:"银行名称"`
	Name        string      `json:"name"        orm:"name"         description:"转账接口名称"`
	BankUrl     string      `json:"bankUrl"     orm:"bank_url"     description:"银行链接"`
	Qrcode      string      `json:"qrcode"      orm:"qrcode"       description:"二维码图片地址"`
	CardAccount string      `json:"cardAccount" orm:"card_account" description:"银行户名或者第三方收款人"`
	CardNo      string      `json:"cardNo"      orm:"card_no"      description:"银行卡号或者第三方账号"`
	DepositBank string      `json:"depositBank" orm:"deposit_bank" description:"开户行"`
	EachMin     float64     `json:"eachMin"     orm:"each_min"     description:"单笔最低。默认为1元"`
	EachMax     float64     `json:"eachMax"     orm:"each_max"     description:"单笔最高。默认为5000元"`
	DailyMax    float64     `json:"dailyMax"    orm:"daily_max"    description:"单日上限。默认为0。0=不限"`
	TodayCount  int         `json:"todayCount"  orm:"today_count"  description:"今日入款次数"`
	TodayAmount float64     `json:"todayAmount" orm:"today_amount" description:"今日转账总额"`
	Status      int         `json:"status"      orm:"status"       description:"状态。1=可用；0=禁用"`
	Sort        int         `json:"sort"        orm:"sort"         description:"排序。值越小排名越靠前"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:""`
	Remark      string      `json:"remark"      orm:"remark"       description:""`
}
