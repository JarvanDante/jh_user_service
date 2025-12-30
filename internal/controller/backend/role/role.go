package role

import (
	"context"
	v2 "jh_admin_service/api/backend/role/v1"
	"jh_admin_service/internal/service/backend"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v2.UnimplementedRoleServer
}

// Register 注册角色服务
func Register(s *grpcx.GrpcServer) {
	v2.RegisterRoleServer(s.Server, &Controller{})
}

// GetRoleList 获取角色列表
func (*Controller) GetRoleList(ctx context.Context, req *v2.GetRoleListReq) (*v2.GetRoleListRes, error) {
	return backend.Role().GetRoleList(ctx, req)
}

// CreateRole 创建角色
func (*Controller) CreateRole(ctx context.Context, req *v2.CreateRoleReq) (*v2.CreateRoleRes, error) {
	return backend.Role().CreateRole(ctx, req)
}

// UpdateRole 更新角色
func (*Controller) UpdateRole(ctx context.Context, req *v2.UpdateRoleReq) (*v2.UpdateRoleRes, error) {
	return backend.Role().UpdateRole(ctx, req)
}

// DeleteRole 删除角色
func (*Controller) DeleteRole(ctx context.Context, req *v2.DeleteRoleReq) (*v2.DeleteRoleRes, error) {
	return backend.Role().DeleteRole(ctx, req)
}
