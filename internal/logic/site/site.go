package site

import (
	"context"
	"fmt"

	v1 "jh_user_service/api/site/v1"
	"jh_user_service/internal/dao"
	"jh_user_service/internal/middleware"
	"jh_user_service/internal/model/do"
	"jh_user_service/internal/model/entity"
	"jh_user_service/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type (
	sSite struct{}
)

func init() {
	service.RegisterSite(&sSite{})
}

// GetBasicSetting 获取站点基本设置
func (s *sSite) GetBasicSetting(ctx context.Context, req *v1.GetBasicSettingReq) (*v1.GetBasicSettingRes, error) {
	middleware.LogWithTrace(ctx, "info", "获取站点基本设置请求 - SiteId: %d", req.SiteId)

	// 默认站点ID为1，如果请求中有指定则使用指定的
	siteId := int32(1)
	if req.SiteId > 0 {
		siteId = req.SiteId
	}

	// 查询站点配置
	var siteConfig *entity.SiteConfig
	err := dao.SiteConfig.Ctx(ctx).Where(do.SiteConfig{
		SiteId: siteId,
	}).Scan(&siteConfig)

	if err != nil {
		middleware.LogWithTrace(ctx, "error", "查询站点配置失败: %v", err)
		return nil, fmt.Errorf("查询站点配置失败: %v", err)
	}

	if siteConfig == nil {
		middleware.LogWithTrace(ctx, "error", "站点配置不存在 - SiteId: %d", siteId)
		return nil, fmt.Errorf("站点配置不存在")
	}

	// 构建响应数据
	res := &v1.GetBasicSettingRes{
		Code:                 fmt.Sprintf("site_%d", siteId), // 站点代码
		Name:                 fmt.Sprintf("站点_%d", siteId),   // 站点名称，这里可以从其他表获取
		RegisterTimeInterval: int32(siteConfig.RegisterTimeInterval),
		SwitchRegister:       siteConfig.SwitchRegister == 1,
		IsClose:              siteConfig.IsClose == 1,
		CloseReason:          siteConfig.CloseReason,
		ServiceUrl:           siteConfig.UrlService,
		AgentUrl:             siteConfig.UrlAgentPc,
		MobileUrl:            siteConfig.UrlMobile,
		AgentRegisterUrl:     siteConfig.UrlAgentRegister,
		MinWithdraw:          siteConfig.MinWithdraw,
		MaxWithdraw:          siteConfig.MaxWithdraw,
		MobileLogo:           siteConfig.MobileLogo,
		DefaultAgentId:       int32(siteConfig.DefaultAgentId),
		DefaultAgentName:     "", // 这里可以根据 DefaultAgentId 查询代理名称
		Balance:              0,  // 这里可以查询站点余额
		BalanceReset:         0,  // 这里可以查询剩余额度
	}

	// 如果关闭原因为空，设置默认消息
	if res.CloseReason == "" {
		res.CloseReason = "您好，系统正在维护中，估计需要1个小时左右，给您造成不便，敬请谅解！"
	}

	middleware.LogWithTrace(ctx, "info", "获取站点基本设置成功 - SiteId: %d", siteId)
	return res, nil
}

// UpdateBasicSetting 更新站点基本设置
func (s *sSite) UpdateBasicSetting(ctx context.Context, req *v1.UpdateBasicSettingReq) (*v1.UpdateBasicSettingRes, error) {
	middleware.LogWithTrace(ctx, "info", "更新站点基本设置请求 - SiteId: %d", req.SiteId)

	// 默认站点ID为1，如果请求中有指定则使用指定的
	siteId := int32(1)
	if req.SiteId > 0 {
		siteId = req.SiteId
	}

	// 参数验证
	if err := s.validateUpdateBasicSetting(req); err != nil {
		middleware.LogWithTrace(ctx, "error", "参数验证失败: %v", err)
		return nil, fmt.Errorf("参数验证失败: %v", err)
	}

	// 查询现有配置
	var existingConfig *entity.SiteConfig
	err := dao.SiteConfig.Ctx(ctx).Where(do.SiteConfig{
		SiteId: siteId,
	}).Scan(&existingConfig)

	if err != nil {
		middleware.LogWithTrace(ctx, "error", "查询现有配置失败: %v", err)
		return nil, fmt.Errorf("查询现有配置失败: %v", err)
	}

	// 准备更新数据
	updateData := do.SiteConfig{
		RegisterTimeInterval: req.IpRegisterTime,
		SwitchRegister:       g.Map{"true": 1, "false": 0}[fmt.Sprintf("%t", req.OpenRegister)],
		IsClose:              g.Map{"true": 1, "false": 0}[fmt.Sprintf("%t", req.Close)],
		CloseReason:          req.CloseReason,
		UrlAgentPc:           req.AgentUrl,
		UrlMobile:            req.MobileUrl,
		UrlAgentRegister:     req.AgentRegisterUrl,
		UrlService:           req.ServiceUrl,
		MobileLogo:           req.MobileLogo,
	}

	// 设置提现金额，确保最小值
	if req.MinWithdraw < 1 {
		updateData.MinWithdraw = 1
	} else {
		updateData.MinWithdraw = req.MinWithdraw
	}

	if req.MaxWithdraw < 1 {
		updateData.MaxWithdraw = 9999999
	} else {
		updateData.MaxWithdraw = req.MaxWithdraw
	}

	// 如果配置不存在，创建新配置
	if existingConfig == nil {
		updateData.SiteId = siteId
		_, err = dao.SiteConfig.Ctx(ctx).Data(updateData).Insert()
		if err != nil {
			middleware.LogWithTrace(ctx, "error", "创建站点配置失败: %v", err)
			return nil, fmt.Errorf("创建站点配置失败: %v", err)
		}
	} else {
		// 更新现有配置
		_, err = dao.SiteConfig.Ctx(ctx).Where(do.SiteConfig{
			SiteId: siteId,
		}).Data(updateData).Update()

		if err != nil {
			middleware.LogWithTrace(ctx, "error", "更新站点配置失败: %v", err)
			return nil, fmt.Errorf("更新站点配置失败: %v", err)
		}
	}

	middleware.LogWithTrace(ctx, "info", "更新站点基本设置成功 - SiteId: %d", siteId)
	return &v1.UpdateBasicSettingRes{
		Message: "设置成功",
	}, nil
}

// validateUpdateBasicSetting 验证更新基本设置的参数
func (s *sSite) validateUpdateBasicSetting(req *v1.UpdateBasicSettingReq) error {
	// 验证URL格式（简单验证）
	if req.AgentUrl != "" && len(req.AgentUrl) < 7 {
		return fmt.Errorf("代理链接地址格式错误")
	}

	if req.MobileUrl != "" && len(req.MobileUrl) < 7 {
		return fmt.Errorf("手机域名地址格式错误")
	}

	if req.AgentRegisterUrl != "" && len(req.AgentRegisterUrl) < 7 {
		return fmt.Errorf("代理推广地址格式错误")
	}

	// 验证提现金额
	if req.MinWithdraw < 0 {
		return fmt.Errorf("单笔最低提现金额不能为负数")
	}

	if req.MaxWithdraw < 0 {
		return fmt.Errorf("单笔最高提现金额不能为负数")
	}

	if req.MinWithdraw > req.MaxWithdraw && req.MaxWithdraw > 0 {
		return fmt.Errorf("单笔最低提现金额不能大于最高提现金额")
	}

	// 验证关闭原因长度
	if len(req.CloseReason) > 255 {
		return fmt.Errorf("暂时关闭网站原因不能超过255个字符")
	}

	return nil
}
