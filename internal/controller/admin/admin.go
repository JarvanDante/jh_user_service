package admin

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	v1 "jh_user_service/api/admin/v1"
	"jh_user_service/internal/dao"
	"jh_user_service/internal/model/do"
	"jh_user_service/internal/model/entity"
)

type Controller struct {
	v1.UnimplementedAdminServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterAdminServer(s.Server, &Controller{})
}

// RegisterHTTP 注册 HTTP 路由
func RegisterHTTP(s *ghttp.Server) {
	s.Group("/api/admin", func(group *ghttp.RouterGroup) {
		group.POST("/login", (*Controller).LoginHTTP)
		group.GET("/refresh-token", (*Controller).RefreshTokenHTTP)
		group.POST("/create-admin", (*Controller).CreateAdminHTTP)
	})
}

// Login 管理员登录
func (*Controller) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	// 获取站点ID (这里需要根据实际情况获取，可能从上下文或配置中获取)
	siteId := 1 // 临时硬编码，实际应该从请求中获取

	g.Log().Infof(ctx, "登录请求 - 用户名: %s, 站点ID: %d", req.Username, siteId)

	// 测试数据库连接和查询
	g.Log().Infof(ctx, "开始数据库查询调试...")

	// 查看admin表的记录数
	countResult, countErr := dao.Admin.DB().GetValue(ctx, "SELECT COUNT(*) FROM admin")
	if countErr != nil {
		g.Log().Errorf(ctx, "查询admin表记录数失败: %v", countErr)
	} else {
		g.Log().Infof(ctx, "admin表总记录数: %s", countResult.String())
	}

	// 使用DAO查询所有管理员
	var admins []*entity.Admin
	scanErr := dao.Admin.Ctx(ctx).Scan(&admins)
	if scanErr != nil {
		g.Log().Errorf(ctx, "DAO查询失败: %v", scanErr)
		return nil, fmt.Errorf("数据库查询错误: %v", scanErr)
	}
	g.Log().Infof(ctx, "DAO查询到 %d 条管理员记录", len(admins))

	// 打印前几条记录用于调试
	for i, admin := range admins {
		if i >= 3 { // 只打印前3条
			break
		}
		g.Log().Infof(ctx, "管理员 %d: ID=%d, Username=%s, Status=%d",
			i+1, admin.Id, admin.Username, admin.Status)
	}
	// 查询管理员 - 使用原生SQL避免自动添加条件
	g.Log().Infof(ctx, "开始查询管理员 - 用户名: %s, 站点ID: %d", req.Username, siteId)

	var admin *entity.Admin
	err = dao.Admin.Ctx(ctx).Where("username = ? AND site_id = ?", req.Username, siteId).Scan(&admin)

	if err != nil {
		g.Log().Errorf(ctx, "数据库查询错误: %v", err)
		return nil, fmt.Errorf("数据库查询错误: %v", err)
	}

	if admin == nil {
		g.Log().Warningf(ctx, "未找到管理员记录 - 用户名: %s, 站点ID: %d", req.Username, siteId)
		return nil, fmt.Errorf("用户名或密码错误")
	}

	g.Log().Infof(ctx, "找到管理员记录 - ID: %d, 用户名: %s, 状态: %d", admin.Id, admin.Username, admin.Status)

	// 检查状态（状态检查放在找到记录之后）
	if admin.Status != 1 {
		g.Log().Warningf(ctx, "管理员状态异常 - 用户名: %s, 状态: %d", req.Username, admin.Status)
		return nil, fmt.Errorf("账号已被禁用")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password))
	if err != nil {
		g.Log().Warningf(ctx, "密码验证失败 - 用户名: %s, 错误: %v", req.Username, err)
		return nil, fmt.Errorf("用户名或密码错误")
	}

	g.Log().Infof(ctx, "密码验证成功 - 用户名: %s", req.Username)

	// 验证Google 2FA (如果开启)
	if admin.SwitchGoogle2Fa == 1 {
		if req.Code == "" {
			return nil, fmt.Errorf("请输入动态验证码")
		}
		// 这里需要实现Google 2FA验证逻辑
		// valid := validateGoogle2FA(admin.Google2FaSecret, req.Code)
		// if !valid {
		//     return nil, fmt.Errorf("动态验证码错误")
		// }
	}

	// 生成JWT token
	token, err := generateJWTToken(admin)
	if err != nil {
		return nil, fmt.Errorf("生成token失败: %v", err)
	}

	// 更新最后登录信息
	_, err = dao.Admin.Ctx(ctx).Where(do.Admin{Id: admin.Id}).Update(do.Admin{
		LastLoginIp:   getClientIP(ctx),
		LastLoginTime: gtime.Now(),
	})
	if err != nil {
		g.Log().Error(ctx, "更新登录信息失败:", err)
	}

	// 记录登录日志
	err = addAdminLog(ctx, admin, "登录成功")
	if err != nil {
		g.Log().Error(ctx, "记录登录日志失败:", err)
	}

	// 获取socket地址 (从配置中获取)
	socketAddr := g.Cfg().MustGet(ctx, "workerman.host", "").String()
	if port := g.Cfg().MustGet(ctx, "workerman.port", "").String(); port != "" {
		socketAddr = socketAddr + ":" + port
	}

	res = &v1.LoginRes{
		Token:  token,
		Socket: socketAddr,
	}

	g.Log().Infof(ctx, "登录成功 - 用户名: %s", req.Username)
	return res, nil
}

// RefreshToken 刷新token
func (*Controller) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error) {
	// 从上下文中获取当前用户信息 (需要中间件解析JWT)
	// 这里简化处理，实际需要从JWT中解析用户信息

	// 重新生成token
	// admin := getCurrentAdmin(ctx)
	// token, err := generateJWTToken(admin)
	// if err != nil {
	//     return nil, fmt.Errorf("刷新token失败: %v", err)
	// }

	res = &v1.RefreshTokenRes{
		Token: "new_token", // 临时返回
	}

	return res, nil
}

// CreateAdmin 创建管理员
func (*Controller) CreateAdmin(ctx context.Context, req *v1.CreateAdminReq) (res *v1.CreateAdminRes, err error) {
	// 获取站点ID (这里需要根据实际情况获取，可能从上下文或配置中获取)
	siteId := 1 // 临时硬编码，实际应该从请求中获取

	// 验证参数
	if req.Username == "" || req.Password == "" || req.Nickname == "" {
		return nil, fmt.Errorf("用户名、密码和昵称不能为空")
	}

	// 验证用户名长度和格式
	if len(req.Username) < 4 || len(req.Username) > 12 {
		return nil, fmt.Errorf("用户名长度必须在4-12个字符之间")
	}

	// 验证密码长度
	if len(req.Password) < 6 || len(req.Password) > 20 {
		return nil, fmt.Errorf("密码长度必须在6-20个字符之间")
	}

	// 验证昵称长度
	if len(req.Nickname) < 2 || len(req.Nickname) > 20 {
		return nil, fmt.Errorf("昵称长度必须在2-20个字符之间")
	}

	// 检查用户名是否已存在
	var existingAdmin *entity.Admin
	err = dao.Admin.Ctx(ctx).Where(do.Admin{
		Username: req.Username,
		SiteId:   siteId,
	}).Scan(&existingAdmin)

	if err != nil {
		return nil, fmt.Errorf("数据库查询错误: %v", err)
	}

	if existingAdmin != nil {
		return nil, fmt.Errorf("用户名已经被使用")
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("密码加密失败: %v", err)
	}

	// 创建管理员
	g.Log().Infof(ctx, "创建管理员 - 用户名: %s, 昵称: %s, 角色: %d, 状态: %d",
		req.Username, req.Nickname, req.Role, req.Status)

	// 确保状态值正确 (如果为0则设为1)
	status := int(req.Status)
	if status == 0 {
		status = 1
	}

	_, err = dao.Admin.Ctx(ctx).Insert(do.Admin{
		SiteId:      siteId,
		Username:    req.Username,
		Nickname:    req.Nickname,
		Password:    string(hashedPassword),
		AdminRoleId: int(req.Role),
		Status:      status,
		CreatedAt:   gtime.Now(),
		UpdatedAt:   gtime.Now(),
	})

	if err != nil {
		g.Log().Errorf(ctx, "创建管理员失败: %v", err)
		return nil, fmt.Errorf("创建管理员失败: %v", err)
	}

	g.Log().Infof(ctx, "创建管理员成功 - 用户名: %s, 实际状态: %d", req.Username, status)

	// 记录操作日志
	// 这里需要获取当前操作的管理员信息，暂时跳过
	// err = addAdminLog(ctx, currentAdmin, "添加员工："+req.Username)

	res = &v1.CreateAdminRes{}
	return res, nil
}

// HTTP 处理函数
func (c *Controller) LoginHTTP(r *ghttp.Request) {
	var req v1.LoginReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 400, "msg": "参数错误"})
		return
	}

	res, err := c.Login(r.Context(), &req)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 403, "msg": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{"code": 0, "msg": "登录成功", "data": res})
}

func (c *Controller) RefreshTokenHTTP(r *ghttp.Request) {
	var req v1.RefreshTokenReq

	res, err := c.RefreshToken(r.Context(), &req)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 666, "msg": "刷新安全令牌失败"})
		return
	}

	r.Response.WriteJson(g.Map{"code": 0, "msg": "刷新安全令牌成功", "data": res.Token})
}

func (c *Controller) CreateAdminHTTP(r *ghttp.Request) {
	var req v1.CreateAdminReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 401, "msg": "参数错误"})
		return
	}

	res, err := c.CreateAdmin(r.Context(), &req)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{"code": 0, "msg": "添加员工成功", "data": res})
}

// 辅助函数

// generateJWTToken 生成JWT token
func generateJWTToken(admin *entity.Admin) (string, error) {
	// JWT密钥 (实际应该从配置文件获取)
	jwtSecret := []byte("your-secret-key")

	// 创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin_id": admin.Id,
		"username": admin.Username,
		"site_id":  admin.SiteId,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 24小时过期
		"iat":      time.Now().Unix(),
	})

	// 签名token
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// getClientIP 获取客户端IP
func getClientIP(ctx context.Context) string {
	// 从上下文中获取HTTP请求
	if r := g.RequestFromCtx(ctx); r != nil {
		return r.GetClientIp()
	}
	return "127.0.0.1"
}

// addAdminLog 添加管理员日志
func addAdminLog(ctx context.Context, admin *entity.Admin, message string) error {
	_, err := dao.AdminLog.Ctx(ctx).Insert(do.AdminLog{
		SiteId:        admin.SiteId,
		AdminId:       int(admin.Id),
		AdminUsername: admin.Username,
		Ip:            getClientIP(ctx),
		Remark:        message,
		CreatedAt:     gtime.Now(),
	})
	return err
}
