package internal

import (
	"context"
	"encoding/json"
	"gf-demo/app/console/internal/model"
	"gf-demo/app/console/internal/packed"
)

type ChartImpl struct {
}

func (p *ChartImpl) GetKline(c context.Context, in model.ChartsKlineInput) (out *model.ChartsKlineOutput, err error) {
	out = &model.ChartsKlineOutput{}
	data, err := packed.GetCharts()

	if err = json.Unmarshal(data, &out.List); err != nil {
		return nil, err
	}
	return out, nil
}
