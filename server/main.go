package main

import (
	"blog/server/blogpb"
	"context"
	"database/sql"
	"fmt"
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
	newTitle := req.GetBlog().GetTitle()
	newPost := req.GetBlog().GetPost()

	query := readSQL("create_blog.sql")
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(context.Background(), newTitle, newPost)

	var id, title, createdAt, post string

	err = row.Scan(&id, &title, &createdAt, &post)
	if err != nil {
		log.Fatal(err)
	}

	// Remove later
	rows, err := s.db.Query("Select * from blog;")
	if err != nil {
		fmt.Println("Error ", err)
	}

	for rows.Next() {
		var id, title, createdAt, post string
		if err := rows.Scan(&id, &title, &createdAt, &post); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, title, createdAt, post)
	}

	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:    id,
			Title: title,
			Date:  createdAt,
			Post:  post,
		},
		Msg:     "blog created successfully",
		Success: true,
	}, nil
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

	createBlogTable := readSQL("create_table.sql")
	_, err = db.Exec(createBlogTable)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("table created")

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
