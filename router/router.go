package router

import (
	"example_middleware/controllers"
	"example_middleware/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StarrtApp() *gin.Engine{
	r:=gin.Default()
	r.GET("/coba", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "berhasil")
	})
	userRouter:=r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}
	productRouter:=r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:productId", middlewares.ProductAuthorization(), controllers.UpdateProduct)
	}
	return r
}