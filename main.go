package main

import (
	"context"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/piotsik/moviesiec/db"
	"github.com/piotsik/moviesiec/listing"
)

func main() {
	client := db.Init()
	defer client.Disconnect(context.TODO())

	router := httprouter.New()
	router.GET("/movies", listing.MakeGetMoviesEndpoint(client))

	log.Fatal(http.ListenAndServe(":8080", router))
}
