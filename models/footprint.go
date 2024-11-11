package models

type Footprint struct {
	Id      int `json:"id" gorm:"primary_key;AUTO_INCREMENT;type:int"`
	UserId  int `json:"userId" gorm:"type:int;default:0"`
	GoodsId int `json:"goodsId" gorm:"type:int;default:0"`
	AddTime int `json:"add_time" gorm:"type:int;default:0"`
}
