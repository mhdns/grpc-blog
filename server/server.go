package main

import (
	"blog/blogpb"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type server struct{}

func (*server) SayHello(ctx context.Context, req *blogpb.SayHelloRequest) (*blogpb.SayHelloResponse, error) {
	return &blogpb.SayHelloResponse{
		Message: &blogpb.Hello{
			Text: "Hello World!",
		},
	}, nil
}

type trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	// li, err := net.Listen("tcp", ":5000")
	// if err != nil {
	// 	log.Fatalln("unable to create tcp connecion: ", err)
	// }

	// s := grpc.NewServer()
	// blogpb.RegisterGreetingsServer(s, &server{})

	// if err = s.Serve(li); err != nil {
	// 	log.Fatalln("unable to create grpc server: ", err)
	// }

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://su_mhdns:Password123@cluster0-dewl6.mongodb.net/test?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("test").Collection("trainers")

	ash := trainer{"Ash", 10, "Pallet Town"}
	misty := trainer{"Misty", 10, "Cerulean City"}
	brock := trainer{"Brock", 15, "Pewter City"}

	insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	trainers := []interface{}{misty, brock}

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	filter := bson.D{{"name", "Ash"}}

	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	objid, _ := primitive.ObjectIDFromHex("5ef9aa839ac8d1911ac5286a")
	fmt.Println(".......", objid.Hex())
	filter = bson.D{{"_id", objid}}
	// filter = bson.D{{"name", "Ash"}}

	var result trainer

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

}
