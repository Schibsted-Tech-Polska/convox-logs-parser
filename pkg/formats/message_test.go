package formats

import (
	"testing"
)

type TestMessage struct {
	message string
}

func (tm TestMessage) String() string {
	return tm.message
}

func TestNew(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want Message
	}{
		{
			"Plain Message",
			args{"A message"},
			TestMessage{"A message"},
		},
		{
			"Simple JSON Message",
			args{"{\"message\": \"A message\"}"},
			TestMessage{"{\"message\": \"A message\"}"},
		},
		{
			"Java JSON Message missing thread",
			args{"{\"message\": \"A message\", \"level\": \"DE-BUG\", \"loggerName\": \"com.super.cls.Bro\"}"},
			TestMessage{"{\"message\": \"A message\", \"level\": \"DE-BUG\", \"loggerName\": \"com.super.cls.Bro\"}"},
		},
		{
			"Java JSON Message missing level",
			args{"{\"message\": \"A message\", \"loggerName\": \"com.super.cls.Bro\", \"thread\": \"a-thread-1\"}"},
			TestMessage{"{\"message\": \"A message\", \"loggerName\": \"com.super.cls.Bro\", \"thread\": \"a-thread-1\"}"},
		},
		{
			"Java JSON Message missing message",
			args{"{\"level\": \"DE-BUG\", \"loggerName\": \"com.super.cls.Bro\", \"thread\": \"a-thread-1\"}"},
			TestMessage{"{\"level\": \"DE-BUG\", \"loggerName\": \"com.super.cls.Bro\", \"thread\": \"a-thread-1\"}"},
		},
		{
			"Java JSON Message missing loggerName",
			args{"{\"message\": \"A message\", \"level\": \"DE-BUG\", \"thread\": \"a-thread-1\"}"},
			TestMessage{"{\"message\": \"A message\", \"level\": \"DE-BUG\", \"thread\": \"a-thread-1\"}"},
		},
		{
			"Java JSON Message missing thread and loggerName",
			args{"{\"message\": \"A message\", \"level\": \"DE-BUG\"}"},
			TestMessage{"{\"message\": \"A message\", \"level\": \"DE-BUG\"}"},
		},
		{
			"Java JSON Message sufficient fields",
			args{"{\"message\": \"A message\", \"level\": \"DE-BUG\", \"loggerName\": \"com.super.cls.Bro\", \"thread\": \"a-thread-1\"}"},
			TestMessage{"[a-thread-1] DE-BUG com.super.cls.Bro - A message"},
		},
		{
			"Java JSON Message all fields",
			args{"{\"thread\":\"vert.x-eventloop-thread-0\",\"level\":\"INFO\",\"loggerName\":\"com.example.demo.MainVerticle\",\"message\":\"Received request!\",\"endOfBatch\":true,\"loggerFqcn\":\"io.vertx.core.logging.Logger\",\"instant\":{\"epochSecond\":1523444015,\"nanoOfSecond\":299000000},\"contextMap\":{},\"threadId\":15,\"threadPriority\":5,\"hostname\":\"4cbd459504f5\",\"service\":\"ci-example\",\"release\":\"RMBYTXJUUGO\"}"},
			TestMessage{"[vert.x-eventloop-thread-0] INFO com.example.demo.MainVerticle - Received request!"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.message); got.String() != tt.want.String() {
				t.Errorf("New() = [%v], want [%v]", got, tt.want)
			}
		})
	}
}
