package controller

import (
	"net/http"
	"source-station/model"
	"source-station/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// post
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
	var post model.Post

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

// user
func (app *App) GetUsers(c *gin.Context) {
	users, err := app.DB.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (app *App) RegisterUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := app.DB.InsertUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	inserteID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error registering user, please try again"})
		return
	}

	token, err := utils.GenerateJWTToken(primitive.ObjectID(inserteID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error registering user, please try again"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully", "token": token})
}

func (app *App) LoginUser(c *gin.Context) {
	var loginData struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := app.DB.FindUser(loginData.Username)
	println(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// TODO: compare hash password
	if loginData.Password != user.Password {
		println("2")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := utils.GenerateJWTToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (app *App) UpdateUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tokenString := c.GetHeader("Authorization")
	userID, err := utils.GetClaimFromJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	result, err := app.DB.UpdateUser(userID, &user)
	println(result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "user updated successfully"})
}

// comment
func (app *App) AddComment(c *gin.Context) {
	var comment model.Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := app.DB.InsertComment(&comment)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id post ID"})
		return
	}

	comments, err := app.DB.GetPostComments(postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}
