package bot

import (
	"fyne.io/fyne/v2"
	"github.com/sysatom/linkit/internal/pkg/types"
)

var clipboard = []Executor{
	{
		Flag: "clipboard_share",
		Run: func(app fyne.App, window fyne.Window, data types.KV) error {
			txt, _ := data.String("txt")
			if txt != "" {
				app.SendNotification(fyne.NewNotification("clipboard", "share text from chat"))
				window.Clipboard().SetContent(txt)
			}
			return nil
		},
	},
}
