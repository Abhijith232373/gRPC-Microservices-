package main

import (
    "context"
    "log"
    "net" // UNCOMMENTED: Needed for TCP connection
    "google.golang.org/grpc" // UNCOMMENTED: Needed for gRPC server
    user "my-microservices/proto/user" 
)

type server struct {
    user.UnimplementedUserServiceServer
}

func (s *server) ValidateUser(ctx context.Context, req *user.ValidateRequest) (*user.ValidateResponse, error) {
    log.Printf("Received Validation Request for UserID: %s", req.UserId)
    
    // Simple logic for the review: 123 is valid, everything else is not
    isValid := req.UserId == "123"
    
    return &user.ValidateResponse{IsValid: isValid}, nil
}

func main() {
    // 1. Create a TCP listener on port 50051
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    // 2. Create the gRPC server instance
    s := grpc.NewServer()

    // 3. Register our service logic with the server
    user.RegisterUserServiceServer(s, &server{})

    log.Println("✅ User Service is running on port :50051...")
    
    // 4. Start serving requests
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}