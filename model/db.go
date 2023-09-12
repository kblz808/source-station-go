package model

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func (db *DB) Connect() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	clientOptions := options.Client().ApplyURI(os.Getenv("CONNECTION_STRING"))

	db.client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) Ping() error {
	err := db.client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) InsertUser(newUser *User) (*mongo.InsertOneResult, error) {
	usersCollection := db.client.Database("mydb").Collection("users")
	result, err := usersCollection.InsertOne(context.Background(), newUser)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (db *DB) InsertPost(newPost Post) (*mongo.InsertOneResult, error) {
	postsCollection := db.client.Database("mydb").Collection("posts")
	result, err := postsCollection.InsertOne(context.Background(), newPost)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (db *DB) GetAllPosts() ([]Post, error) {
	collection := db.client.Database("mydb").Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{}
	options := options.Find().SetSort(bson.D{{Key: "time", Value: -1}}).SetLimit(10)

	cursor, err := collection.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var posts []Post
	for cursor.Next(ctx) {
		var post Post
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (db *DB) GetAllUsers() ([]User, error) {
	collection := db.client.Database("mydb").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []User
	for cursor.Next(ctx) {
		var user User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
