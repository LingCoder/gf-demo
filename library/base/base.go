package base

import "github.com/gogf/gf/v2/frame/g"

type Res struct {
	g.Meta `mime:"application/json"`
}

type PageReq struct {
	Page
}

type PageRes struct {
	Res
	TotalInt
}

type TotalInt struct {
	Total int `json:"total"`
}

type Keywords struct {
	Keywords int `json:"keywords"`
}

type PageInput struct {
	Page
}
