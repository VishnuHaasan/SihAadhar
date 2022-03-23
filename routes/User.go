package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishnuhaasan/PushNotifications/controllers"
)

var RegisterUserRoutes = func(r *gin.Engine) {
	router := r.Group("user")
	{
		router.POST("/", controllers.CreateUser)
		router.GET("/:aadhaar", controllers.GetUser)
		router.PUT("/:aadhaar/service", controllers.AddServiceProvider)
		router.PUT("/:aadhaar/service/toggle", controllers.ToggleIsActiveForProvider)
	}
}
