package controller

import (
	"net/http"
	"source-station/database"
	"source-station/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *App) AddComment(c *gin.Context) {
	var comment database.Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tokenString := c.GetHeader("Authorization")
	userID, err := utils.GetClaimFromJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	result, err := app.DB.InsertComment(userID, &comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "comment created successfully", "id": result.InsertedID})
}

func (app *App) GetComments(c *gin.Context) {
	postIDParam := c.Param("postID")

	postID, err := primitive.ObjectIDFromHex(postIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	}

	comments, err := app.DB.GetPostComments(postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (app *App) UpdateComment(c *gin.Context) {
	var comment database.Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err := app.DB.UpdateComment(&comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "comment updateed"})
}

func (app *App) DeleteComment(c *gin.Context) {
	var comment database.Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := app.DB.DeleteComment(&comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.DeletedCount > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "document deleted successfully"})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "document not found"})
}
