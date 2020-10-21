package users

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Repository interface {
	SignIn() error
	SignOut() error
	SignOn() error
}

func SignOn(client *mongo.Client, user User) error {
	usersCollection := client.Database("moviesiec").Collection("users")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	_, err = usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("New user registered:", user.Username)

	return nil
}
