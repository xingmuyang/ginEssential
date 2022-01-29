package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name      string
	Telephone string
	Password  string
}

func main() {
	r := gin.Default()
	db := InitDb()

	r.POST("/api/auth/register", func(ctx *gin.Context) {

		//获取参数
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		telephone := ctx.PostForm("telephone")

		//数据校验
		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
			return
		}

		if len(password) <= 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码必须大于等于6位"})
			return
		}

		if len(username) == 0 {
			username = RandomString(10)
		}

		//查询手机号
		if isExistTelephone(db, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已注册"})
			return
		}

		//创建用户

		newUser := User{
			Name:      username,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)

		ctx.JSON(http.StatusOK, gin.H{"code": "200", "msg": "创建用户成功！"})

	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

func InitDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("ginEssential.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
	return db
}

func isExistTelephone(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false

}
