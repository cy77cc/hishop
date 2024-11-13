package models

import "gorm.io/gorm"

type Region struct {
	gorm.Model
	ParentId uint16 `json:"parent_id" gorm:"type:smallint unsigned;default:0"`
	Name     string `json:"name" gorm:"type:varchar(120);default:''"`
	Type     int8   `json:"type" gorm:"type:tinyint;default:2"`
	AgencyId uint16 `json:"agency_id" gorm:"type:smallint unsigned;default:0"`
	Area     uint16 `json:"area" gorm:"type:smallint unsigned;default:0"`
	AreaCode string `json:"area_code" gorm:"type:varchar(10);default:'0'"`
	FarArea  uint   `json:"far_area" gorm:"type:int unsigned;default:0"`
}
