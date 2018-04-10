package convox

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"bytes"
	"bufio"
)

var standardMessages = []struct {
	in       []byte
	expected []byte
}{
	{
		[]byte("12:23:45 ServiceDesc This is sample message"),
		[]byte("Date: 12:23:45\nServiceDesc: ServiceDesc\nMessage: This is sample message\n\n"),
	},
	{
		[]byte("31:22:12 service/app:RELEASE/docker-id {\"message\": \"It's a message\", \"random\": \"Some random field\"}"),
		[]byte("31:22:12 service/app:RELEASE/docker-id It's a message\n"),
	},
}


var ft = Formatter{}
var actual bytes.Buffer
var writer = bufio.NewWriter(&actual)

func setup() {
	writer.Flush()
	actual.Reset()
	output = writer
}

func TestMessages(t *testing.T) {
	for _, tt := range standardMessages {
		setup()

		n, err := ft.Write(tt.in)
		writer.Flush()

		assert.Equal(t, tt.expected, actual.Bytes())
		assert.Equal(t, len(tt.in), n)
		assert.NoError(t, err)
	}
}