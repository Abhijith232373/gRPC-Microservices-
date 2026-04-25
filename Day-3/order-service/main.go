package main

import (
    "log"
    "net"
    "google.golang.org/grpc"

    // These must start with your module name 'my-microservices'
    "my-microservices/order-service/service"
    pbUser "my-microservices/proto/user"
    pbOrder "my-microservices/proto/order"
)

funcA main() {
	// 2. Setup connection to User Service (Order Service acts as a CLIENT here)
	// 'Dial' starts the TCP connection process
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {      
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	userClient := pbUser.NewUserServiceClient(conn)

	// 3. Start Order Server (Order Service acts as a SERVER here)
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	s := grpc.NewServer()

	// 4. FIX: Use the struct from the service package and pass the client
	orderSvc := &service.OrderServer{
		UserClient: userClient,
	}
	
	pbOrder.RegisterOrderServiceServer(s, orderSvc)

	log.Println("Order Service running on :50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}