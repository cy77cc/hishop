package models

import "gorm.io/gorm"

type Ad struct {
	gorm.Model
	LinkType  uint   `json:"link_type" gorm:"type:tinyint;default:0"`
	Link      string `json:"link" gorm:"varchar(255);default:''"`
	GoodsId   int    `json:"goods_id" gorm:"type:int;default:0"`
	ImageUrl  string `json:"image_url" gorm:"text;"`
	EndTime   int    `json:"end_time" gorm:"type:int;default:0"`
	Enabled   uint8  `json:"enabled" gorm:"type:tinyint unsigned;default:0"`
	SortOrder int8   `json:"sort_order" gorm:"type:tinyint;default:0"`
	IsDelete  int8   `json:"is_delete" gorm:"type:tinyint;default:0"`
}
