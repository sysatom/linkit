package client

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

type Tinode struct {
	c           *resty.Client
	accessToken string
}

func NewTinode(accessToken string) *Tinode {
	v := &Tinode{accessToken: accessToken}

	v.c = resty.New()
	v.c.SetBaseURL("http://127.0.0.1:6060")
	v.c.SetTimeout(time.Minute)

	return v
}

type Operate string

const (
	Agent Operate = "agent"
	Bots  Operate = "bots"
	Pull  Operate = "pull"
	Help  Operate = "help"
)

func (v *Tinode) fetcher(action Operate, content interface{}) ([]byte, error) {
	resp, err := v.c.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", v.accessToken)).
		SetBody(map[string]interface{}{
			"action":  action,
			"version": 1,
			"content": content,
		}).
		Post("/extra/linkit")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.Body(), nil
	} else {
		return nil, fmt.Errorf("%d, %s (%s)",
			resp.StatusCode(),
			resp.Header().Get("X-Error-Code"),
			resp.Header().Get("X-Error"))
	}
}

func (v *Tinode) Bots() (*BotsResult, error) {
	data, err := v.fetcher(Bots, nil)
	if err != nil {
		return nil, err
	}
	var r BotsResult
	err = json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}
	return &r, err
}

type BotsResult struct {
	Bots []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"bots"`
}

func (v *Tinode) Pull() (*InstructResult, error) {
	data, err := v.fetcher(Pull, nil)
	if err != nil {
		return nil, err
	}
	var r InstructResult
	err = json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}
	return &r, err
}

type InstructResult struct {
	Instruct []struct {
		No       string      `json:"no"`
		Bot      string      `json:"bot"`
		Flag     string      `json:"flag"`
		Content  interface{} `json:"content"`
		ExpireAt string      `json:"expire_at"`
	} `json:"instruct"`
}
