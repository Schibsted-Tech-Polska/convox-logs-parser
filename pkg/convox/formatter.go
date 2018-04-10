package convox

import (
	"strings"
	"fmt"
	"os"
	"io"
	"encoding/json"
)

type Formatter struct{}

var output io.Writer = os.Stdout

func (f *Formatter) Write(p []byte) (n int, err error) {
	var ps = string(p[:])
	//log.Println("Got write request: " + ps)
	var lines = strings.Split(ps, "\n")

	for _, line := range lines {
		if len(line) > 0 {
			pl := parseLogLine(line)
			fmt.Fprintln(output, pl)
		}
	}

	return len(p), nil
}

func isJSON(msg []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(msg, &js) == nil
}

type convoxFrame struct {
	Timestamp   string
	ServiceDesc string
	Message     string
}

func (cf *convoxFrame) IsJsonMessage() bool {
	return isJSON([]byte(cf.Message))
}

func (cf *convoxFrame) getMessageField() string {
	var jsRoot map[string]interface{}
	json.Unmarshal([]byte(cf.Message), &jsRoot)

	return jsRoot["message"].(string)
}

func getConvoxFrame(s string) convoxFrame {
	var sp = strings.Split(s, " ")
	var cf = convoxFrame{
		sp[0],
		sp[1],
		strings.Join(sp[2:], " "),
	}

	return cf
}

func parseLogLine(l string) string {
	cf := getConvoxFrame(l)
	if cf.IsJsonMessage() {
		var sa = []string{cf.Timestamp, cf.ServiceDesc, cf.getMessageField()}
		return strings.Join(sa, " ")
	} else {
		return "Date: " + cf.Timestamp + "\nServiceDesc: " + cf.ServiceDesc + "\nMessage: " + cf.Message + "\n"
	}
}
