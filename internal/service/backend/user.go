package backend

import (
	"context"
	v1 "jh_admin_service/api/backend/user/v1"
)

type (
	IUser interface {
		GetUserList(ctx context.Context, req *v1.GetUserListReq) (*v1.GetUserListRes, error)
		UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (*v1.UpdateUserRes, error)
		GetUserBasicInfo(ctx context.Context, req *v1.GetUserBasicInfoReq) (*v1.GetUserBasicInfoRes, error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
