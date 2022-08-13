package controller

import (
	"context"
	v1 "gf-demo/app/console/api/v1"
	"gf-demo/app/console/internal/model"
	"gf-demo/app/console/internal/service"
)

var User = cUser{}

type cUser struct {
	cBase
}

// UserDetail show page data
func (c *cUser) UserDetail(ctx context.Context, req *v1.UserDetailReq) (res *v1.UserDetailRes, err error) {
	out, err := service.User().GetUserInfo(ctx, model.UserDetailInput{})
	if err != nil {
		return nil, err
	}
	res.UserDetailItem = out.UserDetailItem
	return
}
