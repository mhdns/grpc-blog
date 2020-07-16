package main

import (
	"blog/server/blogpb"
	"database/sql"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type blog struct {
	Title string
	Date  string
	Post  string
}

type server struct {
	db *sql.DB
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

	createBlogTable := readSQL("queries/create_table.sql")
	_, err = db.Exec(createBlogTable)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("table created")

	// gRPC Server

	creds, err := credentials.NewServerTLSFromFile("server.crt", "server.pem")
	if err != nil {
		log.Fatalf("unable to create credentials: %v", err)
	}

	li, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("unable to create listener: %v", err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	blogpb.RegisterBlogServiceServer(s, &server{db: db})
	defer db.Close()
	err = s.Serve(li)
	if err != nil {
		log.Fatalf("unable to server grpc server: %v", err)
	}
}
