package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"net/url"
)

func newAboutTab(app fyne.App) *container.TabItem {
	const (
		https   = "https"
		github  = "github.com"
		version = "v0.0.1"
	)

	repoURL := &url.URL{Scheme: https, Host: github, Path: "/sysatom/linkit"}
	icon := newClickableIcon(app.Icon(), repoURL, app)

	nameLabel := newBoldLabel("Linkit")
	spacerLabel := newBoldLabel("-")

	releaseURL := &url.URL{
		Scheme: https, Host: github,
		Path: "/sysatom/linkit/releases/tag/" + version,
	}
	hyperlink := &widget.Hyperlink{Text: version, URL: releaseURL, TextStyle: fyne.TextStyle{Bold: true}}

	spacer := &layout.Spacer{}

	content := container.NewVBox(
		spacer,
		container.NewHBox(spacer, icon, spacer),
		container.NewHBox(
			spacer,
			nameLabel,
			spacerLabel,
			hyperlink,
			spacer,
		),
		spacer,
	)

	return &container.TabItem{
		Text:    "About",
		Icon:    theme.InfoIcon(),
		Content: content,
	}
}

type clickableIcon struct {
	widget.BaseWidget
	app  fyne.App
	url  *url.URL
	icon *canvas.Image
}

func (c *clickableIcon) Tapped(_ *fyne.PointEvent) {
	err := c.app.OpenURL(c.url)
	if err != nil {
		fyne.LogError("Failed to open repository: ", err)
	}
}

func (c *clickableIcon) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}

func (c *clickableIcon) CreateRenderer() fyne.WidgetRenderer {
	c.ExtendBaseWidget(c)
	return widget.NewSimpleRenderer(c.icon)
}

func (c *clickableIcon) MinSize() fyne.Size {
	return fyne.Size{Width: 256, Height: 256}
}

func newClickableIcon(res fyne.Resource, url *url.URL, app fyne.App) *clickableIcon {
	return &clickableIcon{app: app, url: url, icon: canvas.NewImageFromResource(res)}
}
