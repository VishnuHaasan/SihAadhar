package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vishnuhaasan/PushNotifications/config"
	"github.com/vishnuhaasan/PushNotifications/routes"
)

func main() {
	godotenv.Load(".env")
	router := gin.Default()
	routes.RegisterUserRoutes(router)
	routes.RegisterRequestRoutes(router)
	cancel := config.ConnectDB()
	config.PingDb()
	defer cancel()
	server := &http.Server{
		Addr:           ":4999",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
