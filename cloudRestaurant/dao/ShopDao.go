package dao

import (
	"github.com/cloudRestaurant/model"
	"github.com/cloudRestaurant/tool"
	"strconv"
)

type ShopDao struct {
	Engine *tool.GormEngine
}

func NewShopDao() *ShopDao {
	return &ShopDao{
		Engine: tool.DbEngine,
	}
}

const DEFAULT_RANGE = 5

func (shopDao *ShopDao) QueryShops(longitude, latitude string) []model.Shop {
	var shopList []model.Shop
	lo, _ := strconv.ParseFloat(longitude, 64)
	la, _ := strconv.ParseFloat(latitude, 64)
	shopDao.Engine.DB.Where("longitude>? AND longitude<? AND latitude>? AND latitude<?", lo-DEFAULT_RANGE, lo+DEFAULT_RANGE, la-DEFAULT_RANGE, la+DEFAULT_RANGE).Find(&shopList)
	return shopList
}
func (shaoDao *ShopDao) QueryShopsByName(keyword string) []model.Shop {
	var shops []model.Shop
	shaoDao.Engine.DB.Where("name=? AND status=?", keyword, 1).Find(&shops)
	return shops
}
func (shaoDao *ShopDao) QueryShopsById(shopId int64) []model.Service {
	var service []model.Service
	//shaoDao.Engine.DB.Table("services").Joins("INNER", "shop_service", " service.id = shop_services.service_id and shop_services.shop_id = ? ", shopId).Find(&service)
	//shaoDao.Engine.DB.Model(&model.Service{}).Select("services.*").Joins("JOIN shop_services ON services.id = shop_services.service_id and shop_services.shop_id =? ", shopId).Find(&service)
	//tx := shaoDao.Engine.DB.Joins("JOIN shop_services ON  shop_services.service_id=services.id  ").Where("shop_services.shop_id=?", shopId).Find(&service)
	//fmt.Println(tx.Error)
	shaoDao.Engine.DB.Table("services").Select("services.*").Joins("left join shop_services ON  shop_services.service_id=services.id and shop_services.shop_id =?", shopId).Find(&service)
	return service
}
