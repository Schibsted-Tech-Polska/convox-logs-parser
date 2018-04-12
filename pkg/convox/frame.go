package convox

import (
	"fmt"
	"github.com/Schibsted-Tech-Polska/convox-logs-parser/pkg/formats"
	"strings"
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
