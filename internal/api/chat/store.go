package chat

import (
	"context"
	"log"

	desc "github.com/pavel-pj/go-chat/pkg/chat_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ChatImplementation) Store(ctx context.Context, req *desc.StoreRequest) (*desc.StoreResponse, error) {

	users := req.GetChat().GetUsernames()
	//fmt.Println(len(users))
	if len(users) == 0 {
		return nil, status.Error(codes.InvalidArgument, "No users provided!")
	}

	id, err := s.serv.Store(ctx, req.GetChat())
	if err != nil {
		log.Printf("store chat: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to store chat: %v", err)
	}

	return &desc.StoreResponse{
		Id: id,
	}, nil

}
