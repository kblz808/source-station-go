package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"source-station/controllers"
	"source-station/utils"
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
	_, err := app.DB.InsertUser(utils.RandomUser())
	if err != nil {
		log.Fatal(err)
	}
	println("added user")

	router := gin.Default()
	router.GET("/posts", app.GetPosts)
	router.GET("/users", app.GetUsers)

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
