package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/sysatom/linkit/internal"
	"github.com/sysatom/linkit/internal/assets"
	"github.com/sysatom/linkit/internal/ui"
)

func main() {
	// app
	a := app.NewWithID("com.github.sysatom.linkit")
	assets.SetIcon(a)

	// theme
	t := internal.NewAppTheme()
	a.Settings().SetTheme(t)

	// systray
	if desk, ok := a.(desktop.App); ok {
		ui.SetupSystray(desk)
	}

	// main window
	w := a.NewWindow("Linkit")
	w.SetContent(ui.Create(a, w))
	w.Resize(fyne.NewSize(1280, 720))
	w.SetMaster()
	w.ShowAndRun()
}
