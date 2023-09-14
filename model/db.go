package model

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func (db *DB) Connect() error {
	clientOptions := options.Client().ApplyURI(os.Getenv("CONNECTION_STRING"))

	var err error
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
	collection := db.client.Database("mydb").Collection("users")
	result, err := collection.InsertOne(context.Background(), newUser)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (db *DB) FindUser(username string) (*User, error) {
	collection := db.client.Database("mydb").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User

	result := collection.FindOne(ctx, bson.M{"username": username})
	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *DB) UpdateUser(userIDString string, updatedUser *User) (*mongo.UpdateResult, error) {
	collection := db.client.Database("mydb").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userID, err := primitive.ObjectIDFromHex(userIDString)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": userID}
	update := bson.M{
		"$set": bson.M{
			"username": updatedUser.Username,
			"email":    updatedUser.Email,
			"password": updatedUser.Password,
			"bio":      updatedUser.Bio,
		},
	}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (db *DB) InsertPost(newPost *Post) (*mongo.InsertOneResult, error) {
	collecton := db.client.Database("mydb").Collection("posts")
	result, err := collecton.InsertOne(context.Background(), newPost)
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

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (db *DB) InsertComment(newComment *Comment) (*mongo.InsertOneResult, error) {
	collection := db.client.Database("mydb").Collection("comments")
	result, err := collection.InsertOne(context.Background(), newComment)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (db *DB) GetPostComments(postID primitive.ObjectID) ([]Comment, error) {
	collection := db.client.Database("mydb").Collection("comments")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pipeline := bson.A{
		bson.M{
			"$match": bson.M{"post_id": postID},
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var comments []Comment
	for cursor.Next(ctx) {
		var comment Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
