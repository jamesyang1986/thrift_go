package main

import (
	"context"
	"fmt"
	"net"
	"os"
	rpc "thrift/demo/rpc"
	thrift "thrift/src"
	"time"
)

func main() {
	startTime := currentTimeMillis()
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "9098"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport, err := transportFactory.GetTransport(transport)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error get the transport", err)
		os.Exit(1)
	}
	client := rpc.NewRpcServiceClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:9098", " ", err)
		os.Exit(1)
	}
	defer transport.Close()

	for i := 0; i < 1000; i++ {
		paramMap := make(map[string]string)
		paramMap["name"] = "qinerg"
		paramMap["passwd"] = "123456"
		r1, e1 := client.FunCall(context.Background(), currentTimeMillis(), "login", paramMap)
		fmt.Println(i, "Call->", r1, e1)
	}

	endTime := currentTimeMillis()
	fmt.Println("Program exit. time->", endTime, startTime, (endTime - startTime))
}

// 转换成毫秒
func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
