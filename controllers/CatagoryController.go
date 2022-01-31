package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"learn/ginEssential/common"
	"learn/ginEssential/models"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() CategoryController {
	db := common.GetDB()
	_ = db.AutoMigrate(models.Category{})
	return CategoryController{DB: db}
}

func (c CategoryController) Create(ctx *gin.Context) {
	//绑定body中的参数
	var requestCategory models.Category
	ctx.Bind(&requestCategory)

	//数据校验
	if requestCategory.Name == "" {
		common.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}

	//创建
	c.DB.Create(&requestCategory)

	common.Success(ctx, gin.H{"category": requestCategory}, "创建成功")

}

func (c CategoryController) Update(ctx *gin.Context) {
	//绑定body 参数
	var requestCategory models.Category
	ctx.Bind(&requestCategory)

	//数据校验
	if requestCategory.Name == "" {
		common.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}

	//获取path 参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var updateCategory models.Category

	if err := c.DB.First(&updateCategory, categoryId).Error; err != nil {
		common.Fail(ctx, nil, "分类不存在")
		return
	}

	//更新操作
	c.DB.Model(&updateCategory).Update("name", requestCategory.Name)
	common.Success(ctx, nil, "更新成功")

}

func (c CategoryController) Show(ctx *gin.Context) {
	// 获取path 参数
	var category models.Category
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	// 数据校验
	if err := c.DB.First(&category, categoryId).Error; err != nil {
		common.Fail(ctx, nil, "分类不存在")
		return
	}

	// 展示操作
	common.Success(ctx, gin.H{"category": category}, "查询成功")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取path 参数
	var category models.Category
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	// 数据校验
	if err := c.DB.First(&category, categoryId).Error; err != nil {
		common.Fail(ctx, nil, "分类不存在")
		return
	}

	//删除操作
	c.DB.Delete(models.Category{}, categoryId)
	common.Success(ctx, nil, "删除成功")

}
