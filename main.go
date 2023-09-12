package main

import (
	"log"

	"source-station/controller"
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
	app.InitializeRoutes()
	err := app.Router.Run(app.Port)
	if err != nil {
		log.Fatal(err)
	}
}
