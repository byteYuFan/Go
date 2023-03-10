package service

import (
	"github.com/cloudRestaurant/dao"
	"github.com/cloudRestaurant/model"
)

// service的定义
type ShopService struct {
}

// 实例化service
func NewShopService() *ShopService {
	return &ShopService{}
}

/**
 * 返回商铺列表数据
 */
func (shopService *ShopService) ShopList(longtitude, latitude string) []model.Shop {
	shopDao := dao.NewShopDao()
	return shopDao.QueryShops(longtitude, latitude)
}

func (ShopService *ShopService) SearchShops(keyword string) []model.Shop {
	shopDao := dao.NewShopDao()
	return shopDao.QueryShopsByName(keyword)
}

func (ShopService *ShopService) GetShopService(shopId int64) []model.Service {
	shopDao := dao.NewShopDao()
	return shopDao.QueryShopsById(shopId)
}
