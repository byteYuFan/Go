package controller

import (
	"github.com/cloudRestaurant/service"
	"github.com/cloudRestaurant/tool"
	"github.com/gin-gonic/gin"
)

type ShopController struct {
}

func (sc *ShopController) Router(app *gin.Engine) {
	app.GET("/api/shops", sc.GetShopList)
	app.GET("/api/search_shops", sc.SearchShop)
}
func (sc *ShopController) GetShopList(ctx *gin.Context) {
	//调用服务层
	longitude, _ := ctx.GetQuery("longitude")
	latitude, _ := ctx.GetQuery("latitude")
	if longitude == "" || longitude == "undefined" || latitude == "" || latitude == "undefined" {
		//默认
		longitude = "116.36868"
		latitude = "40.10039"
	}
	shopService := service.NewShopService()
	shopList := shopService.ShopList(longitude, latitude)
	if shopList != nil {
		tool.Success(ctx, shopList)
		return
	}
	tool.Failed(ctx, "无结果")

}
func (sc *ShopController) SearchShop(ctx *gin.Context) {
	keyword, ok := ctx.GetQuery("keyword")
	if !ok {
		keyword = "好适口"
	}
	shopService := service.NewShopService()
	shops := shopService.SearchShops(keyword)
	if shops == nil {
		tool.Failed(ctx, "暂无商家信息")
		return
	}
	//查询shop 支持服务信息

	tool.Success(ctx, shops)
}
