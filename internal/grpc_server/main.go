package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	//	"google.golang.org/protobuf/types/known/timestamppb"

	"chat-server/internal/grpc_server/app"
	"chat-server/internal/grpc_server/config"
	"chat-server/internal/grpc_server/service"
	desc "chat-server/pkg/chat_v1"
)

//const (
//	port   = 50311
//	DbInfo = "host=localhost port=5090 user=chat password=1Pfem&dnX2 dbname=mydb sslmode=disable"
//)

type server struct {
	desc.UnimplementedChatV1Server
	svc service.Service
}

func (s *server) Store(ctx context.Context, req *desc.StoreRequest) (*desc.StoreResponse, error) {

	users := req.GetChat().GetUsernames()
	fmt.Println(len(users))
	if len(users) == 0 {
		return nil, status.Error(codes.InvalidArgument, "No users provided!")
	}

	id, err := s.svc.Store(ctx, req.GetChat())
	if err != nil {
		log.Printf("store chat: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to store chat: %v", err)
	}

	return &desc.StoreResponse{
		Id: id,
	}, nil

}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {

	id := req.GetId()
	//fmt.Println(id)
	data, err := s.svc.Get(ctx, id)
	if err != nil {
		log.Printf("store chat: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get chat: %v", err)
	}
	return &desc.GetResponse{Chat: data}, nil

}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {

	id := req.GetId()
	fmt.Println("Delete id:", id)
	err := s.svc.Delete(ctx, id)
	if err != nil {
		log.Printf("store chat: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to store chat: %v", err)
	}
	return &desc.DeleteResponse{}, nil

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
	desc.RegisterChatV1Server(s, &server{svc: application.Services})

	log.Printf("✅ Server up on port:%s", cfg.ServerPort)
	if err = s.Serve(listener); err != nil {
		log.Fatal("failed to server #{err}")
	}

}
