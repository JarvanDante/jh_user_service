package site

import (
	"context"
	v1 "jh_admin_service/api/backend/site/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedSiteServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterSiteServer(s.Server, &Controller{})
}

func (*Controller) GetBasicSetting(ctx context.Context, req *v1.GetBasicSettingReq) (res *v1.GetBasicSettingRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) UpdateBasicSetting(ctx context.Context, req *v1.UpdateBasicSettingReq) (res *v1.UpdateBasicSettingRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
