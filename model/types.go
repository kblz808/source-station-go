package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	Bio       string             `bson:"bio"`
	CreatedAt primitive.DateTime `bson:"createdAt,omitempty"`
	UpdatedAt primitive.DateTime `bson:"updatedAt,omitemty"`
}

type Post struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Title          string             `bson:"title"`
	Content        string             `bson:"content"`
	ContentType    string             `bson:"contentType"`
	User           primitive.ObjectID `bson:"user,omitempty"`
	Visibility     string             `bson:"visibility"`
	ExpirationDate primitive.DateTime `bson:"expirationDate"`
	ViewCount      int                `bson:"viewCount"`
	LikesCount     int                `bson:"likesCount"`
	CreatedAt      primitive.DateTime `bson:"createdAt,omitempty"`
	UpdatedAt      primitive.DateTime `bson:"updatedAt,omitempty"`
}

type Comment struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Content string
	User    primitive.ObjectID `bson:"user,omitempty"`
	Post    primitive.ObjectID `bson:"post,omitempty"`
}
