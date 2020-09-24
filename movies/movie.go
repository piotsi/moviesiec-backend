package movies

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Movie struct {
	Title    string
	Director string
	Year     int
	Rating   float64
	Ratings  int
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
