package routers

import (
	"github.com/gin-gonic/gin"
	"learn/ginEssential/controllers"
)


func LoadRouter(r *gin.Engine) {
	r.POST("/api/auth/register", controllers.Register)
	r.POST("/api/auth/login", controllers.Login)

}