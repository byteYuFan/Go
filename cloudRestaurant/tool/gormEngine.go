package tool

import (
	"github.com/cloudRestaurant/model"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbEngine *GormEngine

type GormEngine struct {
	DB *gorm.DB
}

func NewGormEngine() (*GormEngine, error) {
	databaseConfig := _cfg.Database
	dsn := databaseConfig.User + ":" + databaseConfig.Password + "@(" +
		databaseConfig.Host + ":" + databaseConfig.Port + ")/" + databaseConfig.Name +
		"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), nil)
	if err != nil {
		return nil, err
	}
	//创建表 自动迁移
	err = db.AutoMigrate(
		new(model.SmsCode),
		new(model.Member),
		new(model.FoodCategory),
		new(model.Shop),
	)

	if err != nil {
		return nil, err
	}
	ge := new(GormEngine)
	ge.DB = db
	DbEngine = ge
	return ge, nil
}

/**
 * 向Shop表中插入初始测试数据
 */
func InitShopData() {
	shops := []model.Shop{
		model.Shop{Id: 1, Name: "嘉禾一品（温都水城）", Address: "北京市昌平区宏福苑温都水城F1", Longitude: 116.36868, Latitude: 40.10039,
			Phone: "13437850035", Status: 1, RecentOrderNum: 106, RatingCount: 961, Rating: 4.7, PromotionInfo: "欢迎光临，用餐高峰请提前下单，谢谢",
			OpeningHours: "8:30/20:30", IsNew: true, IsPremium: true, ImagePath: "", MinimumOrderAmount: 20, DeliveryFee: 5},
		model.Shop{Id: 479, Name: "杨国福麻辣烫", Address: "北京市市蜀山区南二环路天鹅湖万达广场8号楼1705室", Longitude: 117.22124, Latitude: 31.81948, Phone: "13167583411",
			Status: 1, RecentOrderNum: 755, RatingCount: 167, Rating: 4.2, PromotionInfo: "欢迎光临，用餐高峰请提前下单，谢谢", OpeningHours: "8:30/20:30",
			IsNew: true, IsPremium: true, ImagePath: "", MinimumOrderAmount: 20, DeliveryFee: 5},
		model.Shop{Id: 485, Name: "好适口", Address: "北京市海淀区西二旗大街58号", Longitude: 120.65355, Latitude: 31.26578, Phone: "12345678901",
			Status: 1, RecentOrderNum: 58, RatingCount: 576, Rating: 4.6, PromotionInfo: "欢迎光临，用餐高峰请提前下单，谢谢", OpeningHours: "8:30/20:30",
			IsNew: true, IsPremium: true, ImagePath: "", MinimumOrderAmount: 20, DeliveryFee: 5},
		model.Shop{Id: 486, Name: "东来顺旗舰店", Address: "北京市天河区东圃镇汇彩路38号1领汇创展商务中心401", Longitude: 113.41724, Latitude: 23.1127, Status: 1,
			Phone: "13544323775", RecentOrderNum: 542, RatingCount: 372, Rating: 4.2, PromotionInfo: "老北京正宗涮羊肉,非物质文化遗产",
			OpeningHours: "09:00/21:30", IsNew: true, IsPremium: true, ImagePath: "", MinimumOrderAmount: 20, DeliveryFee: 5},
		model.Shop{Id: 487, Name: "北京酒家", Address: "北京市海淀区上下九商业步行街内", Longitude: 113.24826, Latitude: 23.11488, Phone: "13257482341", Status: 0,
			RecentOrderNum: 923, RatingCount: 871, Rating: 4.2, PromotionInfo: "北京第一家传承300年酒家", OpeningHours: "8:30/20:30", IsNew: true, IsPremium: true, ImagePath: "",
			MinimumOrderAmount: 20, DeliveryFee: 5},
		model.Shop{Id: 488, Name: "和平鸽饺子馆", Address: "北京市越秀区德政中路171", Longitude: 113.27521, Latitude: 23.12092,
			Phone: "17098764762", Status: 1, RecentOrderNum: 483, RatingCount: 273, Rating: 4.2, PromotionInfo: "吃饺子就来和平鸽饺子馆", OpeningHours: "8:30/20:30",
			IsNew: true, IsPremium: true, ImagePath: "", MinimumOrderAmount: 20, DeliveryFee: 5}}
	tx := DbEngine.DB.Session(&gorm.Session{
		SkipDefaultTransaction: true,
	})
	tx.Begin()
	r := tx.Create(&shops)
	if r.RowsAffected < int64(len(shops)) {
		tx.Rollback()
		InitShopData()
	}
	tx.Commit()
}
