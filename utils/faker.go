package utils

import (
	"source-station/model"
	"time"

	"github.com/brianvoe/gofakeit/v6"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RandomUser() *model.User {
	return &model.User{
		Username:  gofakeit.Username(),
		Email:     gofakeit.Email(),
		Password:  gofakeit.Password(true, true, true, false, false, 10),
		Bio:       gofakeit.Hobby(),
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}
}

func RandomPost() *model.Post {
	return &model.Post{
		Title:          gofakeit.BookTitle(),
		Content:        gofakeit.Quote(),
		ContentType:    "md",
		User:           primitive.NewObjectIDFromTimestamp(time.Now()),
		Visibility:     "public",
		ExpirationDate: primitive.NewDateTimeFromTime(time.Now()),
		ViewCount:      0,
		LikesCount:     0,
		CreatedAt:      primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:      primitive.NewDateTimeFromTime(time.Now()),
	}
}
