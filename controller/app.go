package controller

import (
	"source-station/model"

	"github.com/gin-gonic/gin"
)

type App struct {
	Name    string
	Port    string
	Version string
	DB      model.DB
	Router  *gin.Engine
}

func NewApp(name string, port string) (*App, error) {
	app := &App{
		Name:    "source-station",
		Port:    ":3000",
		Version: "v1",
		Router:  gin.Default(),
	}
	if err := app.DB.Connect(); err != nil {
		return nil, err
	}
	return app, nil
}

func (app *App) InitializeRoutes() {
	app.Router.GET("/posts", app.GetPosts)
	app.Router.GET("/users", app.GetUsers)
	app.Router.POST("/users", app.AddUser)
	app.Router.POST("/posts", app.AddPost)
}
