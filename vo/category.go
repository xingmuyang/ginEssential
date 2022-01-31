package vo

type CreateCategoryRequest struct {
	//免去 校验数据是否为空
	Name string `json:"name" binding:"required"`
}
