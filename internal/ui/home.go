package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
)

func newHomeTab(app fyne.App, window fyne.Window) *container.TabItem {
	return &container.TabItem{
		Text: "Home",
		Icon: theme.HomeIcon(),
		Content: container.New(
			layout.NewVBoxLayout(),
			layout.NewSpacer(),
		),
	}
}
