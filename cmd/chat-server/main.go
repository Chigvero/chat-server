package main

import (
	"chat-server/internal/api/chat_v1"
	"chat-server/internal/repository"
	"chat-server/internal/service"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
	"log"
	"net"
	"os"

	desc "github.com/Chigvero/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcHost = "localhost:"
	grpcPort = 50050
)

func main() {
	//server ready
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatal(err)
	}
	connStr := os.Getenv("MIGRATION_DSN_L")
	if len(connStr) == 0 {
		log.Fatalf("Len connstr is 0")
	}
	ctx := context.Background()
	pgxConn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewRepository(pgxConn)
	services := service.NewService(repo)
	implemTransport := chat_v1.NewImplementation(services)

	//Server starting
	lis, err := net.Listen("tcp", fmt.Sprintf("%s%d", grpcHost, grpcPort))
	if err != nil {
		log.Fatalf("Error with listnening:%v\n", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, implemTransport)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to listen:%v\n", err)
	}
}
