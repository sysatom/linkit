package agent

import (
	"context"
	"fyne.io/fyne/v2"
	"github.com/allegro/bigcache/v3"
	"github.com/robfig/cron/v3"
	"github.com/sysatom/linkit/internal/pkg/client"
	"github.com/sysatom/linkit/internal/pkg/logs"
	"github.com/sysatom/linkit/internal/pkg/setting"
	"time"
)

func Cron(app fyne.App, window fyne.Window) {
	c := cron.New(cron.WithSeconds())
	// agent job
	if setting.Get().AccessToken != "" {
		cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(24*time.Hour))
		if err != nil {
			logs.Panic(err.Error())
		}
		job := &agentJob{app: app, window: window, cache: cache, client: client.NewTinode(setting.Get().AccessToken)}
		job.RunClipboard(c)
		job.RunAnki(c)
		job.RunDev(c)
	}
	c.Start()
}
