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

	// req := &blogpb.CreateBlogRequest{
	// 	Blog: &blogpb.Blog{
	// 		Title: "My First Blog",
	// 		Post:  "This is my first ever blog. I hope everyone likes it. If you get to read it, good for you",
	// 	},
	// }

	// createBlog(conn, req)

	getBlog(conn, &blogpb.GetBlogRequest{
		BlogId: "5efc038b73cd517b65e735ed",
	})

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
