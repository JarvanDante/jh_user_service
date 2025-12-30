package admin

import (
	"context"
	v1 "jh_admin_service/api/backend/admin/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedAdminServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterAdminServer(s.Server, &Controller{})
}

func (*Controller) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) CreateAdmin(ctx context.Context, req *v1.CreateAdminReq) (res *v1.CreateAdminRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
