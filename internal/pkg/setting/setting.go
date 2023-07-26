package setting

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"github.com/sysatom/linkit/internal/pkg/constant"
	"github.com/sysatom/linkit/internal/pkg/types"
)

var s = &Setting{}

type Setting struct {
	ServerHost      string
	LogPath         string
	AccessToken     string
	RequestInterval int
	InstructSwitch  types.KV
}

func LoadPreferences(p fyne.Preferences) {
	s.ServerHost = p.String(constant.ServerPreferenceKey)
	s.LogPath = p.String(constant.LogPreferenceKey)
	s.AccessToken = p.String(constant.TokenPreferenceKey)
	s.RequestInterval = p.Int(constant.IntervalPreferenceKey)

	data := p.String(constant.InstructPreferenceKey)
	instructSwitch := types.KV{}
	_ = json.Unmarshal([]byte(data), &instructSwitch)
	s.InstructSwitch = instructSwitch
}

func Get() *Setting {
	return s
}
