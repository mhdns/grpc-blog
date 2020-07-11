package main

import (
	"blog/server/blogpb"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
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

	query := readSQL("queries/create_blog.sql")
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
	// rows, err := s.db.Query("Select * from blog;")
	// if err != nil {
	// 	fmt.Println("Error ", err)
	// }

	// for rows.Next() {
	// 	var id, title, createdAt, post string
	// 	if err := rows.Scan(&id, &title, &createdAt, &post); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(id, title, createdAt, post)
	// }

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
	getID := req.GetBlogId()

	query := readSQL("queries/get_blog.sql")
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "database error: %v", err.Error())
	}
	defer stmt.Close()

	var id, title, createdAt, post string
	err = stmt.QueryRowContext(context.Background(), getID).Scan(&id, &title, &createdAt, &post)
	if err == sql.ErrNoRows {
		fmt.Println(err)
		return nil, status.Errorf(codes.NotFound, "blog with id %v not found: ", getID, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "database err: %v", err.Error())
	}

	return &blogpb.GetBlogResponse{
		Blog: &blogpb.Blog{
			Id:    id,
			Title: title,
			Date:  createdAt,
			Post:  post,
		},
		Msg:     "fetched one blog successfully",
		Success: true,
	}, nil
}

func (s *server) UpdateBlog(ctx context.Context, req *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {
	updateID := req.GetBlogId()
	updateTitle := req.GetBlog().GetTitle()
	updatePost := req.GetBlog().GetPost()

	query := readSQL("queries/update_blog.sql")
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "database error: %v", err.Error())
	}
	defer stmt.Close()

	var id, title, createdAt, post string
	err = stmt.QueryRowContext(context.Background(), updateID, updateTitle,
		updatePost).Scan(&id, &title, &createdAt, &post)
	if err == sql.ErrNoRows {
		fmt.Println(err)
		return nil, status.Errorf(codes.NotFound, "blog with id %v not found: ", updateID, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "database err: %v", err.Error())
	}

	return &blogpb.UpdateBlogResponse{
		Blog: &blogpb.Blog{
			Id:    id,
			Title: title,
			Date:  createdAt,
			Post:  post,
		},
		Msg:     "updated one blog successfully",
		Success: true,
	}, nil
}

func (s *server) DeleteBlog(ctx context.Context, req *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {
	deleteID := req.GetBlogId()

	query := readSQL("queries/delete_blog.sql")
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "database error: %v", err.Error())
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(context.Background(), deleteID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "database error: %v", err.Error())
	}
	val, _ := res.RowsAffected()

	return &blogpb.DeleteBlogResponse{
		Msg:     fmt.Sprintf("successfully deleted %v rows", val),
		Success: true,
	}, nil
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
