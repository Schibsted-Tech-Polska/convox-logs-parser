package formats

import "fmt"

// StandardJavaMessage parses log4j2 standard JSON format and displays it
// as it would be displayed normally in stdout without JSON layout.
type StandardJavaMessage struct {
	JSONMessage
	IsValidJavaMessage bool

	thread     string
	level      string
	loggerName string
	message    string
}

// NewStandardJavaMessage is a StandardJavaMessage factory.
// It creates a new StandardJavaMessage object with provided unmarshalled JSON object
//
// Params: jsonObject map[string]interface{}
func NewStandardJavaMessage(jsonObject map[string]interface{}) StandardJavaMessage {
	var sj StandardJavaMessage
	sj.JSONMessage = NewJSONMessage(jsonObject)
	sj.parseJSONObject()
	return sj
}

func (sj *StandardJavaMessage) parseJSONObject() {
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

// SetToValid sets the property of StandardJavaMessage indicating that the provided jsonObject
// contains valid Log4j2 standard message (all required fields are in place)
func (sj *StandardJavaMessage) SetToValid() {
	sj.IsValidJavaMessage = true
}

// SetToInvalid sets the property of StandardJavaMessage indicating that the provided jsonObject
// does not contain valid Log4j2 standard message (some required fields are missing)
func (sj *StandardJavaMessage) SetToInvalid() {
	sj.IsValidJavaMessage = false
}
