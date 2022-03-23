package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vishnuhaasan/PushNotifications/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateRequest(c *gin.Context) {
	newRequest := &models.Request{
		ID:         primitive.NewObjectID(),
		RaisedTime: int(time.Now().Unix()),
		IsDone:     false,
	}
	if err := c.BindJSON(newRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": false,
			"msg":    "Check Body",
			"error":  err,
		})
		panic(err)
	}
	fmt.Println(newRequest)
	_, err := newRequest.CreateRequest()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg":    "Unable to create request",
			"result": false,
			"error":  err,
		})
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   newRequest,
	})
}

func UpdateRequest(c *gin.Context) {
	aadhaar := c.Param("aadhaar")
	res, err := models.UpdateRequest(aadhaar)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": false,
			"error":  err,
		})
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   res,
	})
}

func GetRequests(c *gin.Context) {
	res, err := models.GetRequests()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": false,
			"error":  err,
		})
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   res,
	})
}
