package main

import (
	"chat-server/internal/api/chat_v1"
	"chat-server/internal/config"
	"chat-server/internal/config/env"
	"chat-server/internal/repository"
	"chat-server/internal/service"
	"flag"
	desc "github.com/Chigvero/chat-server/pkg/chat_v1"
	"github.com/jackc/pgx/v5"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	grpcHost = "localhost:"
	grpcPort = 50050
)

var configString string

func init() {
	flag.StringVar(&configString, "config", ".env", "to specify the config file")
}

func main() {
	//config ready
	flag.Parse()
	err := config.Load(configString)
	if err != nil {
		log.Fatalf("Error with loading configPath:%v\n", err)
	}
	pgConfig, err := env.NewPgConfig()
	if err != nil {
		log.Fatalf("failed to get grpcConfig:%v", err)
	}
	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpcConfig:%v", err)
	}

	//server ready
	ctx := context.Background()
	pgxConn, err := pgx.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewRepository(pgxConn)
	services := service.NewService(repo)
	implemTransport := chat_v1.NewImplementation(services)

	//Server starting
	lis, err := net.Listen("tcp", grpcConfig.Address())
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
