// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TransferAccountDao is the data access object for the table transfer_account.
type TransferAccountDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  TransferAccountColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// TransferAccountColumns defines and stores column names for the table transfer_account.
type TransferAccountColumns struct {
	Id          string //
	SiteId      string // 站点ID
	BankType    string // 银行类型
	BankName    string // 银行名称
	Name        string // 转账接口名称
	BankUrl     string // 银行链接
	Qrcode      string // 二维码图片地址
	CardAccount string // 银行户名或者第三方收款人
	CardNo      string // 银行卡号或者第三方账号
	DepositBank string // 开户行
	EachMin     string // 单笔最低。默认为1元
	EachMax     string // 单笔最高。默认为5000元
	DailyMax    string // 单日上限。默认为0。0=不限
	TodayCount  string // 今日入款次数
	TodayAmount string // 今日转账总额
	Status      string // 状态。1=可用；0=禁用
	Sort        string // 排序。值越小排名越靠前
	CreatedAt   string //
	UpdatedAt   string //
	Remark      string //
}

// transferAccountColumns holds the columns for the table transfer_account.
var transferAccountColumns = TransferAccountColumns{
	Id:          "id",
	SiteId:      "site_id",
	BankType:    "bank_type",
	BankName:    "bank_name",
	Name:        "name",
	BankUrl:     "bank_url",
	Qrcode:      "qrcode",
	CardAccount: "card_account",
	CardNo:      "card_no",
	DepositBank: "deposit_bank",
	EachMin:     "each_min",
	EachMax:     "each_max",
	DailyMax:    "daily_max",
	TodayCount:  "today_count",
	TodayAmount: "today_amount",
	Status:      "status",
	Sort:        "sort",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	Remark:      "remark",
}

// NewTransferAccountDao creates and returns a new DAO object for table data access.
func NewTransferAccountDao(handlers ...gdb.ModelHandler) *TransferAccountDao {
	return &TransferAccountDao{
		group:    "default",
		table:    "transfer_account",
		columns:  transferAccountColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TransferAccountDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TransferAccountDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TransferAccountDao) Columns() TransferAccountColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TransferAccountDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TransferAccountDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TransferAccountDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
