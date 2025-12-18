// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityDao is the data access object for the table activity.
type ActivityDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ActivityColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ActivityColumns defines and stores column names for the table activity.
type ActivityColumns struct {
	Id               string // 活动ID
	SiteId           string // 站点ID
	ActivityName     string // 活动名称
	Describe         string // 活动描述
	PcCover          string // pc端封面
	WapCover         string // wap端封面
	StartTime        string // 开始时间
	EndTime          string // 结束时间
	Status           string // 活动状态。1=开启；0=关闭
	IsRelateGame     string // 是否关联游戏
	RelateGame       string // 关联游戏
	TemplateId       string // 活动模板ID
	ActivityDetails  string // 活动详情
	ActivityModuleId string // 活动模块ID
	IsPublish        string // 发布状态。1=发布；0=未发布
	Sort             string // 排序
	CreatedAt        string //
	UpdatedAt        string //
	ActivityType     string // 活动类型
	Name             string // 名称
	MobileCover      string // 手机端封面
	Intro            string //
}

// activityColumns holds the columns for the table activity.
var activityColumns = ActivityColumns{
	Id:               "id",
	SiteId:           "site_id",
	ActivityName:     "activity_name",
	Describe:         "describe",
	PcCover:          "pc_cover",
	WapCover:         "wap_cover",
	StartTime:        "start_time",
	EndTime:          "end_time",
	Status:           "status",
	IsRelateGame:     "is_relate_game",
	RelateGame:       "relate_game",
	TemplateId:       "template_id",
	ActivityDetails:  "activity_details",
	ActivityModuleId: "activity_module_id",
	IsPublish:        "is_publish",
	Sort:             "sort",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	ActivityType:     "activity_type",
	Name:             "name",
	MobileCover:      "mobile_cover",
	Intro:            "intro",
}

// NewActivityDao creates and returns a new DAO object for table data access.
func NewActivityDao(handlers ...gdb.ModelHandler) *ActivityDao {
	return &ActivityDao{
		group:    "default",
		table:    "activity",
		columns:  activityColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ActivityDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ActivityDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ActivityDao) Columns() ActivityColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ActivityDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ActivityDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ActivityDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
