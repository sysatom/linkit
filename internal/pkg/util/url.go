package util

import (
	"fmt"
	"strings"
)

func FillScheme(host string) string {
	scheme := "https"
	if strings.HasPrefix(host, "127.") || strings.HasPrefix(host, "192.") {
		scheme = "http"
	}
	return fmt.Sprintf("%s://%s", scheme, host)
}
