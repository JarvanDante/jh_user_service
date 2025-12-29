package role

import (
	"context"
	v1 "jh_user_service/api/role/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedRoleServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterRoleServer(s.Server, &Controller{})
}

func (*Controller) GetRoleList(ctx context.Context, req *v1.GetRoleListReq) (res *v1.GetRoleListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
