package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username  string             `bson:"username" json:"userName"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	Bio       string             `bson:"bio" json:"bio"`
	CreatedAt primitive.DateTime `bson:"createdAt,omitempty,default:currentTimestamp" json:"createdAt"`
	UpdatedAt primitive.DateTime `bson:"updatedAt,omitempty,default:currentTimestamp" json:"updatedAt"`
}

type Post struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Content     string             `bson:"content" json:"content"`
	ContentType string             `bson:"contentType" json:"contentTypes"`
	User        primitive.ObjectID `bson:"user,omitempty" json:"user_id"`
	Visibility  string             `bson:"visibility" json:"visibility"`
	ViewCount   int                `bson:"viewCount" json:"viewCount"`
	LikesCount  int                `bson:"likesCount" json:"likesCount"`
	CreatedAt   primitive.DateTime `bson:"createdAt,omitempty,default:currentTimestamp" json:"createdAt"`
	UpdatedAt   primitive.DateTime `bson:"updatedAt,omitempty,default:currentTimestamp" json:"updatedAt"`
}

type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Content   string
	User      primitive.ObjectID `bson:"user,omitempty"`
	Post      primitive.ObjectID `bson:"post,omitempty"`
	CreatedAt primitive.DateTime `bson:"createdAt,omitempty,default:currentTimestamp"`
	UpdatedAt primitive.DateTime `bson:"updatedAt,omitempty,default:currentTimestamp"`
}
