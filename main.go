package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"source-station/model"
)

func init() {
	if err := model.Connect(); err != nil {
		log.Fatal(err)
	}
	println("connected")
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

	_, err := model.InsertUser(newUser)
	if err != nil {
		log.Fatal(err)
	}

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

	err = model.InsertPost(newPost)
	if err != nil {
		log.Fatal(err)
	}

	println("new post added")

	router := gin.Default()
	router.GET("/posts", func(c *gin.Context) {
		posts, err := model.GetAllPosts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, posts)
	})

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
