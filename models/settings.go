package models

import "gorm.io/gorm"

type Settings struct {
	gorm.Model
	AutoDelivery       int8   `json:"auto_delivery" gorm:"type:tinyint;default:0"`
	Name               string `json:"name" gorm:"type:varchar(100)"`
	Tel                string `json:"tel" gorm:"type:varchar(20)"`
	ProvinceName       string `json:"province_name" gorm:"type:varchar(20)"`
	CityName           string `json:"city_name" gorm:"type:varchar(20)"`
	ExpAreaName        string `json:"exp_area_name" gorm:"type:varchar(20)"`
	Address            string `json:"address" gorm:"type:varchar(100)"`
	DiscoveryImgHeight int    `json:"discovery_img_height" gorm:"type:int;default:0"`
	DiscoveryImg       string `json:"discovery_img" gorm:"type:varchar(255)"`
	GoodsId            int    `json:"goods_id" gorm:"type:int;default:0"`
	CityId             int    `json:"city_id" gorm:"type:int;default:0"`
	ProvinceId         int    `json:"province_id" gorm:"type:int;default:0"`
	DistrictId         int    `json:"district_id" gorm:"type:int;default:0"`
	Countdown          int    `json:"countdown" gorm:"type:int;default:0"`
	Reset              int8   `json:"reset" gorm:"type:int;default:0"`
}
