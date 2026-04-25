package internal

import (
	"context"
	"database/sql"
	"user-service/proto"
)

type UserServer struct {
	proto.UnimplementedUserServiceServer
	DB *sql.DB
}

func (s *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.User, error) {
	u := &proto.User{Name: req.Name, Email: req.Email}
	return u, s.DB.QueryRowContext(ctx, "INSERT INTO users(name,email) VALUES($1,$2) RETURNING id", u.Name, u.Email).Scan(&u.Id)
}

func (s *UserServer) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.User, error) {
	u := &proto.User{}
	return u, s.DB.QueryRowContext(ctx, "SELECT id,name,email FROM users WHERE id=$1", req.Id).Scan(&u.Id, &u.Name, &u.Email)
}