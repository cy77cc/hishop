package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Name       string `gorm:"type:varchar(50);default:''" json:"name"`
	UserId     uint   `json:"user_id" gorm:"type:mediumint;default:0"`
	CountryId  int16  `json:"country_id" gorm:"type:smallint;default:0"`
	ProvinceId int16  `json:"province_id" gorm:"type:smallint;default:0"`
	CityId     int16  `json:"city_id" gorm:"type:smallint;default:0"`
	DistrictId int16  `json:"district_id" gorm:"type:smallint;default:0"`
	Address    string `json:"address" gorm:"type:varchar(120);default:''"`
	Mobile     string `json:"mobile" gorm:"type:varchar(60);default:''"`
	IsDefault  uint8  `json:"is_default" gorm:"type:tinyint unsigned;default:0"`
	IsDelete   int8   `json:"is_delete" gorm:"type:tinyint;default:0"`
}
