package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"learn/ginEssential/common"
	"learn/ginEssential/routers"
	"os"
)

func init() {
	InitConfig()
	common.InitDb()
}

func InitConfig()  {
	workDir, _ := os.Getwd()
	viper.AddConfigPath(workDir + "/config")
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func run(r *gin.Engine) {
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func main() {
	r := gin.Default()
	routers.LoadRouter(r)
	run(r)
}


