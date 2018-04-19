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
		//"thrown":{"commonElementCount":0,"localizedMessage":"Connection reset by peer","message":"Connection reset by peer","name":"java.io.IOException","extendedStackTrace":"java.io.IOException: Connection reset by peer\n\tat sun.nio.ch.FileDispatcherImpl.read0(Native Method) ~[?:1.8.0_162]\n\tat sun.nio.ch.SocketDispatcher.read(SocketDispatcher.java:39) ~[?:1.8.0_162]\n\tat sun.nio.ch.IOUtil.readIntoNativeBuffer(IOUtil.java:223) ~[?:1.8.0_162]\n\tat sun.nio.ch.IOUtil.read(IOUtil.java:192) ~[?:1.8.0_162]\n\tat sun.nio.ch.SocketChannelImpl.read(SocketChannelImpl.java:380) ~[?:1.8.0_162]\n\tat io.netty.buffer.PooledUnsafeDirectByteBuf.setBytes(PooledUnsafeDirectByteBuf.java:288) ~[subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.buffer.AbstractByteBuf.writeBytes(AbstractByteBuf.java:1108) ~[subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.channel.socket.nio.NioSocketChannel.doReadBytes(NioSocketChannel.java:345) ~[subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.channel.nio.AbstractNioByteChannel$NioByteUnsafe.read(AbstractNioByteChannel.java:126) [subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.channel.nio.NioEventLoop.processSelectedKey(NioEventLoop.java:645) [subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.channel.nio.NioEventLoop.processSelectedKeysOptimized(NioEventLoop.java:580) [subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.channel.nio.NioEventLoop.processSelectedKeys(NioEventLoop.java:497) [subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.channel.nio.NioEventLoop.run(NioEventLoop.java:459) [subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.util.concurrent.SingleThreadEventExecutor$5.run(SingleThreadEventExecutor.java:886) [subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.util.concurrent.FastThreadLocalRunnable.run(FastThreadLocalRunnable.java:30) [subscription-api-jar-with-dependencies.jar:?]\n\tat java.lang.Thread.run(Thread.java:748) [?:1.8.0_162]\n"}
		{
			"Java JSON thrown no extendedStackTrace",
			args{"{\"thread\":\"vert.x-eventloop-thread-0\",\"level\":\"INFO\",\"loggerName\":\"com.example.demo.MainVerticle\",\"message\":\"Received request!\",\"endOfBatch\":true,\"loggerFqcn\":\"io.vertx.core.logging.Logger\",\"instant\":{\"epochSecond\":1523444015,\"nanoOfSecond\":299000000},\"contextMap\":{},\"threadId\":15,\"threadPriority\":5,\"hostname\":\"4cbd459504f5\",\"service\":\"ci-example\",\"release\":\"RMBYTXJUUGO\",\"thrown\":{\"commonElementCount\":0,\"localizedMessage\":\"Connection reset by peer\",\"message\":\"Connection reset by peer\",\"name\":\"java.io.IOException\"}}"},
			TestMessage{"[vert.x-eventloop-thread-0] INFO com.example.demo.MainVerticle - Received request!"},
		},
		{
			"Java JSON thrown",
			args{"{\"thread\":\"vert.x-eventloop-thread-0\",\"level\":\"INFO\",\"loggerName\":\"com.example.demo.MainVerticle\",\"message\":\"Received request!\",\"endOfBatch\":true,\"loggerFqcn\":\"io.vertx.core.logging.Logger\",\"instant\":{\"epochSecond\":1523444015,\"nanoOfSecond\":299000000},\"contextMap\":{},\"threadId\":15,\"threadPriority\":5,\"hostname\":\"4cbd459504f5\",\"service\":\"ci-example\",\"release\":\"RMBYTXJUUGO\",\"thrown\":{\"commonElementCount\":0,\"localizedMessage\":\"Connection reset by peer\",\"message\":\"Connection reset by peer\",\"name\":\"java.io.IOException\",\"extendedStackTrace\":\"java.io.IOException: Connection reset by peer\\n\\tat sun.nio.ch.FileDispatcherImpl.read0(Native Method) ~[?:1.8.0_162]\\n\\tat sun.nio.ch.SocketDispatcher.read(SocketDispatcher.java:39) ~[?:1.8.0_162]\\n\\tat sun.nio.ch.IOUtil.readIntoNativeBuffer(IOUtil.java:223) ~[?:1.8.0_162]\\n\\tat sun.nio.ch.IOUtil.read(IOUtil.java:192) ~[?:1.8.0_162]\\n\\tat sun.nio.ch.SocketChannelImpl.read(SocketChannelImpl.java:380) ~[?:1.8.0_162]\\n\\tat io.netty.buffer.PooledUnsafeDirectByteBuf.setBytes(PooledUnsafeDirectByteBuf.java:288) ~[subscription-api-jar-with-dependencies.jar:?]\\n\\tat io.netty.buffer.AbstractByteBuf.writeBytes(AbstractByteBuf.java:1108) ~[subscription-api-jar-with-dependencies.jar:?]\\n\\tat io.netty.channel.socket.nio.NioSocketChannel.doReadBytes(NioSocketChannel.java:345) ~[subscription-api-jar-with-dependencies.jar:?]\\n\\tat io.netty.channel.nio.AbstractNioByteChannel$NioByteUnsafe.read(AbstractNioByteChannel.java:126) [subscription-api-jar-with-dependencies.jar:?]\\n\\tat io.netty.channel.nio.NioEventLoop.processSelectedKey(NioEventLoop.java:645) [subscription-api-jar-with-dependencies.jar:?]\\n\\tat io.netty.channel.nio.NioEventLoop.processSelectedKeysOptimized(NioEventLoop.java:580) [subscription-api-jar-with-dependencies.jar:?]\\n\\tat io.netty.channel.nio.NioEventLoop.processSelectedKeys(NioEventLoop.java:497) [subscription-api-jar-with-dependencies.jar:?]\\n\\tat io.netty.channel.nio.NioEventLoop.run(NioEventLoop.java:459) [subscription-api-jar-with-dependencies.jar:?]\\n\\tat io.netty.util.concurrent.SingleThreadEventExecutor$5.run(SingleThreadEventExecutor.java:886) [subscription-api-jar-with-dependencies.jar:?]\\n\\tat io.netty.util.concurrent.FastThreadLocalRunnable.run(FastThreadLocalRunnable.java:30) [subscription-api-jar-with-dependencies.jar:?]\\n\\tat java.lang.Thread.run(Thread.java:748) [?:1.8.0_162]\\n\"}}"},
			TestMessage{"[vert.x-eventloop-thread-0] INFO com.example.demo.MainVerticle - Received request!\njava.io.IOException: Connection reset by peer\n\tat sun.nio.ch.FileDispatcherImpl.read0(Native Method) ~[?:1.8.0_162]\n\tat sun.nio.ch.SocketDispatcher.read(SocketDispatcher.java:39) ~[?:1.8.0_162]\n\tat sun.nio.ch.IOUtil.readIntoNativeBuffer(IOUtil.java:223) ~[?:1.8.0_162]\n\tat sun.nio.ch.IOUtil.read(IOUtil.java:192) ~[?:1.8.0_162]\n\tat sun.nio.ch.SocketChannelImpl.read(SocketChannelImpl.java:380) ~[?:1.8.0_162]\n\tat io.netty.buffer.PooledUnsafeDirectByteBuf.setBytes(PooledUnsafeDirectByteBuf.java:288) ~[subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.buffer.AbstractByteBuf.writeBytes(AbstractByteBuf.java:1108) ~[subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.channel.socket.nio.NioSocketChannel.doReadBytes(NioSocketChannel.java:345) ~[subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.channel.nio.AbstractNioByteChannel$NioByteUnsafe.read(AbstractNioByteChannel.java:126) [subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.channel.nio.NioEventLoop.processSelectedKey(NioEventLoop.java:645) [subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.channel.nio.NioEventLoop.processSelectedKeysOptimized(NioEventLoop.java:580) [subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.channel.nio.NioEventLoop.processSelectedKeys(NioEventLoop.java:497) [subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.channel.nio.NioEventLoop.run(NioEventLoop.java:459) [subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.util.concurrent.SingleThreadEventExecutor$5.run(SingleThreadEventExecutor.java:886) [subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.util.concurrent.FastThreadLocalRunnable.run(FastThreadLocalRunnable.java:30) [subscription-api-jar-with-dependencies.jar:?]\n\tat java.lang.Thread.run(Thread.java:748) [?:1.8.0_162]\n"},
		},
		{
			"Java JSON thrown no message",
			args{"{\"thread\":\"vert.x-eventloop-thread-0\",\"level\":\"INFO\",\"loggerName\":\"com.example.demo.MainVerticle\",\"endOfBatch\":true,\"loggerFqcn\":\"io.vertx.core.logging.Logger\",\"instant\":{\"epochSecond\":1523444015,\"nanoOfSecond\":299000000},\"contextMap\":{},\"threadId\":15,\"threadPriority\":5,\"hostname\":\"4cbd459504f5\",\"service\":\"ci-example\",\"release\":\"RMBYTXJUUGO\",\"thrown\":{\"commonElementCount\":0,\"localizedMessage\":\"Connection reset by peer\",\"message\":\"Connection reset by peer\",\"name\":\"java.io.IOException\",\"extendedStackTrace\":\"java.io.IOException: Connection reset by peer\\n\\tat sun.nio.ch.FileDispatcherImpl.read0(Native Method) ~[?:1.8.0_162]\\n\\tat sun.nio.ch.SocketDispatcher.read(SocketDispatcher.java:39) ~[?:1.8.0_162]\\n\\tat sun.nio.ch.IOUtil.readIntoNativeBuffer(IOUtil.java:223) ~[?:1.8.0_162]\\n\\tat sun.nio.ch.IOUtil.read(IOUtil.java:192) ~[?:1.8.0_162]\\n\\tat sun.nio.ch.SocketChannelImpl.read(SocketChannelImpl.java:380) ~[?:1.8.0_162]\\n\\tat io.netty.buffer.PooledUnsafeDirectByteBuf.setBytes(PooledUnsafeDirectByteBuf.java:288) ~[subscription-api-jar-with-dependencies.jar:?]\\n\\tat io.netty.buffer.AbstractByteBuf.writeBytes(AbstractByteBuf.java:1108) ~[subscription-api-jar-with-dependencies.jar:?]\\n\\tat io.netty.channel.socket.nio.NioSocketChannel.doReadBytes(NioSocketChannel.java:345) ~[subscription-api-jar-with-dependencies.jar:?]\\n\\tat io.netty.channel.nio.AbstractNioByteChannel$NioByteUnsafe.read(AbstractNioByteChannel.java:126) [subscription-api-jar-with-dependencies.jar:?]\\n\\tat io.netty.channel.nio.NioEventLoop.processSelectedKey(NioEventLoop.java:645) [subscription-api-jar-with-dependencies.jar:?]\\n\\tat io.netty.channel.nio.NioEventLoop.processSelectedKeysOptimized(NioEventLoop.java:580) [subscription-api-jar-with-dependencies.jar:?]\\n\\tat io.netty.channel.nio.NioEventLoop.processSelectedKeys(NioEventLoop.java:497) [subscription-api-jar-with-dependencies.jar:?]\\n\\tat io.netty.channel.nio.NioEventLoop.run(NioEventLoop.java:459) [subscription-api-jar-with-dependencies.jar:?]\\n\\tat io.netty.util.concurrent.SingleThreadEventExecutor$5.run(SingleThreadEventExecutor.java:886) [subscription-api-jar-with-dependencies.jar:?]\\n\\tat io.netty.util.concurrent.FastThreadLocalRunnable.run(FastThreadLocalRunnable.java:30) [subscription-api-jar-with-dependencies.jar:?]\\n\\tat java.lang.Thread.run(Thread.java:748) [?:1.8.0_162]\\n\"}}"},
			TestMessage{"[vert.x-eventloop-thread-0] INFO com.example.demo.MainVerticle - java.io.IOException: Connection reset by peer\n\tat sun.nio.ch.FileDispatcherImpl.read0(Native Method) ~[?:1.8.0_162]\n\tat sun.nio.ch.SocketDispatcher.read(SocketDispatcher.java:39) ~[?:1.8.0_162]\n\tat sun.nio.ch.IOUtil.readIntoNativeBuffer(IOUtil.java:223) ~[?:1.8.0_162]\n\tat sun.nio.ch.IOUtil.read(IOUtil.java:192) ~[?:1.8.0_162]\n\tat sun.nio.ch.SocketChannelImpl.read(SocketChannelImpl.java:380) ~[?:1.8.0_162]\n\tat io.netty.buffer.PooledUnsafeDirectByteBuf.setBytes(PooledUnsafeDirectByteBuf.java:288) ~[subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.buffer.AbstractByteBuf.writeBytes(AbstractByteBuf.java:1108) ~[subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.channel.socket.nio.NioSocketChannel.doReadBytes(NioSocketChannel.java:345) ~[subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.channel.nio.AbstractNioByteChannel$NioByteUnsafe.read(AbstractNioByteChannel.java:126) [subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.channel.nio.NioEventLoop.processSelectedKey(NioEventLoop.java:645) [subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.channel.nio.NioEventLoop.processSelectedKeysOptimized(NioEventLoop.java:580) [subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.channel.nio.NioEventLoop.processSelectedKeys(NioEventLoop.java:497) [subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.channel.nio.NioEventLoop.run(NioEventLoop.java:459) [subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.util.concurrent.SingleThreadEventExecutor$5.run(SingleThreadEventExecutor.java:886) [subscription-api-jar-with-dependencies.jar:?]\n\tat io.netty.util.concurrent.FastThreadLocalRunnable.run(FastThreadLocalRunnable.java:30) [subscription-api-jar-with-dependencies.jar:?]\n\tat java.lang.Thread.run(Thread.java:748) [?:1.8.0_162]\n"},
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
