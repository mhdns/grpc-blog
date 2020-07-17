package main

import (
	"blog/server/userpb"
	"context"
	"log"
)

func (s *server) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	name := req.GetName()
	email := req.GetEmail()
	password := req.GetPassword()

	// Generate salt
	salt := "1312312313"

	// Generate hashed pw
	hashedPassword := password

	query := readSQL("queries/create_user.sql")
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, email, name, hashedPassword, salt)

	var createdID, createdName string

	err = row.Scan(&createdID, &createdName)
	if err != nil {
		log.Fatal(err)
	}

	return &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:   createdID,
			Name: createdName,
		},
		Msg:     "Created user successfully",
		Success: true,
	}, nil
}

func (s *server) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	return &userpb.GetUserResponse{}, nil
}

func (s *server) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	return &userpb.UpdateUserResponse{}, nil
}

func (s *server) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	return &userpb.DeleteUserResponse{}, nil
}
