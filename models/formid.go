package models

import "gorm.io/gorm"

type Formid struct {
	gorm.Model
	UserId   int    `json:"user_id" gorm:"type:int"`
	OrderId  int    `json:"order_id" gorm:"type:int"`
	FormId   string `json:"form_id" gorm:"type:varchar(255)"`
	AddTime  int    `json:"add_time" gorm:"type:int;default:0"`
	UseTimes int8   `json:"use_times" gorm:"type:tinyint;default:0"`
}
