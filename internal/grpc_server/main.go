package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	//	"google.golang.org/protobuf/types/known/timestamppb"

	"chat-server/internal/grpc_server/app"
	"chat-server/internal/grpc_server/config"
	desc "chat-server/pkg/chat_v1"
)

//const (
//	port   = 50311
//	DbInfo = "host=localhost port=5090 user=chat password=1Pfem&dnX2 dbname=mydb sslmode=disable"
//)

type server struct {
	desc.UnimplementedChatV1Server
}

func (s *server) Store(ctx context.Context, req *desc.StoreRequest) (*desc.StoreResponse, error) {

	users := req.GetChat().GetUsernames()
	log.Println(users)

	return &desc.StoreResponse{
		Id: int64(252),
	}, nil

}

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Создаём приложение
	application, err := app.NewApp(*cfg)
	if err != nil {
		log.Fatal("Failed to create app: ", err)
	}
	defer func() {
		if err := application.Close(); err != nil {
			log.Printf("Warning: failed to close app: %v", err)
		}
	}()

	log.Printf("✅ Successful connect to PostgreSQL")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.ServerPort))
	if err != nil {
		log.Fatal("Server could not start")
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("✅ Server up on port:%s", cfg.ServerPort)
	if err = s.Serve(listener); err != nil {
		log.Fatal("failed to server #{err}")
	}

}
