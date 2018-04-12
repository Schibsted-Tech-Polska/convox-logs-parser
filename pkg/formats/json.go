package formats

import (
	"fmt"
)

// JSONMessage is basic Message containing JSON structure
type JSONMessage struct {
	jsonObject map[string]interface{}
}

// NewJSONMessage is a JSONMessage factory.
func NewJSONMessage(jsonObject map[string]interface{}) JSONMessage {
	jm := JSONMessage{jsonObject}

	return jm
}

func getJSONStringFieldSafe(fn string, jsonObject map[string]interface{}) (string, error) {
	switch v := jsonObject[fn].(type) {
	default:
		return "", fmt.Errorf("Not supported type [%s]", v)
	case string:
		return jsonObject[fn].(string), nil
	}
}
