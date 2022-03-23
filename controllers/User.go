package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishnuhaasan/PushNotifications/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(c *gin.Context) {
	user := &models.User{
		ID:               primitive.NewObjectID(),
		ServiceProviders: []models.ServiceProvider{},
	}
	if err := c.BindJSON(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg":    "Cannot Create user, Check request body",
			"result": false,
			"error":  err,
		})
		panic(err)
	}
	_, err := user.CreateUser()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg":    "Unable to create user",
			"result": false,
			"error":  err,
		})
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":    "Successfully created user",
		"result": true,
		"data":   user,
	})
}

func GetUser(c *gin.Context) {
	aadhaar := c.Param("aadhaar")
	user, err := models.GetUser(aadhaar)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"msg":    "No User found",
			"result": false,
		})
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":    "User found",
		"result": true,
		"data":   user,
	})
}

func AddServiceProvider(c *gin.Context) {
	sp := &models.ServiceProvider{
		ID:       primitive.NewObjectID(),
		IsActive: false,
	}
	if err := c.BindJSON(sp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg":    "Cannot Create Service Provider, Check request body",
			"result": false,
			"error":  err,
		})
		panic(err)
	}
	aadhaar := c.Param("aadhaar")
	user, err := models.AddServiceProvider(sp, aadhaar)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": false,
			"msg":    "No user found with aadhaar",
		})
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":    "Successfully added service provider",
		"result": true,
		"data":   user,
	})
}

func ToggleIsActiveForProvider(c *gin.Context) {
	sp := &models.ServiceProvider{}
	if err := c.BindJSON(sp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg":    "Cannot Create Service Provider, Check request body",
			"result": false,
			"error":  err,
		})
		panic(err)
	}
	aadhaar := c.Param("aadhaar")
	user, err := models.ToggleIsActiveForProvider(aadhaar, sp.ID, sp.IsActive)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"msg":    "No Service provider or user found",
			"result": false,
		})
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":    "Successfully toggled is active of sp",
		"result": true,
		"data":   user,
	})
}
