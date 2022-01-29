package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"learn/ginEssential/common"
	"learn/ginEssential/models"
	"learn/ginEssential/utils"
	"net/http"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
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
		username = utils.RandomString(10)
	}

	//查询手机号
	if isExistTelephone(DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已注册"})
		return
	}

	//创建用户

	newUser := models.User{
		Name:      username,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&newUser)

	ctx.JSON(http.StatusOK, gin.H{"code": "200", "msg": "创建用户成功！"})

}

func isExistTelephone(db *gorm.DB, telephone string) bool {
	var user models.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false

}
