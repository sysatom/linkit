package instruct

import (
	"context"
	"fyne.io/fyne/v2"
	"github.com/allegro/bigcache/v3"
	"github.com/robfig/cron/v3"
	"github.com/sysatom/linkit/internal/pkg/client"
	"github.com/sysatom/linkit/internal/pkg/constant"
	"github.com/sysatom/linkit/internal/pkg/logs"
	"time"
)

func Cron(app fyne.App, window fyne.Window) {
	c := cron.New(cron.WithSeconds())
	// instruct job
	accessToken := app.Preferences().String(constant.TokenPreferenceKey)
	if accessToken != "" {
		cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(time.Hour))
		if err != nil {
			logs.Panic(err.Error())
		}
		job := &instructJob{app: app, window: window, client: client.NewTinode(accessToken), cache: cache}
		_, err = c.AddJob("*/10 * * * * *", job)
		if err != nil {
			logs.Panic(err.Error())
		}
	}
	c.Start()
}
