package main

import (
	"github.com/gin-gonic/gin"
	"learn/ginEssential/common"
	"learn/ginEssential/routers"
)

func init() {
	common.InitDb()

}

func main() {
	r := gin.Default()
	routers.LoadRouter(r)
	panic(r.Run())
}
