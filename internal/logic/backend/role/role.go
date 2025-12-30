package role

import (
	"context"
	"fmt"
	"jh_admin_service/api/backend/role/v1"
	"jh_admin_service/internal/service/backend"

	"jh_admin_service/internal/dao"
	"jh_admin_service/internal/middleware"
	"jh_admin_service/internal/model/do"
	"jh_admin_service/internal/model/entity"
)

type (
	sRole struct{}
)

func init() {
	backend.RegisterRole(&sRole{})
}

// GetRoleList 获取角色列表
func (s *sRole) GetRoleList(ctx context.Context, req *v1.GetRoleListReq) (*v1.GetRoleListRes, error) {
	middleware.LogWithTrace(ctx, "info", "获取角色列表请求 - SiteId: %d", req.SiteId)

	// 默认站点ID为1，如果请求中有指定则使用指定的
	siteId := int32(1)
	if req.SiteId > 0 {
		siteId = req.SiteId
	}

	// 查询角色列表
	var roles []*entity.AdminRole
	err := dao.AdminRole.Ctx(ctx).Where(do.AdminRole{
		SiteId: siteId,
		Status: 1, // 只查询启用的角色
	}).OrderAsc(dao.AdminRole.Columns().Id).Scan(&roles)

	if err != nil {
		middleware.LogWithTrace(ctx, "error", "查询角色列表失败: %v", err)
		return nil, fmt.Errorf("查询角色列表失败: %v", err)
	}

	// 构建响应数据
	var roleList []*v1.RoleInfo
	for _, role := range roles {
		roleInfo := &v1.RoleInfo{
			Id:          int32(role.Id),
			SiteId:      int32(role.SiteId),
			Name:        role.Name,
			Status:      int32(role.Status),
			Permissions: role.Permissions,
		}

		// 处理时间字段
		if role.CreatedAt != nil {
			roleInfo.CreatedAt = role.CreatedAt.Unix()
		}
		if role.UpdatedAt != nil {
			roleInfo.UpdatedAt = role.UpdatedAt.Unix()
		}

		roleList = append(roleList, roleInfo)
	}

	middleware.LogWithTrace(ctx, "info", "获取角色列表成功 - SiteId: %d, 角色数量: %d", siteId, len(roleList))

	return &v1.GetRoleListRes{
		Roles: roleList,
	}, nil
}

// CreateRole 创建角色
func (s *sRole) CreateRole(ctx context.Context, req *v1.CreateRoleReq) (*v1.CreateRoleRes, error) {
	middleware.LogWithTrace(ctx, "info", "创建角色请求 - SiteId: %d, Name: %s", req.SiteId, req.Name)

	// 参数验证
	if req.Name == "" {
		return &v1.CreateRoleRes{
			Success: false,
			Message: "角色名称不能为空",
		}, nil
	}

	if len(req.Name) < 2 || len(req.Name) > 50 {
		return &v1.CreateRoleRes{
			Success: false,
			Message: "角色名称长度必须在2-50个字符之间",
		}, nil
	}

	// 默认站点ID为1
	siteId := int32(1)
	if req.SiteId > 0 {
		siteId = req.SiteId
	}

	// 检查角色名称是否已存在
	count, err := dao.AdminRole.Ctx(ctx).Where(do.AdminRole{
		SiteId: siteId,
		Name:   req.Name,
	}).Count()
	if err != nil {
		middleware.LogWithTrace(ctx, "error", "检查角色名称重复失败: %v", err)
		return &v1.CreateRoleRes{
			Success: false,
			Message: "系统错误，请稍后重试",
		}, nil
	}

	if count > 0 {
		return &v1.CreateRoleRes{
			Success: false,
			Message: "角色名称已存在",
		}, nil
	}

	// 创建角色
	_, err = dao.AdminRole.Ctx(ctx).Data(do.AdminRole{
		SiteId:      siteId,
		Name:        req.Name,
		Status:      1, // 默认启用
		Permissions: "",
	}).Insert()

	if err != nil {
		middleware.LogWithTrace(ctx, "error", "创建角色失败: %v", err)
		return &v1.CreateRoleRes{
			Success: false,
			Message: "创建角色失败",
		}, nil
	}

	middleware.LogWithTrace(ctx, "info", "创建角色成功 - SiteId: %d, Name: %s", siteId, req.Name)

	return &v1.CreateRoleRes{
		Success: true,
		Message: "创建成功",
	}, nil
}

// UpdateRole 更新角色
func (s *sRole) UpdateRole(ctx context.Context, req *v1.UpdateRoleReq) (*v1.UpdateRoleRes, error) {
	middleware.LogWithTrace(ctx, "info", "更新角色请求 - Id: %d, SiteId: %d, Name: %s", req.Id, req.SiteId, req.Name)

	// 参数验证
	if req.Id <= 0 {
		return &v1.UpdateRoleRes{
			Success: false,
			Message: "角色ID无效",
		}, nil
	}

	if req.Name == "" {
		return &v1.UpdateRoleRes{
			Success: false,
			Message: "角色名称不能为空",
		}, nil
	}

	if len(req.Name) < 2 || len(req.Name) > 50 {
		return &v1.UpdateRoleRes{
			Success: false,
			Message: "角色名称长度必须在2-50个字符之间",
		}, nil
	}

	// 默认站点ID为1
	siteId := int32(1)
	if req.SiteId > 0 {
		siteId = req.SiteId
	}

	// 检查角色是否存在
	var role *entity.AdminRole
	err := dao.AdminRole.Ctx(ctx).Where(do.AdminRole{
		Id:     req.Id,
		SiteId: siteId,
	}).Scan(&role)
	if err != nil {
		middleware.LogWithTrace(ctx, "error", "查询角色失败: %v", err)
		return &v1.UpdateRoleRes{
			Success: false,
			Message: "系统错误，请稍后重试",
		}, nil
	}

	if role == nil {
		return &v1.UpdateRoleRes{
			Success: false,
			Message: "角色不存在",
		}, nil
	}

	// 检查角色名称是否已被其他角色使用
	count, err := dao.AdminRole.Ctx(ctx).Where(do.AdminRole{
		SiteId: siteId,
		Name:   req.Name,
	}).WhereNot(dao.AdminRole.Columns().Id, req.Id).Count()
	if err != nil {
		middleware.LogWithTrace(ctx, "error", "检查角色名称重复失败: %v", err)
		return &v1.UpdateRoleRes{
			Success: false,
			Message: "系统错误，请稍后重试",
		}, nil
	}

	if count > 0 {
		return &v1.UpdateRoleRes{
			Success: false,
			Message: "角色名称已存在",
		}, nil
	}

	// 更新角色
	_, err = dao.AdminRole.Ctx(ctx).Where(do.AdminRole{
		Id:     req.Id,
		SiteId: siteId,
	}).Data(do.AdminRole{
		Name: req.Name,
	}).Update()

	if err != nil {
		middleware.LogWithTrace(ctx, "error", "更新角色失败: %v", err)
		return &v1.UpdateRoleRes{
			Success: false,
			Message: "更新角色失败",
		}, nil
	}

	middleware.LogWithTrace(ctx, "info", "更新角色成功 - Id: %d, Name: %s", req.Id, req.Name)

	return &v1.UpdateRoleRes{
		Success: true,
		Message: "更新成功",
	}, nil
}

// DeleteRole 删除角色
func (s *sRole) DeleteRole(ctx context.Context, req *v1.DeleteRoleReq) (*v1.DeleteRoleRes, error) {
	middleware.LogWithTrace(ctx, "info", "删除角色请求 - Id: %d, SiteId: %d", req.Id, req.SiteId)

	// 参数验证
	if req.Id <= 0 {
		return &v1.DeleteRoleRes{
			Success: false,
			Message: "角色ID无效",
		}, nil
	}

	// 默认站点ID为1
	siteId := int32(1)
	if req.SiteId > 0 {
		siteId = req.SiteId
	}

	// 检查角色是否存在
	var role *entity.AdminRole
	err := dao.AdminRole.Ctx(ctx).Where(do.AdminRole{
		Id:     req.Id,
		SiteId: siteId,
	}).Scan(&role)
	if err != nil {
		middleware.LogWithTrace(ctx, "error", "查询角色失败: %v", err)
		return &v1.DeleteRoleRes{
			Success: false,
			Message: "系统错误，请稍后重试",
		}, nil
	}

	if role == nil {
		return &v1.DeleteRoleRes{
			Success: false,
			Message: "角色不存在",
		}, nil
	}

	// 删除角色
	_, err = dao.AdminRole.Ctx(ctx).Where(do.AdminRole{
		Id:     req.Id,
		SiteId: siteId,
	}).Delete()

	if err != nil {
		middleware.LogWithTrace(ctx, "error", "删除角色失败: %v", err)
		return &v1.DeleteRoleRes{
			Success: false,
			Message: "删除角色失败",
		}, nil
	}

	middleware.LogWithTrace(ctx, "info", "删除角色成功 - Id: %d, Name: %s", req.Id, role.Name)

	return &v1.DeleteRoleRes{
		Success: true,
		Message: "删除成功",
	}, nil
}
