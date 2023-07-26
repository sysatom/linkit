package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/sysatom/linkit/internal/pkg/constant"
	"github.com/sysatom/linkit/internal/pkg/util"
	"strconv"
)

type settings struct {
	serverHostEntry *widget.Entry
	logPathEntry    *widget.Entry
	tokenEntry      *widget.Entry
	intervalEntry   *widget.Entry

	preferences fyne.Preferences
	window      fyne.Window
}

func newSettingsTab(app fyne.App, window fyne.Window) *container.TabItem {
	s := &settings{window: window, preferences: app.Preferences()}

	return &container.TabItem{
		Text:    "Settings",
		Icon:    theme.SettingsIcon(),
		Content: s.buildUI(app),
	}
}

func (s *settings) getPreferences(_ fyne.App) {
	s.serverHostEntry.Text = s.preferences.String(constant.ServerPreferenceKey)
	s.logPathEntry.Text = s.preferences.String(constant.LogPreferenceKey)
	s.tokenEntry.Text = s.preferences.String(constant.TokenPreferenceKey)
	s.intervalEntry.Text = fmt.Sprintf("%d", s.preferences.Int(constant.IntervalPreferenceKey))
}

func (s *settings) buildUI(app fyne.App) *container.Scroll {
	s.serverHostEntry = &widget.Entry{PlaceHolder: "Enter server host (eg. 127.0.0.1:6060)", OnChanged: s.onServerChanged}
	pathSelector := &widget.Button{Icon: theme.FolderOpenIcon(), Importance: widget.LowImportance, OnTapped: s.onLogPathSelected}
	s.logPathEntry = &widget.Entry{Wrapping: fyne.TextTruncate, OnSubmitted: s.onDownloadsPathSubmitted, ActionItem: pathSelector}

	s.tokenEntry = &widget.Entry{PlaceHolder: "Enter your bot access token.", Password: true, OnChanged: s.onTokenChanged}
	s.intervalEntry = &widget.Entry{PlaceHolder: "60 sec", OnChanged: s.onIntervalChanged}

	s.getPreferences(app)

	systemContainer := container.NewGridWithColumns(2,
		newBoldLabel("Server host"), s.serverHostEntry,
		newBoldLabel("Logs path"), s.logPathEntry,
	)
	botContainer := container.NewGridWithColumns(2,
		newBoldLabel("Access token"), s.tokenEntry,
		newBoldLabel("Request Interval"), s.intervalEntry,
	)

	return container.NewScroll(container.NewVBox(
		&widget.Card{Title: "System", Content: systemContainer},
		&widget.Card{Title: "Bot", Content: botContainer},
	))
}

func (s *settings) onLogPathSelected() {
	folder := dialog.NewFolderOpen(func(folder fyne.ListableURI, err error) {
		if err != nil {
			fyne.LogError("Error on selecting folder", err)
			dialog.ShowError(err, s.window)
			return
		} else if folder == nil {
			return
		}
		s.preferences.SetString(constant.LogPreferenceKey, folder.Path())
		s.logPathEntry.SetText(folder.Path())
	}, s.window)

	folder.Resize(util.WindowSizeToDialog(s.window.Canvas().Size()))
	folder.Show()
}

func (s *settings) onDownloadsPathSubmitted(d string) {
	fmt.Println(d)
}

func (s *settings) onServerChanged(val string) {
	s.preferences.SetString(constant.ServerPreferenceKey, val)
}

func (s *settings) onTokenChanged(val string) {
	s.preferences.SetString(constant.TokenPreferenceKey, val)
}

func (s *settings) onIntervalChanged(val string) {
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		s.intervalEntry.Text = ""
		d := dialog.NewInformation("Input", "Please input number", s.window)
		d.Show()
		return
	}
	s.preferences.SetInt(constant.IntervalPreferenceKey, int(i))
}

func newBoldLabel(text string) *widget.Label {
	return &widget.Label{Text: text, TextStyle: fyne.TextStyle{Bold: true}}
}
