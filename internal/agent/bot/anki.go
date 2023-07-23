package bot

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/sysatom/linkit/internal/client"
	"github.com/sysatom/linkit/internal/types"
	"log"
	"net/http"
	"strconv"
)

const (
	AnkiAgentVersion = 1

	StatsAgentID  = "stats_agent"
	ReviewAgentID = "review_agent"
)

func AnkiStats(c *client.Tinode) {
	html, err := getCollectionStatsHTML()
	if err != nil {
		log.Println(err)
		return
	}
	_, err = c.Agent(types.AgentContent{
		Id:      StatsAgentID,
		Version: AnkiAgentVersion,
		Content: map[string]interface{}{
			"html": html,
		},
	})
	if err != nil {
		log.Println(err)
	}
}

func AnkiReview(c *client.Tinode) {
	num, err := getNumCardsReviewedToday()
	if err != nil {
		log.Println(err)
		return
	}
	_, err = c.Agent(types.AgentContent{
		Id:      ReviewAgentID,
		Version: AnkiAgentVersion,
		Content: map[string]interface{}{
			"num": num,
		},
	})
	if err != nil {
		log.Println(err)
	}
}

func getCollectionStatsHTML() (string, error) {
	c := resty.New()
	resp, err := c.R().
		SetContext(context.Background()).
		SetBody(Param{
			Action:  "getCollectionStatsHTML",
			Version: ApiVersion,
			Params: map[string]interface{}{
				"wholeCollection": true,
			},
		}).
		SetResult(&Response{}).
		Post(ApiURI)
	if err != nil {
		return "", err
	}

	if resp.StatusCode() == http.StatusOK {
		respResult := resp.Result().(*Response)
		if respResult != nil {
			if respResult.Error != nil {
				return "", errors.New(*respResult.Error)
			}

			return string(respResult.Result), nil
		}
	}
	return "", errors.New("result error")
}

func getNumCardsReviewedToday() (int, error) {
	c := resty.New()
	resp, err := c.R().
		SetContext(context.Background()).
		SetBody(Param{
			Action:  "getNumCardsReviewedToday",
			Version: ApiVersion,
		}).
		SetResult(&Response{}).
		Post(ApiURI)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode() == http.StatusOK {
		respResult := resp.Result().(*Response)
		if respResult != nil {
			if respResult.Error != nil {
				return 0, errors.New(*respResult.Error)
			}

			n, _ := strconv.Atoi(string(respResult.Result))
			return n, nil
		}
	}
	return 0, errors.New("result error")
}

const ApiVersion = 6
const ApiURI = "http://localhost:8765"

type Param struct {
	Action  string      `json:"action"`
	Version int         `json:"version"`
	Params  interface{} `json:"params,omitempty"`
}

type Response struct {
	Error  *string         `json:"error"`
	Result json.RawMessage `json:"result"`
}
