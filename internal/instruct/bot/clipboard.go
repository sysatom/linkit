package bot

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/sysatom/linkit/internal/types"
)

var clipboard = []Do{
	{
		Flag: "clipboard_share",
		Run: func(app fyne.App, window fyne.Window, data types.KV) error {
			txt, _ := data.String("txt")
			if txt != "" {
				d := dialog.NewInformation("clipboard", "share text from chat", window)
				d.Show()
				window.Clipboard().SetContent(txt)
			}
			return nil
		},
	},
}
