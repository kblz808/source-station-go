package utils

import (
	"source-station/database"
	"time"

	"github.com/brianvoe/gofakeit/v6"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RandomUser() *database.User {
	return &database.User{
		Username:  gofakeit.Username(),
		Email:     gofakeit.Email(),
		Password:  gofakeit.Password(true, true, true, false, false, 10),
		Bio:       gofakeit.Hobby(),
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}
}

func RandomPost() *database.Post {
	return &database.Post{
		Title:       gofakeit.BookTitle(),
		Content:     gofakeit.Quote(),
		ContentType: "md",
		User:        primitive.NewObjectIDFromTimestamp(time.Now()),
		Visibility:  "public",
		ViewCount:   0,
		LikesCount:  0,
		CreatedAt:   primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:   primitive.NewDateTimeFromTime(time.Now()),
	}
}
