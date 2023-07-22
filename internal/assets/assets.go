package assets

import (
	_ "embed"
	"fyne.io/fyne/v2"
)

//go:embed icon/app-icon.png
var iconData []byte

func SetIcon(a fyne.App) {
	a.SetIcon(&fyne.StaticResource{
		StaticName:    "app-icon.png",
		StaticContent: iconData,
	})
}
