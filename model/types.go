package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username  string             `bson:"username" json:"username"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	Bio       string             `bson:"bio" json:"bio"`
	CreatedAt primitive.DateTime `bson:"createdAt,omitempty,default:currentTimestamp" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updatedAt,omitempty,default:currentTimestamp" json:"updated_at"`
}

type Post struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Content     string             `bson:"content" json:"content"`
	ContentType string             `bson:"contentType" json:"content_type"`
	User        primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	Visibility  string             `bson:"visibility" json:"visibility"`
	ViewCount   int                `bson:"viewCount" json:"view_count"`
	LikesCount  int                `bson:"likesCount" json:"likes_count"`
	CreatedAt   primitive.DateTime `bson:"createdAt,omitempty,default:currentTimestamp" json:"created_at"`
	UpdatedAt   primitive.DateTime `bson:"updatedAt,omitempty,default:currentTimestamp" json:"updated_at"`
}

type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Content   string             `bson:"content" json:"content"`
	User      primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	Post      primitive.ObjectID `bson:"post_id,omitempty" json:"post_id"`
	CreatedAt primitive.DateTime `bson:"createdAt,omitempty,default:currentTimestamp" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updatedAt,omitempty,default:currentTimestamp" json:"updated_at"`
}
