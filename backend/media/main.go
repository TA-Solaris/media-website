package main

import (
	"context"
	"log"
	"net"

	"github.com/TA-Solaris/common"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {
	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer l.Close()

	upload := NewUpload()
	svc := NewService(upload)

	RegisterGRPCHandler(grpcServer)

	svc.Upload(context.Background())

	log.Println("gRPC Server Started at ", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err)
	}
}
