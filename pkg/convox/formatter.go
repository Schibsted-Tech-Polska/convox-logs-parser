package convox

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Formatter is a type which acts as a io.Writer interface. This way convox-logs-parser is able
// to fetch all bytestreams in order to parse them
type Formatter struct{}

var output io.Writer = os.Stdout

func (f *Formatter) Write(p []byte) (n int, err error) {
	var ps = string(p[:])
	var lines = strings.Split(ps, "\n")

	for _, line := range lines {
		if len(line) > 0 {
			pl := getConvoxFrame(line)
			_,_ = fmt.Fprintln(output, pl)
		}
	}

	return len(p), nil
}
