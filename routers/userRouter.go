package routers

import (
	"github.com/gin-gonic/gin"
	"learn/ginEssential/controllers"
	"learn/ginEssential/middleware"

)


func LoadRouter(r *gin.Engine) {
	r.POST("/api/auth/register", controllers.Register)
	r.POST("/api/auth/login", controllers.Login)
	r.GET("/api/auth/userInfo", middleware.AuthMiddleware(), controllers.UserInfo)

	categoryRoutes := r.Group("/categories")
	categoryController := controllers.NewCategoryController()
	categoryRoutes.POST("", categoryController.Create)
	categoryRoutes.PUT(":id", categoryController.Update)
	categoryRoutes.GET(":id", categoryController.Show)
	categoryRoutes.DELETE(":id", categoryController.Delete)
}

