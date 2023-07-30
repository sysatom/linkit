package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type KV map[string]any

func (j *KV) Scan(value any) error {
	if bytes, ok := value.([]byte); ok {
		result := make(map[string]any)
		err := json.Unmarshal(bytes, &result)
		if err != nil {
			return err
		}
		*j = result
		return nil
	}
	if result, ok := value.(map[string]any); ok {
		*j = result
		return nil
	}
	return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
}

func (j KV) String(key string) (string, bool) {
	if v, ok := j.get(key); ok {
		if t, ok := v.(string); ok {
			return t, ok
		}
	}
	return "", false
}

func (j KV) Int64(key string) (int64, bool) {
	if v, ok := j.get(key); ok {
		switch n := v.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			return reflect.ValueOf(n).Convert(reflect.TypeOf(int64(0))).Int(), true
		case float32, float64:
			return reflect.ValueOf(n).Convert(reflect.TypeOf(int64(0))).Int(), true
		}
	}
	return 0, false
}

func (j KV) Uint64(key string) (uint64, bool) {
	if v, ok := j.get(key); ok {
		if t, ok := v.(float64); ok {
			return uint64(t), ok
		}
	}
	return 0, false
}

func (j KV) Float64(key string) (float64, bool) {
	if v, ok := j.get(key); ok {
		if t, ok := v.(float64); ok {
			return t, ok
		}
	}
	return 0, false
}

func (j KV) Map(key string) (map[string]any, bool) {
	if v, ok := j.get(key); ok {
		if t, ok := v.(map[string]any); ok {
			return t, ok
		}
	}
	return nil, false
}

func (j KV) get(key string) (any, bool) {
	v, ok := j[key]
	return v, ok
}

func (j KV) StringValue() (string, bool) {
	return j.String("value")
}

func (j KV) Int64Value() (int64, bool) {
	return j.Int64("value")
}

func (j KV) Uint64Value() (uint64, bool) {
	return j.Uint64("value")
}

func (j KV) Float64Value() (float64, bool) {
	return j.Float64("value")
}

type Action string

type Data struct {
	Action  Action `json:"action"`
	Version int    `json:"version"`
	Content KV     `json:"content"`
}

// ClientComMessage is a wrapper for client messages.
type ClientComMessage struct {
	Data Data `json:"data"`
}

// ServerComMessage is a wrapper for server-side messages.
type ServerComMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
