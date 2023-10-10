package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// post
func (db *DB) InsertPost(userIDString string, newPost *Post) (*mongo.InsertOneResult, error) {
	collecton := db.client.Database("mydb").Collection("posts")

	userID, err := primitive.ObjectIDFromHex(userIDString)
	if err != nil {
		return nil, err
	}
	newPost.User = userID

	result, err := collecton.InsertOne(context.Background(), newPost)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (db *DB) GetPosts(userIDString string) ([]Post, error) {
	collection := db.client.Database("mydb").Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userID, err := primitive.ObjectIDFromHex(userIDString)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"user_id": userID}

	cursor, err := collection.Find(ctx, filter)
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

func (db *DB) UpdatePost(newPost *Post) (*mongo.UpdateResult, error) {
	collection := db.client.Database("mydb").Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": newPost.ID}
	update := bson.M{
		"$set": bson.M{
			"title":       newPost.Title,
			"content":     newPost.Content,
			"contentType": newPost.ContentType,
			"visibility":  newPost.Visibility,
		},
	}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}
