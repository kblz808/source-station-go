package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"source-station/controllers"
	"source-station/model"
)

var app *controllers.App

func init() {
	app = &controllers.App{
		Name: "source-station",
	}
	if err := app.DB.Connect(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	newUser := model.User{
		Username:  "admin",
		Email:     "admin@mail.com",
		Password:  "admin123",
		FirstName: "alex",
		LastName:  "morty",
		Bio:       "the admin",
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	_, err := app.DB.InsertUser(newUser)
	if err != nil {
		log.Fatal(err)
	}

	println("added user")

	newPost := model.Post{
		Title:          "first post",
		Content:        "console.log(code)",
		User:           primitive.NewObjectID(),
		Visibility:     "public",
		ExpirationDate: primitive.NewDateTimeFromTime(time.Now()),
		ViewCount:      0,
		LikesCount:     0,
		CreatedAt:      primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:      primitive.NewDateTimeFromTime(time.Now()),
	}

	err = app.DB.InsertPost(newPost)
	if err != nil {
		log.Fatal(err)
	}

	println("new post added")

	router := gin.Default()
	router.GET("/posts", app.GetPosts)
	router.GET("/users", app.GetUsers)

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
