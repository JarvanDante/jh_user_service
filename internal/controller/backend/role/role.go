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
