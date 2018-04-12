package formats

import "fmt"

type StandardJavaMessage struct {
	JsonMessage
	IsValidJavaMessage bool

	thread     string
	level      string
	loggerName string
	message    string
}

func NewStandardJavaMessage(jsonObject map[string]interface{}) StandardJavaMessage {
	var sj StandardJavaMessage
	sj.JsonMessage = NewJsonMessage(jsonObject)
	sj.parseJsonObject()
	return sj
}

func (sj *StandardJavaMessage) parseJsonObject() {
	var err error
	sj.SetToValid()

	if sj.thread, err = getJSONStringFieldSafe("thread", sj.jsonObject); err != nil {
		sj.SetToInvalid()
	}

	if sj.loggerName, err = getJSONStringFieldSafe("loggerName", sj.jsonObject); err != nil {
		sj.SetToInvalid()
	}

	if sj.level, err = getJSONStringFieldSafe("level", sj.jsonObject); err != nil {
		sj.SetToInvalid()
	}

	if sj.message, err = getJSONStringFieldSafe("message", sj.jsonObject); err != nil {
		sj.SetToInvalid()
	}
}

func (sj StandardJavaMessage) String() string {
	out := fmt.Sprintf("[%s] %s %s - %s", sj.thread, sj.level, sj.loggerName, sj.message)

	return out
}

func (sj *StandardJavaMessage) SetToValid() {
	sj.IsValidJavaMessage = true
}

func (sj *StandardJavaMessage) SetToInvalid() {
	sj.IsValidJavaMessage = false
}
