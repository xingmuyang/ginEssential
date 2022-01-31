package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"learn/ginEssential/common"
	"learn/ginEssential/dto"
	"learn/ginEssential/models"
	"learn/ginEssential/utils"
	"log"
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
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}

	if len(password) <= 6 {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码必须大于等于6位")
		return
	}

	if len(username) == 0 {
		username = utils.RandomString(10)
	}

	//查询手机号
	if isExistTelephone(DB, telephone) {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已注册")
		return
	}

	//创建用户

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		common.Response(ctx, http.StatusInternalServerError, 500, nil, "加密失败")
		return
	}
	newUser := models.User{
		Name:      username,
		Telephone: telephone,
		Password:  string(hashPassword),
	}
	DB.Create(&newUser)
	// 发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		common.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error : %v", err)
		return
	}

	// 返回结果
	common.Success(ctx, gin.H{"token": token}, "注册成功")

}

func Login(ctx *gin.Context) {
	DB := common.GetDB()
	//获取数据
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	//数据校验
	if len(telephone) != 11 {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}

	if len(password) <= 6 {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码必须大于等于6位")
		return
	}

	// 判断手机号是否存在
	var user models.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}

	//密码校验
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		common.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}


	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		common.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error : %v", err)
		return
	}

	// 返回结果
	common.Success(ctx, gin.H{"token": token}, "登录成功")

}

func isExistTelephone(db *gorm.DB, telephone string) bool {
	var user models.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false

}

func UserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	log.Printf("%T", user)

	common.Success(ctx, gin.H{"user": dto.ToUserDto(user.(models.User))}, "获取用户信息成功")

}
