package model

type Shop struct {
	//id
	Id int64 `gorm:"pk auto_increment" json:"id"`
	//商铺名称
	Name string `gorm:"varchar(12)" json:"name"`
	//宣传信息
	PromotionInfo string `gorm:"varchar(30)" json:"promotion_info"`
	//地址
	Address string `gorm:"varchar(100)" json:"address"`
	//联系电话
	Phone string `gorm:"varchar(11)" json:"phone"`
	//店铺营业状态
	Status int `gorm:"tinyint" json:"status"`
	//经度
	Longitude float64 `gorm:"" json:"longitude"`
	//纬度
	Latitude  float64 `gorm:"" json:"latitude"`
	ImagePath string  `gorm:"varchar(255)" json:"image_path"`
	//
	IsNew bool `gorm:"bool" json:"is_new"`
	//
	IsPremium bool `gorm:"bool" json:"is_premium"`
	//商铺评分
	Rating float32 `gorm:"float" json:"rating"`
	//评分总数
	RatingCount int64 `gorm:"int" json:"rating_count"`
	//当前订单总数
	RecentOrderNum int64 `gorm:"int" json:"recent_order_num"`

	//配送起送价
	MinimumOrderAmount int32 `gorm:"int" json:"minimum_order_amount"`
	//配送费
	DeliveryFee int32 `gorm:"int" json:"delivery_fee"`

	//营业时间
	OpeningHours string `gorm:"varchar(20)" json:"opening_hours"`
}
