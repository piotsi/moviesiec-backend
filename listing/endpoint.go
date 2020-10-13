package listing

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/piotsik/moviesiec-backend/movies"

	"go.mongodb.org/mongo-driver/mongo"
)

type Handler func(http.ResponseWriter, *http.Request, httprouter.Params)

func MakeGetMoviesEndpoint(client *mongo.Client) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		moviesList := movies.GetAll(r, client)

		if moviesList == nil {
			http.Error(w, "No movies to display", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(moviesList)
	}
}

func MakeGetMovieEndpoint(client *mongo.Client) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		movie, err := movies.Get(p.ByName("longid"), client)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(movie)
	}
}

func MakeAddMovieEndpoint(client *mongo.Client) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		var movie movies.Movie

		err := json.NewDecoder(r.Body).Decode(&movie)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = movies.Add(client, movie)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
