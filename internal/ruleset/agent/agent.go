package agent

import (
	"context"
	"fyne.io/fyne/v2"
	"github.com/allegro/bigcache/v3"
	"github.com/robfig/cron/v3"
	"github.com/sysatom/linkit/internal/pkg/client"
	"github.com/sysatom/linkit/internal/pkg/constant"
	"time"
)

func Cron(app fyne.App, window fyne.Window) {
	c := cron.New(cron.WithSeconds())
	// agent job
	accessToken := app.Preferences().String(constant.TokenPreferenceKey)
	if accessToken != "" {
		cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(24*time.Hour))
		if err != nil {
			panic(err)
		}
		job := &agentJob{app: app, window: window, cache: cache, client: client.NewTinode(accessToken)}
		job.RunClipboard(c)
		job.RunAnki(c)
		job.RunDev(c)
	}
	c.Start()
}
