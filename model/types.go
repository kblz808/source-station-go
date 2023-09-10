package model

import "go.mongodb.org/mongo-driver/bson/primitive"

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

type Post struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Title          string             `bson:"title"`
	Content        string             `bson:"content"`
	User           primitive.ObjectID `bson:"user,omitempty"`
	Visibility     string             `bson:"visibility"`
	ExpirationDate primitive.DateTime `bson:"expirationDate"`
	ViewCount      int                `bson:"viewCount"`
	LikesCount     int                `bson:"likesCount"`
	CreatedAt      primitive.DateTime `bson:"createdAt,omitempty"`
	UpdatedAt      primitive.DateTime `bson:"updatedAt,omitempty"`
}
