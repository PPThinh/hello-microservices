package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "hello-service/proto"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) GetHello(ctx context.Context, in *emptypb.Empty) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "hello"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	log.Printf("Hello Service listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
