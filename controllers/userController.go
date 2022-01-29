package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密失败"})
		return
	}
	newUser := models.User{
		Name:      username,
		Telephone: telephone,
		Password:  string(hashPassword),
	}
	DB.Create(&newUser)

	ctx.JSON(http.StatusOK, gin.H{"code": "200", "msg": "创建用户成功！"})

}

func Login(ctx *gin.Context) {
	DB := common.GetDB()
	//获取数据
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	//数据校验
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}

	if len(password) <= 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码必须大于等于6位"})
		return
	}

	// 判断手机号是否存在
	var user models.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户存在"})
		return
	}

	//密码校验
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}

	//返回token
	ctx.JSON(http.StatusOK, gin.H{"code": "200", "msg": "登录成功"})

}

func isExistTelephone(db *gorm.DB, telephone string) bool {
	var user models.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false

}
