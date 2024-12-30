package main

import (
	"context"
	"log"

	pb "github.com/TA-Solaris/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcHandler struct {
	pb.UnimplementedMediaServiceServer

	service MediaService
}

func RegisterGRPCHandler(grpcServer *grpc.Server, service MediaService) {
	handler := &grpcHandler{
		service: service,
	}
	pb.RegisterMediaServiceServer(grpcServer, handler)
}

func (h *grpcHandler) Upload(ctx context.Context, p *pb.UploadRequest) (*pb.Media, error) {
	log.Printf("Upload request recieved from user with ID %s", p.UserID)

	// TODO Implement

	return nil, status.Errorf(codes.Unimplemented, "Method Upload not implemented")
}
