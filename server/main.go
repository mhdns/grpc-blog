package main

import (
	"blog/blogpb"
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type blog struct {
	Title string
	Date  string
	Post  string
}

type server struct {
	client *mongo.Client
}

func (s *server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	collection := s.client.Database("test").Collection("blogs")
	currentTime := time.Now().Format("Mon Jan 2 15:04:05 MST 2006")
	input := blog{
		Title: req.GetBlog().Title,
		Date:  currentTime,
		Post:  req.GetBlog().GetPost(),
	}

	insertResult, err := collection.InsertOne(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	result := insertResult.InsertedID.(primitive.ObjectID).Hex()

	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:    result,
			Title: input.Title,
			Date:  input.Date,
			Post:  input.Post,
		},
		Msg:     "Blog created Successfully",
		Success: true,
	}, nil
}

func (s *server) GetBlog(ctx context.Context, req *blogpb.GetBlogRequest) (*blogpb.GetBlogResponse, error) {
	collection := s.client.Database("test").Collection("blogs")

	objID, err := primitive.ObjectIDFromHex(req.GetBlogId())
	if err != nil {
		fmt.Printf("invalid objectId: %v", err)
		return nil, err
	}

	result := new(blog)

	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(result)
	if err != nil {
		fmt.Printf("unable to get document: %v", err)
		return nil, err
	}

	return &blogpb.GetBlogResponse{
		Blog: &blogpb.Blog{
			Id:    req.GetBlogId(),
			Title: result.Title,
			Date:  result.Date,
			Post:  result.Post,
		},
	}, nil
}

func (s *server) UpdateBlog(ctx context.Context, req *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {
	collection := s.client.Database("test").Collection("blogs")

	objID, err := primitive.ObjectIDFromHex(req.GetBlogId())
	if err != nil {
		fmt.Printf("invalid objectId: %v", err)
		return nil, err
	}

	filter := bson.M{"_id": objID}

	update := bson.D{
		{"$set", bson.D{{"title", req.GetBlog().GetTitle},
			{"post", req.GetBlog().GetPost}},
		}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Printf("unable to update document: %v", err)
		return nil, err
	}

	fmt.Println(result.ModifiedCount)

	return &blogpb.UpdateBlogResponse{}, nil
}

func (s *server) DeleteBlog(ctx context.Context, req *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {
	collection := s.client.Database("test").Collection("blogs")

	objID, err := primitive.ObjectIDFromHex(req.GetBlogId())
	if err != nil {
		fmt.Printf("invalid objectId: %v", err)
		return nil, err
	}

	filter := bson.M{"_id": objID}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("unable to delete document: %v", err)
		return nil, err
	}

	fmt.Println(result.DeletedCount)
	return &blogpb.DeleteBlogResponse{}, nil
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
	defer client.Disconnect(context.TODO())

	err = s.Serve(li)
	if err != nil {
		log.Fatalf("unable to server grpc server: %v", err)
	}
}
