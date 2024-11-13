package models

import "gorm.io/gorm"

type Keywords struct {
	gorm.Model
	Keyword   string `json:"keyword" gorm:"type:varchar(90);default:''"`
	IsHot     uint8  `json:"is_hot" gorm:"type:tinyint unsigned;default:0"`
	IsDefault uint8  `json:"is_default" gorm:"type:tinyint unsigned;default:0"`
	IsShow    uint8  `json:"is_show" gorm:"type:tinyint unsigned;default:1"`
	SortOrder uint   `json:"sort_order" gorm:"type:int unsigned;default:100"`
	SchemeUrl string `json:"scheme_url" gorm:"type:varchar(255);default:''"`
	Type      uint   `json:"type" gorm:"type:int unsigned;default:0"`
}
