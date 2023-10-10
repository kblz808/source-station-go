package controller

import (
	"net/http"
	"source-station/database"
	"source-station/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *App) GetUsers(c *gin.Context) {
	users, err := app.DB.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (app *App) RegisterUser(c *gin.Context) {
	var user database.User

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
	var user database.User

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

	_, err = app.DB.UpdateUser(userID, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "user updated successfully"})
}
