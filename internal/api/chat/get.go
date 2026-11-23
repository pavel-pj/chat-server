package chat

import (
	"context"
	"log"

	desc "github.com/pavel-pj/go-chat/pkg/chat_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ChatImplementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {

	id := req.GetId()
	//fmt.Println(len(users))

	chat, err := s.serv.Get(ctx, id)
	if err != nil {
		log.Printf("store chat: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to store chat: %v", err)
	}

	chat

	return &desc.GetResponse{
		Id: id,
	}, nil

}
