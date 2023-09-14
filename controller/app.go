package controller

import (
	"source-station/model"

	"github.com/gin-gonic/gin"
)

type App struct {
	Name        string
	Port        string
	Version     string
	Environment string
	DB          model.DB
	Router      *gin.Engine
}

func NewApp(name, port, environment string) (*App, error) {
	if environment == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	app := &App{
		Name:        name,
		Port:        port,
		Version:     "v1",
		Environment: environment,
		Router:      gin.Default(),
	}

	if err := app.DB.Connect(); err != nil {
		return nil, err
	}

	app.InitializeRoutes()

	return app, nil
}

func (app *App) InitializeRoutes() {
	// app.Router.GET("/users", app.GetUsers)
	app.Router.POST("/register", app.RegisterUser)
	app.Router.POST("/login", app.LoginUser)

	app.Router.GET("/posts", app.JWTMiddleware, app.GetPosts)
	app.Router.POST("/posts", app.JWTMiddleware, app.AddPost)

	app.Router.POST("/comments/:postID", app.JWTMiddleware, app.GetComments)
	app.Router.POST("/comments", app.JWTMiddleware, app.AddComment)
}
