package movies

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Movie struct {
	Title    string  `json:"title"`
	Director string  `json:"director"`
	Year     int     `json:"year"`
	Rating   float64 `json:"rating"`
	Ratings  int     `json:"ratings"`
}

type Repository interface {
	GetAll() []Movie
	Get(int) (Movie, error)
	Add(Movie) error
}

func GetAll(client *mongo.Client) []*Movie {
	var moviesList []*Movie

	moviesCollection := client.Database("moviesiec").Collection("movies")

	findOptions := options.Find().SetLimit(2)

	cur, err := moviesCollection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem Movie
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		moviesList = append(moviesList, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	return moviesList
}

func Add(client *mongo.Client, movie Movie) {
	moviesCollection := client.Database("moviesiec").Collection("movies")

	fmt.Println(movie)

	_, err := moviesCollection.InsertOne(context.TODO(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Added new movie")
}
