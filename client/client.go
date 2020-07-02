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
			Title: "My First Blog",
			Post:  "This is my first ever blog. I hope everyone likes it. If you get to read it, good for you",
		},
	}

	createBlog(conn, req)

	// testCall(conn)
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

// func testCall(conn *grpc.ClientConn) {
// 	c := blogpb.NewGreetingsClient(conn)

// 	res, err := c.SayHello(context.Background(), &blogpb.SayHelloRequest{})
// 	if err != nil {
// 		fmt.Printf("error getting response: %v\n", err)
// 	}

// 	fmt.Println(res.GetMessage().GetText())
// }
