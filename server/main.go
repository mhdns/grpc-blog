package main

import (
	"blog/blogpb"
	"context"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type server struct {
	client *mongo.Client
}

func (*server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	return nil, nil
}

func (*server) GetBlog(ctx context.Context, req *blogpb.GetBlogRequest) (*blogpb.GetBlogResponse, error) {
	return nil, nil
}

func (*server) UpdateBlog(ctx context.Context, req *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {
	return nil, nil
}

func (*server) DeleteBlog(ctx context.Context, req *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {
	return nil, nil
}

func main() {
	client, err := dbConnect("mongodb+srv://su_mhdns:Password123@cluster0-dewl6.mongodb.net/test?retryWrites=true&w=majority")
	if err != nil {
		log.Fatalf("unable to connect to db: %v", err)
	}
	log.Println("Connected to DB!")

	li, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("unable to create listener: %v", err)
	}

	s := grpc.NewServer()
	blogpb.RegisterBlogServiceServer(s, &server{client: client})

	err = s.Serve(li)
	if err != nil {
		log.Fatalf("unable to server grpc server: %v", err)
	}
}
