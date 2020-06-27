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

	testCall(conn)
}

func testCall(conn *grpc.ClientConn) {
	c := blogpb.NewGreetingsClient(conn)

	res, err := c.SayHello(context.Background(), &blogpb.SayHelloRequest{})
	if err != nil {
		fmt.Printf("error getting response: %v\n", err)
	}

	fmt.Println(res.GetMessage().GetText())
}
