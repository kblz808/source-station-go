package controller

import (
	"net/http"
	"source-station/database"
	"source-station/utils"

	"github.com/gin-gonic/gin"
)

func (app *App) GetPosts(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	userID, err := utils.GetClaimFromJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	posts, err := app.DB.GetPosts(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (app *App) AddPost(c *gin.Context) {
	var post database.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tokenString := c.GetHeader("Authorization")
	userID, err := utils.GetClaimFromJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	result, err := app.DB.InsertPost(userID, &post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "post created successfully", "id": result.InsertedID})
}

func (app *App) UpdatePost(c *gin.Context) {
	var post database.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err := app.DB.UpdatePost(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "post updated successfully"})
}
