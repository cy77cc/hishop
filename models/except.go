package models

type ExceptArea struct {
	Id       uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT;type:mediumint unsigned"`
	Content  string `json:"content" gorm:"type:varchar(255)"`
	Area     string `json:"area" gorm:"type:varchar(3000);default:''"`
	IsDelete int8   `json:"is_delete" gorm:"type:tinyint;default:0"`
}

type ExceptAreaDetail struct {
	Id           uint `json:"id" gorm:"primary_key;AUTO_INCREMENT;type:mediumint unsigned"`
	ExceptAreaId int  `json:"except_area_id" gorm:"type:int;default:0"`
	Area         int  `json:"area" gorm:"type:int;default:0"`
	IsDelete     int8 `json:"is_delete" gorm:"type:tinyint;default:0"`
}
