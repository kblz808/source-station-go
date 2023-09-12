package controller

import (
	"net/http"
	"source-station/model"

	"github.com/gin-gonic/gin"
)

func (app *App) GetPosts(c *gin.Context) {
	posts, err := app.DB.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, posts)
}

func (app *App) GetUsers(c *gin.Context) {
	users, err := app.DB.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, users)
}

func (app *App) AddUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	result, err := app.DB.InsertUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"ok": result.InsertedID})
}
