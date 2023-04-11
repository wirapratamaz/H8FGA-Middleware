package router

import (
	"jwtgo/controllers"
	"jwtgo/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegistrasi)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.GET("/", controllers.GetUser)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:productId", middlewares.ProductAuthorization(), controllers.UpdateProduct)
		productRouter.GET("/", middlewares.Authentication(), controllers.GetProduct)
		productRouter.DELETE("/:productId", middlewares.ProductAuthorization(), controllers.DeleteProduct)
	}

	roleRouter := r.Group("/roles")
	{
		roleRouter.Use(middlewares.Authentication())
		roleRouter.POST("/", controllers.AddRole)
	}

	return r
}
