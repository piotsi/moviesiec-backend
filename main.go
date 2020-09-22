package main

import (
	"context"

	"github.com/piotsik/moviesiec/db"
)

func main() {
	client := db.Init()
	defer client.Disconnect(context.TODO())
}
