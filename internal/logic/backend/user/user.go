package user

import (
	"context"
	"fmt"
	v1 "jh_admin_service/api/backend/user/v1"
	"jh_admin_service/internal/service/backend"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/crypto/bcrypt"

	"jh_admin_service/internal/dao"
	"jh_admin_service/internal/middleware"
	"jh_admin_service/internal/model/do"
	"jh_admin_service/internal/model/entity"
	"jh_admin_service/internal/tracing"
)

type (
	sUser struct{}
)

func init() {
	backend.RegisterUser(&sUser{})
}

// GetUserList 获取用户列表
func (s *sUser) GetUserList(ctx context.Context, req *v1.GetUserListReq) (*v1.GetUserListRes, error) {
	// 参数验证
	if req == nil {
		return nil, fmt.Errorf("请求参数不能为空")
	}

	// 创建Jaeger span
	ctx, span := tracing.StartSpan(ctx, "user.GetUserList", trace.WithAttributes(
		attribute.String("method", "GetUserList"),
		attribute.String("username", req.Username),
		attribute.Int("page", int(req.Page)),
		attribute.Int("size", int(req.Size)),
	))
	defer span.End()

	// 获取站点ID (从配置中获取，如果没有则尝试查找)
	siteId := 1 // 默认站点ID

	// 尝试从配置获取站点代码对应的ID
	siteCode := g.Cfg().MustGet(ctx, "site.code", "site_1").String()
	middleware.LogWithTrace(ctx, "info", fmt.Sprintf("配置的站点代码: %s", siteCode))

	tracing.SetSpanAttributes(span, attribute.Int("site_id", siteId))

	middleware.LogWithTrace(ctx, "info", fmt.Sprintf("开始查询用户列表，站点ID: %d, 页码: %d, 每页: %d", siteId, req.Page, req.Size))

	// 设置默认分页参数
	page := req.Page
	size := req.Size
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 20
	}

	// 构建查询条件
	query := dao.User.Ctx(ctx).Where("site_id = ?", siteId)

	middleware.LogWithTrace(ctx, "info", fmt.Sprintf("基础查询条件 - 站点ID: %d", siteId))

	// 添加筛选条件
	if req.Status > 0 {
		query = query.Where("status = ?", req.Status)
	}
	if req.GradeId > 0 {
		query = query.Where("grade_id = ?", req.GradeId)
	}
	if req.LevelId > 0 {
		query = query.Where("level_id = ?", req.LevelId)
	}
	if req.Username != "" {
		query = query.Where("username = ?", req.Username)
	}
	if req.Realname != "" {
		query = query.Where("realname = ?", req.Realname)
	}
	if req.Mobile != "" {
		query = query.Where("mobile = ?", req.Mobile)
	}

	// 代理用户名筛选
	if req.AgentUsername != "" {
		// 先查找代理ID
		var agent entity.Agent
		err := dao.Agent.Ctx(ctx).Where("site_id = ? AND username = ?", siteId, req.AgentUsername).Scan(&agent)
		if err != nil {
			return &v1.GetUserListRes{List: []*v1.UserInfo{}, Count: 0}, nil
		}
		query = query.Where("agent_id = ?", agent.Id)
	}

	// 域名筛选
	if req.Domain != "" {
		query = query.WhereLike("register_url", "%"+req.Domain+"%")
	}

	// 时间范围筛选
	if req.StartDate != "" {
		query = query.Where("created_at >= ?", req.StartDate)
	}
	if req.EndDate != "" {
		query = query.Where("created_at <= ?", req.EndDate)
	}

	// 数据库查询span - 获取总数
	ctx, countSpan := tracing.StartSpan(ctx, "db.query.user_count", trace.WithAttributes(
		attribute.String("db.operation", "count"),
		attribute.String("db.table", "user"),
	))
	total, err := query.Count()
	countSpan.End()
	if err != nil {
		tracing.SetSpanError(span, err)
		return nil, fmt.Errorf("获取用户总数失败: %v", err)
	}

	middleware.LogWithTrace(ctx, "info", fmt.Sprintf("查询到用户总数: %d", total))

	// 添加调试：检查数据库中是否有任何用户数据
	allUsersCount, _ := dao.User.Ctx(ctx).Count()
	middleware.LogWithTrace(ctx, "info", fmt.Sprintf("数据库中所有用户总数: %d", allUsersCount))

	// 检查不同站点ID的用户数量
	for i := 1; i <= 3; i++ {
		siteCount, _ := dao.User.Ctx(ctx).Where("site_id = ?", i).Count()
		middleware.LogWithTrace(ctx, "info", fmt.Sprintf("站点ID %d 的用户数量: %d", i, siteCount))
	}

	// 数据库查询span - 获取列表数据
	ctx, listSpan := tracing.StartSpan(ctx, "db.query.user_list", trace.WithAttributes(
		attribute.String("db.operation", "select"),
		attribute.String("db.table", "user"),
		attribute.Int("limit", int(size)),
		attribute.Int("offset", int((page-1)*size)),
	))

	// 排序处理
	orderBy := "created_at DESC"
	if req.SortField != "" {
		orderField := s.getOrderByField(req.SortField)
		orderRule := "ASC"
		if req.SortRule == 0 {
			orderRule = "DESC"
		}
		orderBy = fmt.Sprintf("%s %s", orderField, orderRule)
	}

	var users []entity.User
	err = query.Page(int(page), int(size)).
		Order(orderBy).
		Scan(&users)
	listSpan.End()

	if err != nil {
		tracing.SetSpanError(span, err)
		return nil, fmt.Errorf("获取用户列表失败: %v", err)
	}

	// 转换为响应格式
	var userList []*v1.UserInfo
	for _, user := range users {
		userInfo := &v1.UserInfo{
			Id:            int32(user.Id),
			Username:      user.Username,
			GradeId:       int32(user.GradeId),
			LevelId:       int32(user.LevelId),
			AgentId:       int32(user.AgentId),
			Status:        int32(user.Status),
			RegisterIp:    user.RegisterIp,
			Realname:      user.Realname,
			Mobile:        s.maskMobile(user.Mobile),
			Email:         user.Email,
			FocusLevel:    int32(user.FocusLevel),
			BalanceStatus: int32(user.BalanceStatus),
			IsOnline:      int32(user.IsOnline),
			PayTimes:      int32(user.PayTimes),
		}

		// 时间格式化
		if user.RegisterTime != nil {
			userInfo.RegisterTime = user.RegisterTime.Format("2006-01-02 15:04:05")
		}
		if user.LastLoginTime != nil {
			userInfo.LastLoginTime = user.LastLoginTime.Format("2006-01-02 15:04:05")
		}

		// 获取等级名称
		var grade entity.UserGrade
		dao.UserGrade.Ctx(ctx).Where(do.UserGrade{
			SiteId: siteId,
			Id:     uint(user.GradeId),
		}).Scan(&grade)
		userInfo.GradeName = grade.Name

		// 获取层级名称
		var level entity.UserLevel
		dao.UserLevel.Ctx(ctx).Where(do.UserLevel{
			SiteId: siteId,
			Id:     uint(user.LevelId),
		}).Scan(&level)
		userInfo.LevelName = level.Name

		// 获取代理用户名
		var agent entity.Agent
		dao.Agent.Ctx(ctx).Where(do.Agent{
			SiteId: siteId,
			Id:     uint(user.AgentId),
		}).Scan(&agent)
		userInfo.AgentUsername = agent.Username

		userList = append(userList, userInfo)
	}

	// 确保返回空数组而不是nil
	if userList == nil {
		userList = []*v1.UserInfo{}
	}

	// 获取统计数据
	totalUsers, _ := dao.User.Ctx(ctx).Where("site_id = ?", siteId).Count()
	totalChargeUsers, _ := dao.User.Ctx(ctx).Where("site_id = ? AND pay_times >= 1", siteId).Count()

	tracing.SetSpanAttributes(span,
		attribute.Int("total_count", total),
		attribute.Int("returned_count", len(userList)),
	)

	middleware.LogWithTrace(ctx, "info", fmt.Sprintf("获取用户列表成功，总数: %d，返回: %d", total, len(userList)))

	return &v1.GetUserListRes{
		List:             userList,
		Count:            int32(total),
		TotalUsers:       int32(totalUsers),
		TotalChargeUsers: int32(totalChargeUsers),
	}, nil
}

// UpdateUser 更新用户信息
func (s *sUser) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (*v1.UpdateUserRes, error) {
	// 参数验证
	if req == nil {
		return nil, fmt.Errorf("请求参数不能为空")
	}
	if req.Id <= 0 {
		return &v1.UpdateUserRes{Success: false, Message: "用户ID不能为空"}, nil
	}

	// 创建Jaeger span
	ctx, span := tracing.StartSpan(ctx, "user.UpdateUser", trace.WithAttributes(
		attribute.String("method", "UpdateUser"),
		attribute.Int("user_id", int(req.Id)),
	))
	defer span.End()

	// 获取站点ID
	siteId := 1 // 临时硬编码

	// 查找用户
	var user entity.User
	err := dao.User.Ctx(ctx).Where("site_id = ? AND id = ?", siteId, req.Id).Scan(&user)

	if err != nil || user.Id == 0 {
		return &v1.UpdateUserRes{Success: false, Message: "用户不存在"}, nil
	}

	// 构建更新数据
	updateData := do.User{
		GradeId:       int(req.GradeId),
		LevelId:       int(req.LevelId),
		AgentId:       int(req.AgentId),
		Sex:           int(req.Sex),
		FocusLevel:    int(req.FocusLevel),
		BalanceStatus: uint(req.BalanceStatus),
		Status:        int(req.Status),
		UpdatedAt:     gtime.Now(),
	}

	// 可选字段更新
	if req.LoginPassword != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.LoginPassword), bcrypt.DefaultCost)
		if err != nil {
			return &v1.UpdateUserRes{Success: false, Message: "密码加密失败"}, nil
		}
		updateData.Password = string(hashedPassword)
	}

	if req.PayPassword != "" {
		hashedPayPassword, err := bcrypt.GenerateFromPassword([]byte(req.PayPassword), bcrypt.DefaultCost)
		if err != nil {
			return &v1.UpdateUserRes{Success: false, Message: "资金密码加密失败"}, nil
		}
		updateData.PayPassword = string(hashedPayPassword)
	}

	if req.Realname != "" {
		updateData.Realname = req.Realname
	}
	if req.Mobile != "" {
		updateData.Mobile = req.Mobile
	}
	if req.Email != "" {
		updateData.Email = req.Email
	}
	if req.Qq != "" {
		updateData.Qq = req.Qq
	}
	if req.Birthday != "" {
		birthday, err := gtime.StrToTime(req.Birthday)
		if err == nil {
			updateData.Birthday = birthday
		}
	}
	if req.Remark != "" {
		updateData.Remark = req.Remark
	}

	// 执行更新
	_, err = dao.User.Ctx(ctx).Where(do.User{
		SiteId: siteId,
		Id:     uint(req.Id),
	}).Update(updateData)

	if err != nil {
		tracing.SetSpanError(span, err)
		return &v1.UpdateUserRes{Success: false, Message: "更新用户失败"}, nil
	}

	middleware.LogWithTrace(ctx, "info", fmt.Sprintf("更新用户成功，用户ID: %d", req.Id))

	return &v1.UpdateUserRes{Success: true, Message: "更新用户成功"}, nil
}

// GetUserBasicInfo 获取用户基本信息
func (s *sUser) GetUserBasicInfo(ctx context.Context, req *v1.GetUserBasicInfoReq) (*v1.GetUserBasicInfoRes, error) {
	// 参数验证
	if req == nil {
		return nil, fmt.Errorf("请求参数不能为空")
	}
	if req.Id <= 0 {
		return nil, fmt.Errorf("用户ID不能为空")
	}

	// 创建Jaeger span
	ctx, span := tracing.StartSpan(ctx, "user.GetUserBasicInfo", trace.WithAttributes(
		attribute.String("method", "GetUserBasicInfo"),
		attribute.Int("user_id", int(req.Id)),
	))
	defer span.End()

	// 获取站点ID
	siteId := 1 // 临时硬编码

	// 查找用户
	var user entity.User
	err := dao.User.Ctx(ctx).Where("site_id = ? AND id = ?", siteId, req.Id).Scan(&user)

	if err != nil || user.Id == 0 {
		return nil, fmt.Errorf("用户不存在")
	}

	// 构建基本信息
	basicInfo := &v1.UserBasicInfo{
		Id:            int32(user.Id),
		Username:      user.Username,
		Realname:      user.Realname,
		Mobile:        user.Mobile,
		Email:         user.Email,
		Sex:           int32(user.Sex),
		FocusLevel:    int32(user.FocusLevel),
		BalanceStatus: int32(user.BalanceStatus),
		Status:        int32(user.Status),
		Remark:        user.Remark,
		RegisterIp:    user.RegisterIp,
		LastLoginIp:   user.LastLoginIp,
		IsOnline:      int32(user.IsOnline),
		Qq:            user.Qq,
	}

	// 时间格式化
	if user.RegisterTime != nil {
		basicInfo.RegisterTime = user.RegisterTime.Format("2006-01-02 15:04:05")
	}
	if user.LastLoginTime != nil {
		basicInfo.LastLoginTime = user.LastLoginTime.Format("2006-01-02 15:04:05")
	}
	if user.Birthday != nil {
		basicInfo.Birthday = user.Birthday.Format("2006-01-02")
	}

	// 获取等级名称
	var grade entity.UserGrade
	dao.UserGrade.Ctx(ctx).Where(do.UserGrade{
		SiteId: siteId,
		Id:     uint(user.GradeId),
	}).Scan(&grade)
	basicInfo.GradeName = grade.Name

	// 获取层级名称
	var level entity.UserLevel
	dao.UserLevel.Ctx(ctx).Where(do.UserLevel{
		SiteId: siteId,
		Id:     uint(user.LevelId),
	}).Scan(&level)
	basicInfo.LevelName = level.Name

	// 获取代理名称
	var agent entity.Agent
	dao.Agent.Ctx(ctx).Where(do.Agent{
		SiteId: siteId,
		Id:     uint(user.AgentId),
	}).Scan(&agent)
	basicInfo.AgentName = agent.Username

	// 获取银行卡信息
	var banks []entity.UserBank
	dao.UserBank.Ctx(ctx).Where(do.UserBank{
		SiteId: siteId,
		UserId: int(user.Id),
	}).Scan(&banks)

	var bankList []*v1.BankInfo
	for _, bank := range banks {
		bankList = append(bankList, &v1.BankInfo{
			BankName: bank.BankName,
			CardNo:   bank.CardNo, // 注意：实际应用中可能需要解密
		})
	}
	basicInfo.Banks = bankList

	middleware.LogWithTrace(ctx, "info", fmt.Sprintf("获取用户基本信息成功，用户ID: %d", req.Id))

	return &v1.GetUserBasicInfoRes{User: basicInfo}, nil
}

// getOrderByField 获取排序字段
func (s *sUser) getOrderByField(sortField string) string {
	switch sortField {
	case "register_time":
		return "register_time"
	case "balance":
		return "balance"
	case "balance_frozen":
		return "balance_frozen"
	case "last_login_time":
		return "last_login_time"
	case "points":
		return "points"
	default:
		return "created_at"
	}
}

// maskMobile 手机号脱敏
func (s *sUser) maskMobile(mobile string) string {
	if len(mobile) < 8 {
		return mobile
	}
	return mobile[:3] + "****" + mobile[len(mobile)-4:]
}
