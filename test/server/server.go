package main

import (
	"fmt"
	"google.golang.org/grpc"
	"math"
	"net"
)

func RunServer(port string) {
	address := ":" + port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("server print error:%s", err)
	}
	optionMaxRecvMsgSize := grpc.MaxRecvMsgSize(math.MaxInt32)
	s := grpc.NewServer(optionMaxRecvMsgSize)
	//将服务注册到grpc服务中
	RegisterSingleFlowServer(s)
	s.Serve(lis)
}
