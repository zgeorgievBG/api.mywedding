package main

import (
	"api.mywedding/controllers"
	"api.mywedding/database"
	"api.mywedding/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.Migrate()
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/users/register", controllers.RegisterUser)
		api.POST("/users/token", controllers.GenerateToken)
		api.GET("/partners", controllers.GetAllPartnersCards)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.POST("/partners", controllers.CreatePartnerCard)
			secured.PUT("/partners/:id", controllers.UpdatePartnerCard)
			secured.DELETE("/partners/:id", controllers.DeletePartner)
		}
	}
	return router
}

// user can publish when he have some status
