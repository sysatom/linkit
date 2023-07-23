package bot

import (
	"fmt"
	"fyne.io/fyne/v2"
	"github.com/sysatom/linkit/internal/types"
	"time"
)

var dev = []Do{
	{
		Flag: "dev_example",
		Run: func(app fyne.App, window fyne.Window, data types.KV) error {
			fmt.Println("dev example", data, time.Now())
			return nil
		},
	},
}
