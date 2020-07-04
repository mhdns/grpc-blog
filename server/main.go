package main

import (
	"blog/server/blogpb"
	"context"
	"database/sql"
	"log"
	"net"

	"google.golang.org/grpc"
)

type blog struct {
	Title string
	Date  string
	Post  string
}

type server struct {
	db *sql.DB
}

func (s *server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {

	return &blogpb.CreateBlogResponse{}, nil
}

func (s *server) GetBlog(ctx context.Context, req *blogpb.GetBlogRequest) (*blogpb.GetBlogResponse, error) {

	return &blogpb.GetBlogResponse{}, nil
}

func (s *server) UpdateBlog(ctx context.Context, req *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {

	return &blogpb.UpdateBlogResponse{}, nil
}

func (s *server) DeleteBlog(ctx context.Context, req *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {

	return &blogpb.DeleteBlogResponse{}, nil
}

func main() {
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

	li, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("unable to create listener: %v", err)
	}

	s := grpc.NewServer()
	blogpb.RegisterBlogServiceServer(s, &server{db: db})
	defer db.Close()
	err = s.Serve(li)
	if err != nil {
		log.Fatalf("unable to server grpc server: %v", err)
	}
}
