package mygrpc

import (
	"fmt"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func NewGRPCClient(host string, port string) (*grpc.ClientConn, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found, relying on environment variables")
	}
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err)
	}

	return conn, nil
}
