package models

import "gorm.io/gorm"

type SearchHistory struct {
	gorm.Model
	Keyword string `json:"keyword" gorm:"unique;type:char(50)"`
	From    string `json:"from" gorm:"type:varchar(45);default:''"`
	AddTime int    `json:"add_time" gorm:"type:int;default:0"`
	UserId  string `json:"user_id" gorm:"type:varchar(45)"`
}
