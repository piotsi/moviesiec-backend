package main

import (
	"context"
	"moviesiec/db"
)

func main() {
	client := db.Init()
	defer client.Disconnect(context.TODO())
}
