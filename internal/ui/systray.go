package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func SetupSystray(desk desktop.App, window fyne.Window) {
	menu := fyne.NewMenu("MyApp",
		fyne.NewMenuItem("Open", window.Show),
	)
	desk.SetSystemTrayMenu(menu)
}
