package models

type Cart struct {
	Model
	UserId                      uint    `json:"user_id" gorm:"type:mediumint unsigned;default:0"`
	GoodsId                     uint    `json:"goods_id" gorm:"type:mediumint unsigned;default:0"`
	GoodsSn                     string  `json:"goods_sn" gorm:"type:varchar(60);default:''"`
	ProductId                   uint    `json:"product_id" gorm:"type:mediumint unsigned;default:0"`
	GoodsName                   string  `json:"goods_name" gorm:"type:varchar(120);default:''"`
	GoodsAka                    string  `json:"goods_aka" gorm:"type:varchar(120);default:''"`
	GoodsWeight                 float32 `json:"goods_weight" gorm:"type:double(4,2);default:0.00"`
	AddPrice                    float64 `json:"add_price" gorm:"type:decimal(10,2);default:0.00"`
	RetailPrice                 float64 `json:"retail_price" gorm:"type:decimal(10,2);default:0.00"`
	Number                      uint8   `json:"number" gorm:"type:smallint unsigned;default:0"`
	GoodsSpecificationNameValue string  `json:"goods_specification_name_value" gorm:"type:text"`
	GoodsSpecificationIds       string  `json:"goods_specification_ids" gorm:"type:varchar(60);default:''"`
	Checked                     uint8   `json:"checked" gorm:"type:tinyint unsigned;default:1"`
	ListPicUrl                  string  `json:"list_pic_url" gorm:"type:varchar(255);default:''"`
	FreightTemplateId           uint    `json:"freight_template_id" gorm:"type:mediumint"`
	IsOnSale                    uint8   `json:"is_on_sale" gorm:"type:tinyint unsigned;default:1"`
	AddTime                     uint    `json:"add_time" gorm:"type:int;default:0"`
	IsFast                      int8    `json:"is_fast" gorm:"type:tinyint;default:0"`
	IsDelete                    uint8   `json:"is_delete" gorm:"type:tinyint unsigned;default:0"`
}
