package event

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestSubscribe(t *testing.T) {
	Publish(Login, gctx.New())
	Subscribe(Login, loginAction)
}

func loginAction(ctx context.Context) {
	g.Log().Info(ctx, "login success")
}
