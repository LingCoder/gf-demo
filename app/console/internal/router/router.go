package router

import (
	"gf-demo/app/console/internal/consts"
	"gf-demo/app/console/internal/controller"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gsession"
)

func Setup() *ghttp.Server {
	s := g.Server()
	s.SetSessionStorage(gsession.NewStorageMemory())
	s.Use(ghttp.MiddlewareHandlerResponse)
	base := s.Group(consts.AppName)
	base.GET("/user/detail", controller.User.UserDetail)

	return s
}
