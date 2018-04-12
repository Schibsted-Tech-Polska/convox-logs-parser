package convox

import (
	"github.com/radekl/convox-json-logs/pkg/formats"
	"strings"
	"fmt"
)

type convoxFrame struct {
	Timestamp   string
	ServiceDesc string
	Message     formats.Message
}

func getConvoxFrame(s string) convoxFrame {
	var sp = strings.Split(s, " ")
	var cf = convoxFrame{
		sp[0],
		sp[1],
		formats.New(strings.Join(sp[2:], " ")),
	}

	return cf
}

func (cf convoxFrame) String() string {
	return fmt.Sprintf("%s %s %s", cf.Timestamp, cf.ServiceDesc, cf.Message)
}
