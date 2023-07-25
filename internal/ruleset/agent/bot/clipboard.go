package bot

import (
	"fmt"
	"fyne.io/fyne/v2"
	"github.com/allegro/bigcache/v3"
	"github.com/sysatom/linkit/internal/pkg/client"
	"github.com/sysatom/linkit/internal/pkg/types"
	"log"
)

const (
	ClipboardAgentVersion = 1
	UploadAgentID         = "clipboard_upload"
)

func ClipboardUpload(window fyne.Window, cache *bigcache.BigCache, c *client.Tinode) {
	old, _ := cache.Get("clipboard")
	now := window.Clipboard().Content()
	fmt.Printf("clipboard upload (%s) (%s)\n", string(old), now)
	if string(old) == now {
		return
	}
	_, err := c.Agent(types.AgentContent{
		Id:      UploadAgentID,
		Version: ClipboardAgentVersion,
		Content: types.KV{
			"txt": now,
		},
	})
	if err != nil {
		log.Println(err)
	}
	_ = cache.Set("clipboard", []byte(now))
}
