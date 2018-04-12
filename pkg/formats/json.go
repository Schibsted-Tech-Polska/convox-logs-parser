package formats

import (
	"fmt"
	"errors"
)

type JsonMessage struct {
	jsonObject map[string]interface{}
}



func NewJsonMessage(jsonObject map[string]interface{}) JsonMessage {
	jm := JsonMessage{jsonObject}

	return jm
}

func getJSONStringFieldSafe(fn string, jsonObject map[string]interface{}) (string, error) {
	switch v := jsonObject[fn].(type) {
	default:
		return "", errors.New(fmt.Sprintf("Not supported type [%s]", v))
	case string:
		return jsonObject[fn].(string), nil
	}
}
