package site

import (
	"context"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	v1 "jh_user_service/api/site/v1"
	"jh_user_service/internal/service"
)

type Controller struct {
	v1.UnimplementedSiteServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterSiteServer(s.Server, &Controller{})
}

// GetBasicSetting 获取站点基本设置 (gRPC)
func (*Controller) GetBasicSetting(ctx context.Context, req *v1.GetBasicSettingReq) (res *v1.GetBasicSettingRes, err error) {
	return service.Site().GetBasicSetting(ctx, req)
}

// UpdateBasicSetting 更新站点基本设置 (gRPC)
func (*Controller) UpdateBasicSetting(ctx context.Context, req *v1.UpdateBasicSettingReq) (res *v1.UpdateBasicSettingRes, err error) {
	return service.Site().UpdateBasicSetting(ctx, req)
}

//// GetBasicSettingHTTP 获取站点基本设置 (HTTP)
//func (c *Controller) GetBasicSettingHTTP(r *ghttp.Request) {
//	siteId := r.Get("site_id", 1).Int32()
//	req := &v1.GetBasicSettingReq{
//		SiteId: siteId,
//	}
//
//	res, err := c.GetBasicSetting(r.Context(), req)
//	if err != nil {
//		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
//		return
//	}
//
//	r.Response.WriteJson(g.Map{"code": 0, "msg": "获取数据成功", "data": res})
//}
//
//// UpdateBasicSettingHTTP 更新站点基本设置 (HTTP)
//func (c *Controller) UpdateBasicSettingHTTP(r *ghttp.Request) {
//	var req v1.UpdateBasicSettingReq
//	if err := r.Parse(&req); err != nil {
//		r.Response.WriteJson(g.Map{"code": 400, "msg": "参数解析失败: " + err.Error()})
//		return
//	}
//
//	// 如果请求中没有指定站点ID，使用默认值
//	if req.SiteId == 0 {
//		req.SiteId = 1
//	}
//
//	res, err := c.UpdateBasicSetting(r.Context(), &req)
//	if err != nil {
//		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
//		return
//	}
//
//	r.Response.WriteJson(g.Map{"code": 0, "msg": res.Message})
//}
