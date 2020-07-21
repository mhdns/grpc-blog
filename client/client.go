package main

import (
	"blog/blogpb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds, _ := credentials.NewClientTLSFromFile("ca.crt", "")
	conn, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln("unable to dial grpc server: ", err)
	}
	defer conn.Close()

	// User portion
	createUser(conn, &blogpb.CreateUserRequest{
		Email:    "abc@gmail.com",
		Name:     "Ang Beng Chong",
		Password: "123456",
	})

	getUser(conn, &blogpb.GetUserRequest{
		UserId: "1",
	})

	updateUser(conn, &blogpb.UpdateUserRequest{
		UserId: "1",
		User: &blogpb.User{
			Id:   "1",
			Name: "Ang Beng Cheng",
		},
	})

	deleteUser(conn, &blogpb.DeleteUserRequest{
		UserId: "1",
	})
	// Blog portion
	// req := &blogpb.CreateBlogRequest{
	// 	Blog: &blogpb.Blog{
	// 		Title: "First Blog - Anas",
	// 		Post:  "This is my first blog. I hope everyone likes it. If you get to read it, good for you!!!!!",
	// 	},
	// }

	// createBlog(conn, req)
	// createBlog(conn, req)
	// createBlog(conn, req)

	// getBlog(conn, &blogpb.GetBlogRequest{BlogId: "1"})

	// updateBlog(conn, &blogpb.UpdateBlogRequest{
	// 	BlogId: "1",
	// 	Blog: &blogpb.Blog{
	// 		Title: "Some New Title",
	// 		Post:  "This is a new post",
	// 	},
	// })

	// deleteBlog(conn, &blogpb.DeleteBlogRequest{BlogId: "1"})
	// deleteBlog(conn, &blogpb.DeleteBlogRequest{BlogId: "1"})

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
	fmt.Println(res, err)
	if err != nil {
		return err
	}
	return nil
}

func updateBlog(conn *grpc.ClientConn, req *blogpb.UpdateBlogRequest) error {
	c := blogpb.NewBlogServiceClient(conn)

	res, err := c.UpdateBlog(context.Background(), req)
	fmt.Println(res, err)
	if err != nil {
		return err
	}

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

func createUser(conn *grpc.ClientConn, req *blogpb.CreateUserRequest) error {
	c := blogpb.NewUserServiceClient(conn)

	res, err := c.CreateUser(context.Background(), req)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func getUser(conn *grpc.ClientConn, req *blogpb.GetUserRequest) error {
	c := blogpb.NewUserServiceClient(conn)

	res, err := c.GetUser(context.Background(), req)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func updateUser(conn *grpc.ClientConn, req *blogpb.UpdateUserRequest) error {
	c := blogpb.NewUserServiceClient(conn)

	res, err := c.UpdateUser(context.Background(), req)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func deleteUser(conn *grpc.ClientConn, req *blogpb.DeleteUserRequest) error {
	c := blogpb.NewUserServiceClient(conn)

	res, err := c.DeleteUser(context.Background(), req)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
