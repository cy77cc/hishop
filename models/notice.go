package models

type Notice struct {
	Model
	Content  string `json:"content" gorm:"type:varchar(255);default:'0'"`
	EndTime  int    `json:"end_time" gorm:"type:int"`
	IsDelete int8   `json:"is_delete" gorm:"type:tinyint"`
}
