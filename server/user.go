package main

import (
	"blog/server/userpb"
	"context"
	"fmt"
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
	id := req.GetUserId()

	query := readSQL("queries/get_user.sql")
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	row := stmt.QueryRowContext(ctx, id)

	var returnedID, returnedName, returnedEmail, returnedPW, returnedSalt string

	err = row.Scan(returnedID, returnedName, returnedEmail, returnedPW, returnedSalt)
	if err != nil {
		log.Fatal(err)
	}

	return &userpb.GetUserResponse{
		User: &userpb.User{
			Id:   returnedID,
			Name: returnedName,
		},
		Msg:     "user retrieved successfully",
		Success: true,
	}, nil
}

func (s *server) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	id := req.GetUserId()
	newName := req.GetUser().GetName()

	query := readSQL("queries/update_user.sql")
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	row := stmt.QueryRowContext(ctx, id, newName, "random@something.com")

	var returnedID, returnedName string

	err = row.Scan(returnedID, returnedName)
	if err != nil {
		log.Fatal(err)
	}

	return &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id:   returnedID,
			Name: returnedName,
		},
		Msg:     "updated user successfully",
		Success: true,
	}, nil
}

func (s *server) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	id := req.GetUserId()
	query := readSQL("query/delete_user.sql")

	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, id)

	rowsAffected, _ := result.RowsAffected()
	return &userpb.DeleteUserResponse{
		Msg:     fmt.Sprintf("deleted %v user", rowsAffected),
		Success: true,
	}, nil
}
