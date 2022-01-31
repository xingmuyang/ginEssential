package controllers

import (
	"github.com/gin-gonic/gin"
	"learn/ginEssential/common"
	"learn/ginEssential/models"
	"learn/ginEssential/repository"
	"learn/ginEssential/vo"
	"log"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	Repository repository.ICategoryRepository
}

func NewCategoryController() ICategoryController {
	categoryRepository := repository.NewCategoryRepository()
	categoryController := CategoryController{Repository: categoryRepository}
	_ = categoryController.Repository.(repository.CategoryRepository).DB.AutoMigrate(models.Category{})

	return categoryController
}

func (c CategoryController) Create(ctx *gin.Context) {
	//绑定body中的参数
	var requestCategory vo.CreateCategoryRequest

	if err := ctx.ShouldBind(&requestCategory); err != nil {
		common.Fail(ctx, nil, "数据验证错误")
		return
	}

	//创建
	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		common.Fail(ctx, nil, "数据验证错误")
		return
	}

	common.Success(ctx, gin.H{"category": category}, "创建成功")
}

func (c CategoryController) Update(ctx *gin.Context) {
	// 获取body 参数
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		log.Println(err.Error())
		common.Fail(ctx, nil, "数据验证错误")
		return
	}

	// 获取 path 参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var updateCategory *models.Category
	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		panic(err)
	}
	updateCategory, err = c.Repository.Update(*category, requestCategory.Name)
	if err != nil {
		panic(err)
	}

	common.Success(ctx, gin.H{"category": updateCategory}, "更新成功")

}

func (c CategoryController) Show(ctx *gin.Context) {
	// 获取path 参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		panic(err)
	}

	// 展示操作
	common.Success(ctx, gin.H{"category": category}, "查询成功")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取path 参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if err := c.Repository.DeleteById(categoryId); err != nil {
		panic(err)
	}

	common.Success(ctx, nil, "删除成功")

}
