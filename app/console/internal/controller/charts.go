package controller

import (
	"context"
	v1 "gf-demo/app/console/api/v1"
	"gf-demo/app/console/internal/model"
	"gf-demo/app/console/internal/service"
)

var Charts = cCharts{}

type cCharts struct {
	cBase
}

// RewardKLine show KLine data
func (c *cCharts) RewardKLine(ctx context.Context, req *v1.ChartsKlineReq) (res *v1.ChartsKlineRes, err error) {
	out, err := service.Charts().GetKline(ctx, model.ChartsKlineInput{})
	if err != nil {
		return nil, err
	}
	res.List = out.List
	return
}
