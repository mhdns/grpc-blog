package main

import (
	"blog/server/blogpb"
	"database/sql"
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type blog struct {
	Title string
	Date  string
	Post  string
}

type server struct {
	createBlog, createUser, getBlog, getUserByID, getUserByEmail, updateBlog, updateUser, deleteBlog, deleteUser *sql.Stmt
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("--> unary interceptor", info.FullMethod)
	return handler(ctx, req)
}

func main() {
	// Database Initialization
	dbCred := dbCredentials{
		host:     "postgres",
		port:     5432,
		user:     "postgres",
		password: "example",
		dbname:   "postgres",
	}
	db, err := dbConnect(dbCred)
	if err != nil {
		panic(err)
	}
	log.Println("Connected to DB!")

	// Create Blog and User tables
	ctx := context.Background()

	_, err = db.ExecContext(ctx, createBlogTable)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("blog table created...")

	_, err = db.ExecContext(ctx, createUserTable)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("user table created...")

	// Create STMTs
	createBlogStmt, err := db.Prepare(createBlog)
	if err != nil {
		fmt.Println(err)
	}
	defer createBlogStmt.Close()

	createUserStmt, err := db.Prepare(createUser)
	if err != nil {
		fmt.Println(err)
	}
	defer createUserStmt.Close()

	getBlogStmt, err := db.Prepare(getBlog)
	if err != nil {
		fmt.Println(err)
	}
	defer getBlogStmt.Close()

	getUserByIDStmt, err := db.Prepare(getUserByID)
	if err != nil {
		fmt.Println(err)
	}
	defer getUserByIDStmt.Close()

	getUserByEmailStmt, err := db.Prepare(getUserByEmail)
	if err != nil {
		fmt.Println(err)
	}
	defer getUserByEmailStmt.Close()

	updateBlogStmt, err := db.Prepare(updateBlog)
	if err != nil {
		fmt.Println(err)
	}
	defer updateBlogStmt.Close()

	updateUserStmt, err := db.Prepare(updateUser)
	if err != nil {
		fmt.Println(err)
	}
	defer updateUserStmt.Close()

	deleteBlogStmt, err := db.Prepare(deleteBlog)
	if err != nil {
		fmt.Println(err)
	}
	defer deleteBlogStmt.Close()

	deleteUserStmt, err := db.Prepare(deleteUser)
	if err != nil {
		fmt.Println(err)
	}
	defer deleteUserStmt.Close()

	// gRPC Server

	creds, err := credentials.NewServerTLSFromFile("server.crt", "server.pem")
	if err != nil {
		log.Fatalf("unable to create credentials: %v", err)
	}

	li, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("unable to create listener: %v", err)
	}

	s := grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(unaryInterceptor),
	)
	blogServer := &server{
		createBlog:     createBlogStmt,
		createUser:     createUserStmt,
		getBlog:        getBlogStmt,
		getUserByID:    getUserByIDStmt,
		getUserByEmail: getUserByEmailStmt,
		updateBlog:     updateBlogStmt,
		updateUser:     updateUserStmt,
		deleteBlog:     deleteBlogStmt,
		deleteUser:     deleteUserStmt,
	}
	blogpb.RegisterBlogServiceServer(s, blogServer)
	blogpb.RegisterUserServiceServer(s, blogServer)
	defer db.Close()
	err = s.Serve(li)
	if err != nil {
		log.Fatalf("unable to server grpc server: %v", err)
	}
}
