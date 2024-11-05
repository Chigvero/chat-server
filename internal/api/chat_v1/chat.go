package chat_v1

import (
	"chat-server/internal/service"
	"context"
	"errors"
	desc "github.com/Chigvero/chat-server/pkg/chat_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
)

type Implementation struct {
	desc.UnimplementedChatV1Server
	service *service.Service
}

func NewImplementation(s *service.Service) *Implementation {
	return &Implementation{
		UnimplementedChatV1Server: desc.UnimplementedChatV1Server{},
		service:                   s,
	}
}
func (im *Implementation) Create(ctx context.Context, request *desc.CreateRequest) (*desc.CreateResponse, error) {
	usrs := request.GetUsernames()
	if len(usrs) == 0 {
		return nil, errors.New("Invalid username's data ")
	}
	log.Println(usrs)
	id, err := im.service.ChatService.Create(ctx, usrs)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &desc.CreateResponse{Id: id}, nil
}

func (im *Implementation) Delete(ctx context.Context, r *desc.DeleteRequest) (*empty.Empty, error) {
	id := r.GetId()
	if id == 0 {
		return nil, errors.New("Invalid id format")
	}
	err := im.service.ChatService.Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
func (im *Implementation) SendMessage(_ context.Context, r *desc.SendMessageRequest) (*empty.Empty, error) {
	log.Println(desc.SendMessageRequest{
		From:      r.GetFrom(),
		Text:      r.GetText(),
		Timestamp: r.GetTimestamp(),
	})
	return nil, nil
}
