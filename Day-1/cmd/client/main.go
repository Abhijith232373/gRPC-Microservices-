package main

import (
	"context"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"hello-grpc/pb"
)

func main() {
	conn, _ := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	res, err := pb.NewHelloServiceClient(conn).SayHello(context.Background(), &pb.UserRequest{Name: "Abijith"})
	if err != nil {
		log.Fatal(err)
	}
	log.Print(res.GetMessage())
}