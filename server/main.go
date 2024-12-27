package main

import (
	"clothingecommerce/config"
	"clothingecommerce/db"
	"clothingecommerce/routers"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)



func main() {
	db.InitDB()

	config.LoadEnv()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	routers.AuthRouter(router)
	routers.AdminRouter(router)
	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": "Access granted for API"})
	})

	log.Fatal(router.Run(":" + port))

}