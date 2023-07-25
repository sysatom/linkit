package theme

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/sysatom/linkit/internal/assets"
	"image/color"
)

type AppTheme struct {
	regular, bold, italic, boldItalic, monospace fyne.Resource
}

func NewAppTheme() *AppTheme {
	t := &AppTheme{}
	t.setDefaultResource()
	return t
}

func (t *AppTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (t *AppTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *AppTheme) Font(style fyne.TextStyle) fyne.Resource {
	if style.Monospace {
		return t.monospace
	}
	if style.Bold {
		if style.Italic {
			return t.boldItalic
		}
		return t.bold
	}
	if style.Italic {
		return t.italic
	}
	return t.regular
}

func (t *AppTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func (t *AppTheme) setDefaultResource() {
	t.regular = theme.TextFont()
	t.bold = theme.TextBoldFont()
	t.italic = theme.TextItalicFont()
	t.boldItalic = theme.TextBoldItalicFont()
	t.monospace = theme.TextMonospaceFont()

	abc := &fyne.StaticResource{
		StaticName:    "SourceHanSans-Medium.ttf",
		StaticContent: assets.FontData,
	}

	t.regular = abc
	t.bold = abc
	t.italic = abc
	t.boldItalic = abc
	t.monospace = t.regular
}
