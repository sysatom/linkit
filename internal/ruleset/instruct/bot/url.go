package bot

import (
	"fyne.io/fyne/v2"
	"github.com/sysatom/linkit/internal/pkg/types"
	netUrl "net/url"
)

var url = []Executor{
	{
		Flag: "url_open",
		Run: func(app fyne.App, window fyne.Window, data types.KV) error {
			txt, _ := data.String("url")
			if txt != "" {
				u, err := netUrl.Parse(txt)
				if err != nil {
					return err
				}
				err = app.OpenURL(u)
				if err != nil {
					return err
				}
				app.SendNotification(fyne.NewNotification("url", "open url"))
			}
			return nil
		},
	},
}
