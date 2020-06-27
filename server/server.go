package main

import (
	"blog/blogpb"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) SayHello(ctx context.Context, req *blogpb.SayHelloRequest) (*blogpb.SayHelloResponse, error) {
	return &blogpb.SayHelloResponse{
		Message: &blogpb.Hello{
			Text: "Hello World!",
		},
	}, nil
}

func main() {
	li, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalln("unable to create tcp connecion: ", err)
	}

	s := grpc.NewServer()
	blogpb.RegisterGreetingsServer(s, &server{})

	if err = s.Serve(li); err != nil {
		log.Fatalln("unable to create grpc server: ", err)
	}
}
