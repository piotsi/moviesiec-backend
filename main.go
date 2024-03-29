package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"github.com/rs/cors"

	"github.com/julienschmidt/httprouter"
	"github.com/piotsik/moviesiec-backend/db"
	"github.com/piotsik/moviesiec-backend/listing"
	"github.com/piotsik/moviesiec-backend/login"
)

func main() {
	client := db.Init()
	defer client.Disconnect(context.TODO())

	router := httprouter.New()
	router.GET("/movies", listing.MakeGetMoviesEndpoint(client))
	router.GET("/movies/:longid", listing.MakeGetMovieEndpoint(client))
	router.POST("/movies/add", listing.MakeAddMovieEndpoint(client))
	router.POST("/user/register", login.MakeSignOnEndpoint(client))
	router.GET("/hello", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w, "hello, %s!\n", r.URL.Query().Get("name"))
	})

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
