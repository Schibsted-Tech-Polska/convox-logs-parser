package formats

import (
	"encoding/json"
	"errors"
)

type Message interface {
	String() string
}

func New(message string) Message {
	if jsonObject, err := ParseJsonMessage(message); err == nil {
		if jm := NewStandardJavaMessage(jsonObject); jm.IsValidJavaMessage {
			return jm
		}
	}
	return NewPlainMessage(message)
}

func ParseJsonMessage(message string) (map[string]interface{}, error) {
	var js map[string]interface{}

	if err := json.Unmarshal([]byte(message), &js); err != nil {
		return nil, errors.New("This is not a valid JSON message")
	}

	return js, nil
}
