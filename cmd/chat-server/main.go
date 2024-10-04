package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/Chigvero/chat-server/pkg/chat_v1"
	"github.com/brianvoe/gofakeit"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcHost = "localhost:"
	grpcPort = 50050
)

type server struct {
	desc.UnimplementedChatV1Server
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s%d", grpcHost, grpcPort))
	if err != nil {
		log.Fatalf("Error with listnening:%v\n", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to listen:%v\n", err)
	}
}

func (s *server) Create(_ context.Context, r *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Println(r.GetUsernames())
	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Delete(_ context.Context, r *desc.DeleteRequest) (*empty.Empty, error) {
	log.Println(r.GetId())
	return &empty.Empty{}, nil
}
func (s *server) SendMessage(_ context.Context, r *desc.SendMessageRequest) (*empty.Empty, error) {
	log.Println(desc.SendMessageRequest{
		From:      r.GetFrom(),
		Text:      r.GetText(),
		Timestamp: r.GetTimestamp(),
	})
	return nil, nil
}
