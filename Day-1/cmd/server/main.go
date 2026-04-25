package main

import (
	"context"
	"net"
	"google.golang.org/grpc"
	"hello-grpc/pb"
)

type server struct{ pb.UnimplementedHelloServiceServer }

func (s *server) SayHello(c context.Context, r *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{Message: "Hello " + r.GetName()}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	s.Serve(lis)
}