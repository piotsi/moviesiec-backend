package login

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/piotsik/moviesiec-backend/users"
	"go.mongodb.org/mongo-driver/mongo"
)

func MakeSignOnEndpoint(client *mongo.Client) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		var user users.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = users.SignOn(client, user)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
