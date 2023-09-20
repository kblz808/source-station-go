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

// user
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
