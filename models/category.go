package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name       string `json:"name" gorm:"varchar(90);default:''"`
	Keywords   string `json:"keywords" gorm:"varchar(255);default:''"`
	FrontDesc  string `json:"front_desc" gorm:"varchar(255);default:''"`
	ParentId   uint   `json:"parent_id" gorm:"type:int unsigned;default:0"`
	SortOrder  uint8  `json:"sort_order" gorm:"tinyint unsigned;default:50"`
	ShowIndex  int8   `json:"show_index" gorm:"tinyint;default:0"`
	IsShow     uint8  `json:"is_show" gorm:"tinyint unsigned;default:1"`
	IconUrl    string `json:"icon_url" gorm:"varchar(255);default:''"`
	ImgUrl     string `json:"img_url" gorm:"varchar(255);default:''"`
	Level      string `json:"level" gorm:"varchar(255);default:''"`
	FrontName  string `json:"front_name" gorm:"varchar(255);default:''"`
	PHeight    int    `json:"p_height" gorm:"int;default:0"`
	IsCategory int8   `json:"is_category" gorm:"tinyint;default:0"`
	IsChannel  int8   `json:"is_channel" gorm:"tinyint;default:0"`
}
