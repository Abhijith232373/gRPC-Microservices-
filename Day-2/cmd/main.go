package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq" // Postgres driver
	"google.golang.org/grpc"
	"user-service/internal"
	"user-service/proto"
)

func main() {
connStr := "postgres://postgres:232373@localhost:5432/grpc?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	userServer := &internal.UserServer{DB: db}
	proto.RegisterUserServiceServer(grpcServer, userServer)
	log.Println("Server running on :50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
