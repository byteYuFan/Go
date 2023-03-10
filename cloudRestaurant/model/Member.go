package model

// 会员数据结构体
type Member struct {
	Id           int64   `gorm:"pk auto_increment" json:"id"`
	UserName     string  `gorm:"varchar(20)" json:"user_name"`
	Mobile       string  `gorm:"char(11)" json:"mobile"`
	Password     string  `gorm:"varchar(255);default:''" json:"password"`
	RegisterTime int64   `gorm:"bigint" json:"register_time"`
	Avatar       string  `gorm:"varchar(255)" json:"avatar"`
	Balance      float64 `gorm:"double" json:"balance"`
	IsActive     int8    `gorm:"tinyint" json:"is_active"`
	City         string  `gorm:"varchar(10)" json:"city"`
}
