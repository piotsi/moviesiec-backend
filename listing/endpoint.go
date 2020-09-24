package listing

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/piotsik/moviesiec/movies"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler func(http.ResponseWriter, *http.Request, httprouter.Params)

func MakeGetMoviesEndpoint(client *mongo.Client) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		moviesList := movies.GetAll(client)

		json.NewEncoder(w).Encode(moviesList)
	}
}
