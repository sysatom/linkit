package types

type AgentContent struct {
	Id      string `json:"id"`
	Version int    `json:"version"`
	Content KV     `json:"content"`
}
