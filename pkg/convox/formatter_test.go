package convox

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestFormatter_Write(t *testing.T) {
	var buf bytes.Buffer
	output = bufio.NewWriter(&buf)

	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		f       *Formatter
		args    args
		wantErr bool
		wantMsg string
	}{
		{
			"Standard message",
			&Formatter{},
			args{[]byte("12:23:45 ServiceDesc This is sample message")},
			false,
			"12:23:45 ServiceDesc This is sample message",
		},
		{
			"JSON message",
			&Formatter{},
			args{[]byte("31:22:12 service/app:RELEASE/docker-id {\"message\": \"It's a message\", \"random\": \"Some random field\"}")},
			false,
			"31:22:12 service/app:RELEASE/docker-id {\"message\": \"It's a message\", \"random\": \"Some random field\"}",
		},
		{
			"Java Log4j2 JSON message",
			&Formatter{},
			args{[]byte("2018-04-11T10:53:35Z service/app:RMBYTXJUUGO/2b29bc027044 {\"thread\":\"vert.x-eventloop-thread-0\",\"level\":\"INFO\",\"loggerName\":\"com.example.demo.MainVerticle\",\"message\":\"Received request!\",\"endOfBatch\":true,\"loggerFqcn\":\"io.vertx.core.logging.Logger\",\"instant\":{\"epochSecond\":1523444015,\"nanoOfSecond\":299000000},\"contextMap\":{},\"threadId\":15,\"threadPriority\":5,\"hostname\":\"4cbd459504f5\",\"service\":\"ci-example\",\"release\":\"RMBYTXJUUGO\"}")},
			false,
			"2018-04-11T10:53:35Z service/app:RMBYTXJUUGO/2b29bc027044 [vert.x-eventloop-thread-0] INFO com.example.demo.MainVerticle - Received request!",
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			var gotMsg bytes.Buffer
			var writer = bufio.NewWriter(&gotMsg)

			output = writer

			f := &Formatter{}
			gotN, err := f.Write(tt.args.p)
			writer.Flush()

			if (err != nil) != tt.wantErr {
				t.Errorf("Formatter.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != len(tt.args.p) {
				t.Errorf("Formatter.Write() = %v, want %v", gotN, len(tt.args.p))
				return
			}
			if strings.TrimSuffix(gotMsg.String(), "\n") != tt.wantMsg {
				t.Errorf("Formatter.Write() output = %v, want %v", strings.TrimSuffix(gotMsg.String(), "\n"), tt.wantMsg)
			}
		})
	}
}
