package controller

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type cBase struct {
}

// Init auto run before request
func (c *cBase) Init(r *ghttp.Request) {
	g.Log().Debug(r.Context(), "request...")
}
