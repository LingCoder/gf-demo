package cmd

import (
	"context"
	"gf-demo/app/job/internal/task"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start crontab job",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			task.Init()
			// Block listening the shutdown signal.
			g.Listen()
			return
		},
	}
)
