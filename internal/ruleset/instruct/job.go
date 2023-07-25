package instruct

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"github.com/allegro/bigcache/v3"
	"github.com/sysatom/linkit/internal/pkg/client"
	"github.com/sysatom/linkit/internal/pkg/constant"
	"github.com/sysatom/linkit/internal/pkg/types"
	"github.com/sysatom/linkit/internal/ruleset/instruct/bot"
	"log"
	"time"
)

type instructJob struct {
	app    fyne.App
	window fyne.Window
	cache  *bigcache.BigCache
	client *client.Tinode
}

func (j *instructJob) Run() {
	res, err := j.client.Pull()
	if err != nil {
		log.Println(err)
		return
	}
	if res == nil {
		return
	}
	// get preference
	d := j.app.Preferences().String(constant.InstructPreferenceKey)
	data := types.KV{}
	_ = json.Unmarshal([]byte(d), &data)
	// instruct loop
	for _, item := range res.Instruct {
		// check switch
		s, ok := data.String(item.Bot)
		if !ok || s == "" || s == "Off" {
			continue
		}
		// check has been run
		has, _ := j.cache.Get(item.No)
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
				err = do.Run(j.app, j.window, data)
				if err != nil {
					log.Println("instruct run job failed", item.Bot, item.No)
				}
				err = j.cache.Set(item.No, []byte("1"))
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}
