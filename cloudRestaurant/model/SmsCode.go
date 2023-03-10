package model

// 短信结构体

type SmsCode struct {
	Id         int    `gorm:"pk auto_increment" json:"id"`
	Phone      string `gorm:"char(11)" json:"phone"`
	BizId      string `gorm:"varchar(30)" json:"bizId"`
	Code       string `gorm:"varchar(6)" json:"code"`
	CreateTime int64  `gorm:"bigint" json:"create_time"`
}
