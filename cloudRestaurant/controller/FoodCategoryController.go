package controller

import (
	"github.com/cloudRestaurant/service"
	"github.com/cloudRestaurant/tool"
	"github.com/gin-gonic/gin"
)

type FoodCategoryController struct {
}

func (fcc *FoodCategoryController) Router(engine *gin.Engine) {
	//美食类别
	engine.GET("/api/food_category", fcc.foodCategory)
}

// foodCategory 获取全部食品种类
func (fcc *FoodCategoryController) foodCategory(ctx *gin.Context) {
	//调用service层数据获取
	foodCategoryService := &service.FoodCategoryService{}
	categories, err := foodCategoryService.Categories()
	if err != nil {
		tool.Failed(ctx, "获取食品列表失败")
		return
	}
	//格式转化
	// imgUrl:hello.png
	for _, category := range categories {
		if category.ImageUrl != "" {
			category.ImageUrl = tool.FileServerAddr() + "/" + category.ImageUrl
		}
	}
	tool.Success(ctx, categories)
}
