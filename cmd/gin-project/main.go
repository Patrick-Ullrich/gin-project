package main

import (
	"example/gin-project/pkg/infrastructure/persistence"
	"example/gin-project/pkg/web/controllers"
	"net/http"

	"example/gin-project/pkg/web/middlewares"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "this is the login endpoint!"})
}

func main() {
	DB := persistence.Postgres_Init()
	c := controllers.New(DB)
	router := gin.Default()
	router.Use(middlewares.Logger())

	api := router.Group("/api/v1")

	api.POST("/register", c.Register)
	api.POST("/login", c.Login)

	router.Run(":8080")
}
