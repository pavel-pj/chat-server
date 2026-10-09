package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	//	"google.golang.org/protobuf/types/known/timestamppb"

	desc "chat-server/pkg/chat_v1"
)

const port = 50311

type server struct {
	desc.UnimplementedChatV1Server
}

func (s *server) Store(ctx context.Context, req *desc.StoreRequest) (*desc.StoreResponse, error) {
	return &desc.StoreResponse{
		Id: int64(252),
	}, nil

}

func main() {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("Server could not start")
	}

	log.Println("Server up on port:", port)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("Server listening at %d", port)
	if err = s.Serve(listener); err != nil {
		log.Fatal("failed to server #{err}")
	}

}
