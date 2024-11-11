package models

type Shipper struct {
	Id           int    `json:"id" gorm:"primary_key;AUTO_INCREMENT;type:int"`
	Name         string `json:"name" gorm:"type:varchar(20);default:''"`
	Code         string `json:"code" gorm:"type:varchar(10);default:''"`
	SortOrder    int    `json:"sort_order" gorm:"type:int;default:10"`
	MonthCode    string `json:"month_code" gorm:"type:varchar(100)"`
	CustomerName string `json:"CustomerName" gorm:"type:varchar(100)"`
	Enabled      int8   `json:"enabled" gorm:"type:tinyint;default:0"`
}