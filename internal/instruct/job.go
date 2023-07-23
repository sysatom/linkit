package instruct

import (
	"fyne.io/fyne/v2"
	"github.com/allegro/bigcache/v3"
	"github.com/sysatom/linkit/internal/client"
	"github.com/sysatom/linkit/internal/instruct/bot"
	"github.com/sysatom/linkit/internal/types"
	"log"
	"time"
)

type instructJob struct {
	app    fyne.App
	window fyne.Window
	cache  *bigcache.BigCache
	client *client.Tinode
}

func (i *instructJob) Run() {
	res, err := i.client.Pull()
	if err != nil {
		log.Println(err)
		return
	}
	if res == nil {
		return
	}
	for _, item := range res.Instruct {
		// check has been run
		has, _ := i.cache.Get(item.No)
		if len(has) > 0 {
			continue
		}
		// check expired
		expiredAt, err := time.Parse("2006-01-02T15:04:05Z", item.ExpireAt)
		if err != nil {
			continue
		}
		if time.Now().After(expiredAt) {
			continue
		}
		for id, dos := range bot.DoInstruct {
			if item.Bot != id {
				continue
			}
			for _, do := range dos {
				if item.Flag != do.Flag {
					continue
				}
				// run instruct
				log.Println("instruct run job", item.Bot, item.No)
				data := types.KV{}
				if v, ok := item.Content.(map[string]interface{}); ok {
					data = v
				}
				err = do.Run(i.app, i.window, data)
				if err != nil {
					log.Println("instruct run job failed", item.Bot, item.No)
				}
				err = i.cache.Set(item.No, []byte("1"))
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}
