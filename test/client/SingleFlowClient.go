package main

import (
	"context"
	"flow/test/proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"sync"
	"time"
)

//根据ip,port,超时时间获取客户端对象
var clients = make(map[string]*grpc.ClientConn)
var clientLocker sync.Mutex

func GetSingleFlowClient(ip, port string, timeout int) (*proto.SingleFlowClient, error) {
	clientLocker.Lock()
	defer clientLocker.Unlock()
	var key = ip + ":" + port
	conn, ok := clients[key]
	conn.Connect()
	if ok {
		switch conn.GetState() {
		case connectivity.Ready:
			fallthrough
		case connectivity.Connecting:
			fallthrough
		case connectivity.Idle:
			//直接返回conn生成的client
			client := proto.NewSingleFlowClient(conn)
			return &client, nil
		case connectivity.Shutdown:
			fallthrough
		case connectivity.TransientFailure:
			conn.Close()
		default:
		}
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	conn, err := grpc.DialContext(ctx, key, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	clients[key] = conn
	newClient := proto.NewSingleFlowClient(conn)
	return &newClient, nil
}

//如何进行重连
func GetMessageFromServer(ip, port string, key string) error {
	client, err := GetSingleFlowClient(ip, port, 10)
	if err != nil {
		return err
	}
	var in proto.UpperRequest
	in.ID = []byte(key)
	in.ReqType = "1"
	message, err := (*client).GetMessage(context.Background(), &in)
	if err != nil {
		return err
	}
	for {
		recv, err := message.Recv()
		if err != nil {
			return err
		}
		fmt.Println(recv.String())
	}
	return nil
}
