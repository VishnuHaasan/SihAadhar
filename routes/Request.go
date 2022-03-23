package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishnuhaasan/PushNotifications/controllers"
)

var RegisterRequestRoutes = func(r *gin.Engine) {
	router := r.Group("request")
	{
		router.GET("/", controllers.GetRequests)
		router.POST("/", controllers.CreateRequest)
		router.PUT("/:aadhaar", controllers.UpdateRequest)
	}
}
