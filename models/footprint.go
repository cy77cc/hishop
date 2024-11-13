package models

import "gorm.io/gorm"

type Footprint struct {
	gorm.Model
	UserId  int `json:"userId" gorm:"type:int;default:0"`
	GoodsId int `json:"goodsId" gorm:"type:int;default:0"`
	AddTime int `json:"add_time" gorm:"type:int;default:0"`
}
