package movies

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Movie struct {
	StringID string  `json:"stringid"`
	Title    string  `json:"title"`
	Director string  `json:"director"`
	Year     int     `json:"year"`
	Rating   float64 `json:"rating"`
	Ratings  int     `json:"ratings"`
}

type Repository interface {
	GetAll() []Movie
	Get(string) (Movie, error)
	Add(Movie) error
}

func GetAll(page int64, client *mongo.Client) []*Movie {
	var moviesList []*Movie
	var moviesOnPage int64 = 5

	moviesCollection := client.Database("moviesiec").Collection("movies")

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"title", 1}})
	findOptions.SetLimit(moviesOnPage)
	findOptions.SetSkip(page * moviesOnPage)

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

func Get(stringID string, client *mongo.Client) (Movie, error) {
	var movie Movie

	moviesCollection := client.Database("moviesiec").Collection("movies")

	err := moviesCollection.FindOne(context.TODO(), bson.D{{"stringid", stringID}}).Decode(&movie)
	if err != nil {
		return movie, err
	}

	return movie, nil
}

func Add(client *mongo.Client, movie Movie) error {
	moviesCollection := client.Database("moviesiec").Collection("movies")

	movie.StringID = GenerateStringID(movie.Title, movie.Year)

	_, err := moviesCollection.InsertOne(context.TODO(), movie)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Added new movie", movie)

	return nil
}

func GenerateStringID(title string, year int) string {
	stringID := strings.Replace(strings.ToLower(title), " ", "-", -1)
	stringYear := strconv.Itoa(year)
	stringID = stringID + "-" + stringYear

	return stringID
}
