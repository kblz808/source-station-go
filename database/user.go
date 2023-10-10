package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

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
