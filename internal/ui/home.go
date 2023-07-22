package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"time"
)

func newHomeTab(app fyne.App, window fyne.Window) *container.TabItem {
	clock := widget.NewLabel("")
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	return &container.TabItem{
		Text: "Home",
		Icon: theme.HomeIcon(),
		Content: container.New(
			layout.NewVBoxLayout(),
			layout.NewSpacer(),
			widget.NewLabel("功能一"),
			widget.NewLabel("功能二"),
			clock,
			widget.NewButton("确定", func() {
				//dialog.ShowInformation("标题", "文本说明。。。。", w)

				// win 2
				w2 := app.NewWindow("Larger")
				w2.SetContent(widget.NewLabel("More content"))
				w2.Resize(fyne.NewSize(100, 100))
				w2.Show()
			}),
			layout.NewSpacer(),
		),
	}
}

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("时间: 03:04:05")
	clock.SetText(formatted)
}
