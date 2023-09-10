package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	FirstName string             `bson:"firstName"`
	LastName  string             `bson:"lastName"`
	Bio       string             `bson:"bio"`
	CreatedAt primitive.DateTime `bson:"createdAt,omitempty"`
	UpdatedAt primitive.DateTime `bson:"updatedAt,omitemty"`
}

func main() {
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOptions := options.Client().ApplyURI("mongodb+srv://harofpicvw:mongo-password@cluster0.wlhh0so.mongodb.net/?retryWrites=true&w=majority")

	// connect
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// check connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected")

	// users collection
	usersCollection := client.Database("mydb").Collection("users")

	// new user
	newUser := User{
		Username:  "admin",
		Email:     "admin@mail.com",
		Password:  "admin123",
		FirstName: "alex",
		LastName:  "morty",
		Bio:       "the admin",
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	// insert
	insertResult, err := usersCollection.InsertOne(context.Background(), newUser)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("inserted id: ", insertResult.InsertedID)

}
