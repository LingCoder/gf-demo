package internal

import (
	"context"
	"gf-demo/app/console/internal/model"
)

type UserImpl struct{}

func (p *UserImpl) GetUserInfo(c context.Context, in model.UserDetailInput) (out *model.UserDetailOutput, err error) {

	return out, nil
}
