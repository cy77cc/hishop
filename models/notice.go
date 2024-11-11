package models

type Notice struct {
	Id       int    `json:"id" gorm:"primary_key;AUTO_INCREMENT;type:int"`
	Content  string `json:"content" gorm:"type:varchar(255);default:'0'"`
	EndTime  int    `json:"end_time" gorm:"type:int"`
	IsDelete int8   `json:"is_delete" gorm:"type:tinyint"`
}
