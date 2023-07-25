package agent

import (
	"fyne.io/fyne/v2"
	"github.com/allegro/bigcache/v3"
	"github.com/robfig/cron/v3"
	"github.com/sysatom/linkit/internal/pkg/client"
	"github.com/sysatom/linkit/internal/pkg/util"
	"github.com/sysatom/linkit/internal/ruleset/agent/bot"
)

type agentJob struct {
	app    fyne.App
	window fyne.Window
	cache  *bigcache.BigCache
	client *client.Tinode
}

func (j *agentJob) RunAnki(c *cron.Cron) {
	util.MustAddFunc(c, "0 * * * * *", func() {
		bot.AnkiStats(j.client)
	})
	util.MustAddFunc(c, "0 * * * * *", func() {
		bot.AnkiReview(j.client)
	})
}

func (j *agentJob) RunClipboard(c *cron.Cron) {
	util.MustAddFunc(c, "*/10 * * * * *", func() {
		bot.ClipboardUpload(j.window, j.cache, j.client)
	})
}

func (j *agentJob) RunDev(c *cron.Cron) {
	util.MustAddFunc(c, "0 * * * * *", func() {
		bot.DevImport(j.client)
	})
}
