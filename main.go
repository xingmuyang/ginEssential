package main

import (
	"github.com/gin-gonic/gin"
	"learn/ginEssential/common"
	"learn/ginEssential/routers"
)

func main() {
	common.InitDb()
	r := gin.Default()
	routers.LoadRouter(r)

	panic(r.Run())
}
