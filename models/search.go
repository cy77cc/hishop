package models

type SearchHistory struct {
	Id      uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT;type:int unsigned"`
	Keyword string `json:"keyword" gorm:"unique;type:char(50)"`
	From    string `json:"from" gorm:"type:varchar(45);default:''"`
	AddTime int    `json:"add_time" gorm:"type:int;default:0"`
	UserId  string `json:"user_id" gorm:"type:varchar(45)"`
}
