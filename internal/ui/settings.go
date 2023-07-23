package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/sysatom/linkit/internal/constant"
	"github.com/sysatom/linkit/internal/util"
	"strconv"
)

type settings struct {
	downloadPathEntry *widget.Entry
	tokenEntry        *widget.Entry
	intervalEntry     *widget.Entry

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
	s.tokenEntry.Text = s.preferences.String(constant.TokenPreferenceKey)
	s.intervalEntry.Text = fmt.Sprintf("%d", s.preferences.Int(constant.IntervalPreferenceKey))
}

func (s *settings) buildUI(app fyne.App) *container.Scroll {
	pathSelector := &widget.Button{Icon: theme.FolderOpenIcon(), Importance: widget.LowImportance, OnTapped: s.onDownloadsPathSelected}
	s.downloadPathEntry = &widget.Entry{Wrapping: fyne.TextTruncate, OnSubmitted: s.onDownloadsPathSubmitted, ActionItem: pathSelector}

	s.tokenEntry = &widget.Entry{PlaceHolder: "Enter your bot access token.", Password: true, OnChanged: s.onTokenChanged}
	s.intervalEntry = &widget.Entry{PlaceHolder: "60 sec", OnChanged: s.onIntervalChanged}

	s.getPreferences(app)

	botContainer := container.NewGridWithColumns(2,
		newBoldLabel("Token"), s.tokenEntry,
		newBoldLabel("Interval"), s.intervalEntry,
	)

	return container.NewScroll(container.NewVBox(
		&widget.Card{Title: "Bot settings", Content: botContainer},
	))
}

func (s *settings) onDownloadsPathSelected() {
	folder := dialog.NewFolderOpen(func(folder fyne.ListableURI, err error) {
		if err != nil {
			fyne.LogError("Error on selecting folder", err)
			dialog.ShowError(err, s.window)
			return
		} else if folder == nil {
			return
		}

		fmt.Println(folder.Path())
		s.preferences.SetString("DownloadPath", folder.Path())
		s.downloadPathEntry.SetText(folder.Path())
	}, s.window)

	folder.Resize(util.WindowSizeToDialog(s.window.Canvas().Size()))
	folder.Show()
}

func (s *settings) onDownloadsPathSubmitted(d string) {
	fmt.Println(d)
}

func (s *settings) onTokenChanged(val string) {
	old := s.preferences.String(constant.TokenPreferenceKey)
	if old != "" && old != val {
		d := dialog.NewInformation("Info", "Setting a new token requires a restart of the application to take effect.", s.window)
		d.Show()
	}
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
