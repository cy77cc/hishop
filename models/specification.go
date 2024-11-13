package models

type Specification struct {
	Model
	Name      string `json:"name" gorm:"type:varchar(60);default:''"`
	SortOrder uint   `json:"sort_order" gorm:"type:int unsigned;default:0"`
	Memo      string `json:"memo" gorm:"type:varchar(255);default:''"`
}
