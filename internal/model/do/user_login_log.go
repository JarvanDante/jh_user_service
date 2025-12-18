// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLoginLog is the golang structure of table user_login_log for DAO operations like Where/Data.
type UserLoginLog struct {
	g.Meta       `orm:"table:user_login_log, do:true"`
	Id           any         //
	SiteId       any         //
	UserId       any         // 会员ID
	Username     any         // 会员用户名
	RefererUrl   any         // 来源网址
	LoginUrl     any         // 登录网址
	LoginTime    *gtime.Time // 登录时间
	LoginIp      any         // 登录IP
	LoginAddress any         // 登录地址
	Os           any         // 操作系统
	Network      any         // 网络
	Screen       any         // 分辨率
	Browser      any         // 浏览器
	Device       any         // 终端。1=电脑；2=手机；3=平板
	IsRobot      any         // 判断是否是机器人登录。1=是；0=否
	CreatedAt    *gtime.Time // 创建时间
}
