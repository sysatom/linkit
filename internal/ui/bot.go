package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/sysatom/linkit/internal/client"
)

type bot struct {
	client      *client.Tinode
	accessToken string

	app         fyne.App
	window      fyne.Window
	preferences fyne.Preferences
}

func newBotTab(app fyne.App, window fyne.Window) *container.TabItem {
	b := bot{app: app, window: window, preferences: app.Preferences()}
	return &container.TabItem{
		Text:    "Bot",
		Icon:    theme.ComputerIcon(),
		Content: b.buildUI(app),
	}
}

func (b *bot) buildUI(app fyne.App) *container.Scroll {
	onOffOptions := []string{"On", "Off"}

	b.getPreferences(app)
	b.client = client.NewTinode(b.accessToken)

	var options []string
	var co []fyne.CanvasObject
	co = append(co, &widget.Label{Text: "manage your bots instruct settings here."})

	res, err := b.client.Bots()
	if err != nil {
		fmt.Println(err)
	}
	if res != nil {
		for _, item := range res.Bots {
			options = append(options, item.Name)
			co = append(co, container.NewGridWithColumns(2,
				newBoldLabel(item.Name), &widget.RadioGroup{Options: onOffOptions, Horizontal: true, Required: true, OnChanged: b.onSwitchChanged},
			))
		}
	}

	settingsContainer := &widget.Card{Title: "Bots Instruct Settings", Content: container.NewVBox(co...)}

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
		widget.NewSelectEntry(options),
		&widget.Card{Title: "dev", Content: table},
	)

	panelContainer := container.NewGridWithColumns(2,
		settingsContainer, infoContainer,
	)

	return container.NewScroll(panelContainer)
}

func (b *bot) getPreferences(app fyne.App) {
	b.accessToken = b.preferences.String(tokenSettingKey)
}

func (b *bot) onSwitchChanged(val string) {
	fmt.Println(val)
}

func onOrOff(on bool) string {
	if on {
		return "On"
	}

	return "Off"
}
