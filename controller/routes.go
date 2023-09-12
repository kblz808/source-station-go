package controller

import (
	"net/http"
	"source-station/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully", "id": result.InsertedID})
}

func (app *App) AddPost(c *gin.Context) {
	var post model.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	result, err := app.DB.InsertPost(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "post created successfully", "id": result.InsertedID})
}

func (app *App) AddComment(c *gin.Context) {
	var comment model.Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	result, err := app.DB.InsertComment(&comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "comment created successfully", "id": result.InsertedID})
}

func (app *App) GetComments(c *gin.Context) {
	postIDParam := c.Param("postID")

	postID, err := primitive.ObjectIDFromHex(postIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id post ID"})
		return
	}

	comments, err := app.DB.GetPostComments(postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, comments)
}
