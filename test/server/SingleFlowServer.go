package main

import (
	"context"
	"flow/test/message"
	"flow/test/proto"
	"fmt"
	"google.golang.org/grpc"
	"io"
)

type SingleFlowServer struct {
}

//// 推送服务器消息
//	PutMessage(context.Context, *UpperRequest) (*UpperReply, error)
//	// 获取服务器消息
//	GetMessage(*UpperRequest, SingleFlow_GetMessageServer) error

func (server *SingleFlowServer) PutMessage(ctx context.Context, request *proto.UpperRequest) (response *proto.UpperReply, err error) {
	return
}

func (server *SingleFlowServer) GetMessage(request *proto.UpperRequest, stream proto.SingleFlow_GetMessageServer) error {
	if request != nil {
		fmt.Println("client ", string(request.ID), "connected!")
		message.Register(string(request.ID))
		var rChan = message.GetRChan(string(request.ID))
		var out proto.UpperReply
		for {
			for v := range *rChan {
				out.ReaType = []byte("resp")
				out.Contnet = []byte(v)
				err := stream.Send(&out)
				if err != nil {
					if err != io.EOF {
						return nil
					}
				}
			}
		}
	}

	return nil
}

func RegisterSingleFlowServer(s *grpc.Server) {
	proto.RegisterSingleFlowServer(s, &SingleFlowServer{})
}
