package bot

import (
	"fyne.io/fyne/v2"
	"github.com/sysatom/linkit/internal/pkg/logs"
	"github.com/sysatom/linkit/internal/pkg/types"
	"time"
)

var dev = []Executor{
	{
		Flag: "dev_example",
		Run: func(app fyne.App, window fyne.Window, data types.KV) error {
			logs.Info("dev example %s %s", data, time.Now())
			return nil
		},
	},
}
