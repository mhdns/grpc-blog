package main

import (
	"blog/server/blogpb"
	"context"
	"database/sql"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) CreateUser(ctx context.Context, req *blogpb.CreateUserRequest) (*blogpb.CreateUserResponse, error) {
	name := req.GetName()
	email := req.GetEmail()
	password := req.GetPassword()

	// Generate salt
	salt := "1312312313"

	// Generate hashed pw
	hashedPassword := password

	row := s.createUser.QueryRowContext(ctx, email, name, hashedPassword, salt)

	var createdID, createdName, createdEmail string

	err := row.Scan(&createdID, &createdName, &createdEmail)
	if err != nil {
		log.Fatal("create user:", err)
	}

	return &blogpb.CreateUserResponse{
		User: &blogpb.User{
			Id:   createdID,
			Name: createdName,
		},
		Msg:     "Created user successfully",
		Success: true,
	}, nil
}

func (s *server) LoginUser(ctx context.Context, req *blogpb.LoginUserRequest) (*blogpb.LoginUserResponse, error) {

	email := req.GetEmail()

	row := s.getUserByEmail.QueryRowContext(ctx, email)

	var returnedID, returnedName, returnedEmail, returnedPW, returnedSalt string

	err := row.Scan(&returnedID, &returnedName, &returnedEmail, &returnedPW, &returnedSalt)
	if err == sql.ErrNoRows {
		return nil, status.Errorf(codes.InvalidArgument, "invalid credentials")
	} else if err != nil {
		log.Fatal(err)
	}

	if req.GetPassword() != returnedPW {
		return nil, status.Errorf(codes.InvalidArgument, "invalid credentials")
	}

	return &blogpb.LoginUserResponse{
		Msg:     "login successful",
		Success: "true",
	}, nil
}

func (s *server) GetUser(ctx context.Context, req *blogpb.GetUserRequest) (*blogpb.GetUserResponse, error) {
	id := req.GetUserId()

	row := s.getUserByID.QueryRowContext(ctx, id)

	var returnedID, returnedName, returnedEmail, returnedPW, returnedSalt string

	err := row.Scan(&returnedID, &returnedName, &returnedEmail, &returnedPW, &returnedSalt)
	if err != nil {
		log.Fatal(err)
	}

	return &blogpb.GetUserResponse{
		User: &blogpb.User{
			Id:   returnedID,
			Name: returnedName,
		},
		Msg:     "user retrieved successfully",
		Success: true,
	}, nil
}

func (s *server) UpdateUser(ctx context.Context, req *blogpb.UpdateUserRequest) (*blogpb.UpdateUserResponse, error) {
	id := req.GetUserId()
	newName := req.GetUser().GetName()

	row := s.updateUser.QueryRowContext(ctx, id, newName, "random@something.com")

	var returnedID, returnedName, returnedEmail string

	err := row.Scan(&returnedID, &returnedName, &returnedEmail)
	if err != nil {
		log.Fatal(err)
	}

	return &blogpb.UpdateUserResponse{
		User: &blogpb.User{
			Id:   returnedID,
			Name: returnedName,
		},
		Msg:     "updated user successfully",
		Success: true,
	}, nil
}

func (s *server) DeleteUser(ctx context.Context, req *blogpb.DeleteUserRequest) (*blogpb.DeleteUserResponse, error) {
	id := req.GetUserId()

	result, err := s.deleteUser.ExecContext(ctx, id)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()
	return &blogpb.DeleteUserResponse{
		Msg:     fmt.Sprintf("deleted %v user", rowsAffected),
		Success: true,
	}, nil
}
