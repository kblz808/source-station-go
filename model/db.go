package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Connect() error {
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOptions := options.Client().ApplyURI("mongodb+srv://harofpicvw:mongo-password@cluster0.wlhh0so.mongodb.net/?retryWrites=true&w=majority")

	var err error

	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	return nil
}

func Ping() error {
	err := client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}
	return nil
}

func InsertUser(newUser User) (*mongo.InsertOneResult, error) {
	usersCollection := client.Database("mydb").Collection("users")
	result, err := usersCollection.InsertOne(context.Background(), newUser)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func InsertPost(newPost Post) error {
	postsCollection := client.Database("mydb").Collection("posts")
	_, err := postsCollection.InsertOne(context.Background(), newPost)
	if err != nil {
		return err
	}
	return nil
}

func GetAllPosts() ([]Post, error) {
	collection := client.Database("mydb").Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		println("1-> ", err.Error())
		return nil, err
	}
	defer cursor.Close(ctx)

	var posts []Post
	for cursor.Next(ctx) {
		var post Post
		err := cursor.Decode(&post)
		if err != nil {
			println("2-> ", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
