package routes

import (
	"online_market_project/pkg/controllers"
	"online_market_project/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	router.POST("/cart", controllers.AddToProduc)
	router.GET("/cart/:user_id", controllers.GetProduc)
	router.DELETE("/cart/:id", controllers.RemoveFromProduc)
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	auth := router.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/products", controllers.GetProducts)
		auth.POST("/products", controllers.CreateProduct)
		auth.GET("/profile/:id", controllers.GetProfile)
		auth.PUT("/profile/:id", controllers.UpdateProfile)
		auth.GET("/products/:id", controllers.GetProductByID)
		auth.PUT("/products/:id", controllers.UpdateProduct)
		auth.DELETE("/products/:id", controllers.DeleteProduct)
	}
}
