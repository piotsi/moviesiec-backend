package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	password, ok := os.LookupEnv("MONGO_PASSWORD")
	if !ok {
		log.Fatal("Environmental variable for password not set")
	}
	uri := fmt.Sprintf("mongodb+srv://user:%s@cluster0.hnspb.gcp.mongodb.net/moviesiec?retryWrites=true&w=majority", password)

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database!")

	return client
}
