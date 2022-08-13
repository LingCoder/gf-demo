package v1

import (
	"gf-demo/app/console/internal/model"
	"gf-demo/library/base"
	"github.com/gogf/gf/v2/frame/g"
)

type ChartsKlineReq struct {
	g.Meta `tags:"Charts" summary:"View Charts Kline" dc:"Charts Kline list"`
}

type ChartsKlineRes struct {
	base.Res
	List []*model.ChartsKlineItem `json:"list"`
}
