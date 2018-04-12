package formats

import (
	"encoding/json"
	"errors"
)

// Message is an interface describing possible behaviour of received messages
type Message interface {
	String() string
}

// New is a standard Message factory. The factory checks what kind of message is given
// and it's using proper format class to parse it.
func New(message string) Message {
	if jsonObject, err := ParseJSONMessage(message); err == nil {
		if jm := NewStandardJavaMessage(jsonObject); jm.IsValidJavaMessage {
			return jm
		}
	}
	return NewPlainMessage(message)
}

// ParseJSONMessage is an overlay for json.Unmarshal which additionally creates an object
// to be set instead of accepting a reference.
func ParseJSONMessage(message string) (map[string]interface{}, error) {
	var js map[string]interface{}

	if err := json.Unmarshal([]byte(message), &js); err != nil {
		return nil, errors.New("This is not a valid JSON message")
	}

	return js, nil
}
