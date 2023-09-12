package main

import (
	"log"
	"source-station/controller"

	"github.com/joho/godotenv"
)

var app *controller.App

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	app, err = controller.NewApp("source-station", ":3000", "debug")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	err := app.Router.Run(app.Port)
	if err != nil {
		log.Fatal(err)
	}
}
