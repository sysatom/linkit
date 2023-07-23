package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/sysatom/linkit/internal"
	"github.com/sysatom/linkit/internal/assets"
	"github.com/sysatom/linkit/internal/instruct"
	"github.com/sysatom/linkit/internal/ui"
)

func main() {
	// app
	a := app.NewWithID("com.github.sysatom.linkit")
	assets.SetIcon(a)
	w := a.NewWindow("Linkit")

	// cron
	instruct.Cron(a, w)

	// theme
	t := internal.NewAppTheme()
	a.Settings().SetTheme(t)

	// systray
	if desk, ok := a.(desktop.App); ok {
		ui.SetupSystray(desk)
	}

	// main window
	w.SetContent(ui.Create(a, w))
	w.Resize(fyne.NewSize(1000, 600))
	w.SetMaster()
	w.ShowAndRun()
}
