package util

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/sysatom/linkit/internal/pkg/constant"
	"github.com/sysatom/linkit/internal/pkg/logs"
	"time"
)

func CheckSingleton() {
	if !PortAvailable(constant.EmbedServerPort) {
		resp, err := resty.New().SetTimeout(500 * time.Millisecond).R().
			Get(fmt.Sprintf("http://127.0.0.1:%s/", constant.EmbedServerPort))
		if err != nil {
			logs.Err.Println(err)
			return
		}
		if resp.String() == "ok" {
			panic("app exists")
		}
	}
}
