package chat_v1

import (
	"chat-server/internal/service"
	"context"
	"errors"
	desc "github.com/Chigvero/chat-server/pkg/chat_v1"
)

type Implementation struct {
	desc.UnimplementedChatV1Server
	service service.Service
}

func (im *Implementation) Create(ctx context.Context, request *desc.CreateRequest) (*desc.CreateResponse, error) {
	usrs := request.GetUsernames()
	if len(usrs) == 0 {
		return nil, errors.New("Invalid username's data ")
	}
	im.service.ChatService.Create(ctx, usrs)

}
