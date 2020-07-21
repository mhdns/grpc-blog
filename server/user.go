package main

import (
	"blog/server/blogpb"
	"context"
	"fmt"
	"log"
)

func (s *server) CreateUser(ctx context.Context, req *blogpb.CreateUserRequest) (*blogpb.CreateUserResponse, error) {
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

	return &blogpb.CreateUserResponse{
		User: &blogpb.User{
			Id:   createdID,
			Name: createdName,
		},
		Msg:     "Created user successfully",
		Success: true,
	}, nil
}

func (s *server) GetUser(ctx context.Context, req *blogpb.GetUserRequest) (*blogpb.GetUserResponse, error) {
	id := req.GetUserId()

	query := readSQL("queries/get_user_id.sql")
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	row := stmt.QueryRowContext(ctx, id)

	var returnedID, returnedName, returnedEmail, returnedPW, returnedSalt string

	err = row.Scan(&returnedID, &returnedName, &returnedEmail, &returnedPW, &returnedSalt)
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

	query := readSQL("queries/update_user.sql")
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	row := stmt.QueryRowContext(ctx, id, newName, "random@something.com")

	var returnedID, returnedName string

	err = row.Scan(&returnedID, &returnedName)
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
	query := readSQL("queries/delete_user.sql")

	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, id)

	rowsAffected, _ := result.RowsAffected()
	return &blogpb.DeleteUserResponse{
		Msg:     fmt.Sprintf("deleted %v user", rowsAffected),
		Success: true,
	}, nil
}
