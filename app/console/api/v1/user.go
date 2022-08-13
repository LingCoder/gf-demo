package v1

import (
	"gf-demo/app/console/internal/model"
	"gf-demo/library/base"
	"github.com/gogf/gf/v2/frame/g"
)

type UserDetailReq struct {
	g.Meta `tags:"User" summary:"View User Info" dc:"This is user detail information"`
}

type UserDetailRes struct {
	base.Res
	model.UserDetailItem
}
