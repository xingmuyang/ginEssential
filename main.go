package main

import (
	"github.com/gin-gonic/gin"
	"learn/ginEssential/common"
	"learn/ginEssential/controllers"
)



func main() {
	r := gin.Default()
	common.InitDb()

	r.POST("/api/auth/register", controllers.Register)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}





