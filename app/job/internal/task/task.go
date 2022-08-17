package task

import (
	"context"
	"github.com/go-co-op/gocron"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"time"
)

func Init() {
	s := gocron.NewScheduler(time.Local)
	// Task 1
	s.Every(1).
		Seconds().
		WaitForSchedule().
		SingletonMode().
		Do(HelloWorld, gctx.New())

	// Task 2
	s.Every(6).
		Seconds().
		WaitForSchedule().
		SingletonMode().
		Do(refreshToken, gctx.New())

	s.StartAsync()
}

func HelloWorld(ctx context.Context) {
	g.Log().Info(gctx.New(), "hello world")
}

func refreshToken(ctx context.Context) {
	g.Log().Info(gctx.New(), "refresh Token ...")
}
