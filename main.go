package main

import (
	"log"

	"source-station/controller"
	"source-station/utils"
)

var app *controller.App

func init() {
	var err error
	app, err = controller.NewApp("source-station", "3000")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	_, err := app.DB.InsertUser(utils.RandomUser())
	if err != nil {
		log.Fatal(err)
	}
	println("added user")

	_, err = app.DB.InsertPost(*utils.RandomPost())
	if err != nil {
		log.Fatal(err)
	}
	println("added post")

	app.InitializeRoutes()
	err = app.Router.Run(app.Port)
	if err != nil {
		log.Fatal(err)
	}
}
