package cmd

import (
	"context"
	"gf-demo/app/console/internal/router"
	"github.com/gogf/gf/v2/os/gcmd"
)

var Main = gcmd.Command{
	Name:  "main",
	Usage: "main",
	Brief: "Start Http Server",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		s := router.Setup()
		s.Run()
		return nil
	},
}
