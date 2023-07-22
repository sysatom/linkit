package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func SetupSystray(desk desktop.App) {
	// Set up menu
	// desk.SetSystemTrayIcon(icon.Data)

	menu := fyne.NewMenu("MyApp",
		//fyne.NewMenuItem("Open", mainWindow.Show),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Tomato", func() {
			//startCountdown(DEFAULT_TIMERS["TOMATO"])
		}),
		fyne.NewMenuItem("Short break", func() {
			//startCountdown(DEFAULT_TIMERS["SHORT"])
		}),
		fyne.NewMenuItem("Long break", func() {
			//startCountdown(DEFAULT_TIMERS["LONG"])
		}),
	)
	desk.SetSystemTrayMenu(menu)
}
