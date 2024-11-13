package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	GoodsId               uint    `json:"goods_id"`
	GoodsSpecificationIds string  `json:"goods_specification_ids"`
	GoodsSn               string  `json:"goods_sn"`
	GoodsNumber           uint    `json:"goods_number"`
	RetailPrice           float64 `json:"retail_price"`
	Cost                  float64 `json:"cost"`
	GoodsWeight           float32 `json:"goods_weight"`
	HasChange             int8    `json:"has_change"`
	GoodsName             string  `json:"goods_name"`
	IsOnSale              int8    `json:"is_on_sale"`
	IsDelete              int8    `json:"is_delete"`
}
