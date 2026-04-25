package service

import (
    "context"
    "errors"
    // Use these two only:
    user "my-microservices/proto/user"
    order "my-microservices/proto/order"
)


type OrderServer struct {
    UserClient pbUser.UserServiceClient
}

funcA (s *OrderServer) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.OrderResponse, error) {
    // 1. CALLING THE OTHER SERVICE
    // We send the UserID to the User Service via gRPC
    res, err := s.UserClient.ValidateUser(ctx, &user.ValidateRequest{UserId: req.UserId})
    
    // 2. ERROR HANDLING
    if err != nil || !res.IsValid {
        return nil, errors.New("authentication failed: user is invalid")
    }

    // 3. SUCCESS LOGIC
    return &order.OrderResponse{OrderId: "ORD-123", Status: "Created"}, nil
}