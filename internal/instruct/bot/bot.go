package bot

import (
	"fyne.io/fyne/v2"
	"github.com/sysatom/linkit/internal/types"
)

type Do struct {
	Flag string
	Run  func(app fyne.App, window fyne.Window, data types.KV) error
}

var DoInstruct = map[string][]Do{
	"dev":       dev,
	"clipboard": clipboard,
}
