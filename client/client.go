package main

import (
	"blog/blogpb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("unable to dial grpc server: ", err)
	}
	defer conn.Close()

	req := &blogpb.CreateBlogRequest{
		Blog: &blogpb.Blog{
			Title: "First Blog - Anas",
			Post:  "This is my first blog. I hope everyone likes it. If you get to read it, good for you!!!!!",
		},
	}

	createBlog(conn, req)
	createBlog(conn, req)
	createBlog(conn, req)
	createBlog(conn, req)
	createBlog(conn, req)

	// getBlog(conn, &blogpb.GetBlogRequest{
	// 	BlogId: "5efc038b73cd517b65e735ed",
	// })

	//updateBlog(conn, req)

	// deleteBlog(conn, &blogpb.DeleteBlogRequest{BlogId: "5efc038b73cd517b65e735ee"})

}

func createBlog(conn *grpc.ClientConn, req *blogpb.CreateBlogRequest) error {
	c := blogpb.NewBlogServiceClient(conn)

	res, err := c.CreateBlog(context.Background(), req)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func getBlog(conn *grpc.ClientConn, req *blogpb.GetBlogRequest) error {
	c := blogpb.NewBlogServiceClient(conn)

	res, err := c.GetBlog(context.Background(), req)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func updateBlog(conn *grpc.ClientConn, req *blogpb.UpdateBlogRequest) error {
	c := blogpb.NewBlogServiceClient(conn)

	res, err := c.UpdateBlog(context.Background(), req)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func deleteBlog(conn *grpc.ClientConn, req *blogpb.DeleteBlogRequest) error {
	c := blogpb.NewBlogServiceClient(conn)

	res, err := c.DeleteBlog(context.Background(), req)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
