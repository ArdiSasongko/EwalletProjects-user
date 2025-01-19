package api

import (
	"log"
	"net"

	"github.com/ArdiSasongko/EwalletProjects-user/internal/config/logger"
	"github.com/ArdiSasongko/EwalletProjects-user/internal/env"

	"google.golang.org/grpc"

	"github.com/joho/godotenv"
)

func SetupGRPC() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	// set up logrus
	logrus := logger.NewLogger()
	cfg := Config{
		addr:   env.GetEnv("ADDR_GRPC", ":5000"),
		logger: logrus,
	}

	lis, err := net.Listen("tcp", cfg.addr)
	if err != nil {
		cfg.logger.Fatalf("failed listen grpc port, err:%v", err)
	}

	server := grpc.NewServer()

	cfg.logger.Printf("grpc server has running, port:%v", cfg.addr)

	if err := server.Serve(lis); err != nil {
		cfg.logger.Fatalf("failed to starting grpc server, err:%v", err)
	}
}
