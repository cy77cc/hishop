package models

import "gorm.io/gorm"

type FreightTemplate struct {
	gorm.Model
	Name         string  `json:"name" gorm:"type:varchar(120);default:'0'"`
	PackagePrice float32 `json:"package_price" gorm:"type:decimal(10,2) unsigned;default:0.00"`
	FreightType  int8    `json:"freight_type" gorm:"type:tinyint;default:0"`
	IsDelete     int8    `json:"is_delete" gorm:"type:tinyint;default:0"`
}

type FreightTemplateDetail struct {
	gorm.Model
	TemplateId int  `json:"template_id" gorm:"primary_key;type:int;default:0"`
	GroupId    int  `json:"group_id" gorm:"type:int;default:0"`
	Area       int  `json:"area" gorm:"type:int;default:0"`
	IsDelete   int8 `json:"is_delete" gorm:"type:tinyint;default:0"`
}

type FreightTemplateGroup struct {
	gorm.Model
	TemplateId   int     `json:"template_id" gorm:"type:int;default:0"`
	IsDefault    int8    `json:"is_default" gorm:"type:tinyint;default:0"`
	Area         string  `json:"area" gorm:"type:varchar(3000);default:'0'"`
	Start        int     `json:"start" gorm:"type:int;default:0"`
	StartFee     float64 `json:"start_fee" gorm:"type:decimal(10,2);default:0.00"`
	Add          int     `json:"add" gorm:"type:int;default:1"`
	AddFee       float64 `json:"add_fee" gorm:"type:decimal(10,2);default:0.00"`
	FreeByNumber int8    `json:"free_by_number" gorm:"type:tinyint;default:0"`
	FreeByMoney  float64 `json:"free_by_money" gorm:"type:decimal(10,2);default:0.00"`
	IsDelete     int8    `json:"is_delete" gorm:"type:tinyint;default:0"`
}
