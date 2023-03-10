package service

import (
	"github.com/cloudRestaurant/dao"
	"github.com/cloudRestaurant/model"
)

type FoodCategoryService struct {
}

func NewFoodCategoryService() *FoodCategoryService {
	return &FoodCategoryService{}
}

func (fcs *FoodCategoryService) Categories() ([]model.FoodCategory, error) {
	//调用数据库层
	foodCategoryDao := dao.NewFoodCategoryDap()
	return foodCategoryDao.QueryCategories()
}
