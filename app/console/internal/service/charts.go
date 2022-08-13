package service

import (
	"context"
	"gf-demo/app/console/internal/model"
	"gf-demo/app/console/internal/service/internal"
)

var localCharts = internal.ChartImpl{}

func Charts() ICharts {
	return &localCharts
}

type ICharts interface {
	GetKline(ctx context.Context, in model.ChartsKlineInput) (out *model.ChartsKlineOutput, err error)
}
