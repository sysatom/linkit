package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type page struct {
	app    fyne.App
	window fyne.Window
}

func newBotTab(app fyne.App, window fyne.Window) *container.TabItem {
	p := page{app: app, window: window}
	return &container.TabItem{
		Text:    "Bot",
		Icon:    theme.ComputerIcon(),
		Content: p.buildUI(),
	}
}

func (p *page) buildUI() *container.Scroll {
	onOffOptions := []string{"On", "Off"}

	settingsContainer := &widget.Card{Title: "Bots Instruct Settings", Content: container.NewVBox(
		&widget.Label{Text: "manage your bots instruct settings here."},
		container.NewGridWithColumns(2,
			newBoldLabel("dev"), &widget.RadioGroup{Options: onOffOptions, Horizontal: true, Required: true, OnChanged: p.onSwitchChanged},
		),
		container.NewGridWithColumns(2,
			newBoldLabel("share"), &widget.RadioGroup{Options: onOffOptions, Horizontal: true, Required: true, OnChanged: p.onSwitchChanged},
		),
	)}

	var data = [][]string{[]string{"top left", "top right"},
		[]string{"bottom left", "bottom right"}}

	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		},
	)
	infoContainer := container.NewVBox(
		widget.NewSelectEntry([]string{"dev", "share"}),
		&widget.Card{Title: "dev", Content: table},
	)

	panelContainer := container.NewGridWithColumns(2,
		settingsContainer, infoContainer,
	)

	return container.NewScroll(panelContainer)
}

func (p *page) onSwitchChanged(val string) {
	fmt.Println(val)
}

func onOrOff(on bool) string {
	if on {
		return "On"
	}

	return "Off"
}
