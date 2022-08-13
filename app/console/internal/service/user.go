package service

import (
	"context"
	"gf-demo/app/console/internal/model"
	"gf-demo/app/console/internal/service/internal"
)

var localUser = internal.UserImpl{}

func User() IUser {
	return &localUser
}

type IUser interface {
	GetUserInfo(c context.Context, in model.UserDetailInput) (out *model.UserDetailOutput, err error)
}
