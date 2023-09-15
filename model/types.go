package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Username  string             `bson:"username" json:"username" binding:"required"`
	Email     string             `bson:"email" json:"email" binding:"required"`
	Password  string             `bson:"password" json:"password" binding:"required"`
	Bio       string             `bson:"bio" json:"bio"`
	CreatedAt primitive.DateTime `bson:"createdAt,omitempty,default:currentTimestamp" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updatedAt,omitempty,default:currentTimestamp" json:"updated_at"`
}

type Post struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Title       string             `bson:"title" json:"title" binding:"required"`
	Content     string             `bson:"content" json:"content" binding:"required"`
	ContentType string             `bson:"contentType" json:"content_type" binding:"required"`
	User        primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	Visibility  string             `bson:"visibility" json:"visibility" binding:"required"`
	ViewCount   int                `bson:"viewCount" json:"view_count"`
	LikesCount  int                `bson:"likesCount" json:"likes_count"`
	CreatedAt   primitive.DateTime `bson:"createdAt,omitempty,default:currentTimestamp" json:"created_at"`
	UpdatedAt   primitive.DateTime `bson:"updatedAt,omitempty,default:currentTimestamp" json:"updated_at"`
}

type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Content   string             `bson:"content" json:"content" binding:"required"`
	User      primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	Post      primitive.ObjectID `bson:"post_id,omitempty" json:"post_id" binding:"required"`
	CreatedAt primitive.DateTime `bson:"createdAt,omitempty,default:currentTimestamp" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updatedAt,omitempty,default:currentTimestamp" json:"updated_at"`
}
