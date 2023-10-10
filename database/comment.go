package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *DB) InsertComment(userIDString string, newComment *Comment) (*mongo.InsertOneResult, error) {
	collection := db.client.Database("mydb").Collection("comments")

	userID, err := primitive.ObjectIDFromHex(userIDString)
	if err != nil {
		return nil, err
	}
	newComment.User = userID

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

func (db *DB) UpdateComment(newComment *Comment) (*mongo.UpdateResult, error) {
	collection := db.client.Database("mydb").Collection("comments")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": newComment.ID}
	update := bson.M{
		"$set": bson.M{
			"content": newComment.Content,
		},
	}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (db *DB) DeleteComment(comment *Comment) (*mongo.DeleteResult, error) {
	collection := db.client.Database("mydb").Collection("comments")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": comment.ID}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}
