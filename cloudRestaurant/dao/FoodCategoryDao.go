package dao

import (
	"github.com/cloudRestaurant/model"
	"github.com/cloudRestaurant/tool"
)

type FoodCategoryDao struct {
	Engine *tool.GormEngine
}

// NewFoodCategoryDap 实例化一个对象
func NewFoodCategoryDap() *FoodCategoryDao {
	return &FoodCategoryDao{
		Engine: tool.DbEngine,
	}
}

// QueryCategories 从数据库中查询数据种类
func (fcd *FoodCategoryDao) QueryCategories() ([]model.FoodCategory, error) {
	var category []model.FoodCategory

	result := fcd.Engine.DB.Find(&category)
	return category, result.Error
}
