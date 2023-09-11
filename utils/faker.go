package utils

import (
	"source-station/model"
	"time"

	"github.com/brianvoe/gofakeit/v6"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RandomUser() *model.User {
	return &model.User{
		ID:        primitive.NewObjectIDFromTimestamp(time.Now()),
		Username:  gofakeit.Username(),
		Email:     gofakeit.Email(),
		Password:  gofakeit.Password(true, true, true, false, false, 10),
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Bio:       gofakeit.Hobby(),
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}
}
