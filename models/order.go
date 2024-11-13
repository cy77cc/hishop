package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderSn        string  `json:"order_sn"`
	UserId         uint    `json:"user_id"`
	OrderStatus    uint8   `json:"order_status"`
	OfflinePay     int8    `json:"offline_pay"`
	ShippingStatus uint8   `json:"shipping_status"`
	PrintStatus    int8    `json:"print_status"`
	PayStatus      uint8   `json:"pay_status"`
	Consignee      string  `json:"consignee"`
	Country        uint8   `json:"country"`
	Province       uint8   `json:"province"`
	City           uint8   `json:"city"`
	District       uint8   `json:"district"`
	Address        string  `json:"address"`
	PrintInfo      string  `json:"print_info"`
	Mobile         string  `json:"mobile"`
	Postscript     string  `json:"postscript"`
	AdminMemo      string  `json:"admin_memo"`
	ShippingFee    float64 `json:"shipping_fee"`
	PayName        string  `json:"pay_name"`
	PayId          string  `json:"pay_id"`
	ChangePrice    float64 `json:"change_price"`
	ActualPrice    float64 `json:"actual_price"`
	OrderPrice     float64 `json:"order_price"`
	GoodsPrice     float64 `json:"goods_price"`
	AddTime        uint    `json:"add_time"`
	PayTime        uint    `json:"pay_time"`
	ShippingTime   uint    `json:"shipping_time"`
	ConfirmTime    uint    `json:"confirm_time"`
	DealdoneTime   uint    `json:"dealdone_time"`
	FreightPrice   uint    `json:"freight_price"`
	ExpressValue   float64 `json:"express_value"`
	Remark         string  `json:"remark"`
	OrderType      uint8   `json:"order_type"`
	IsDelete       uint8   `json:"is_delete"`
}

type OrderExpress struct {
	Id           uint   `gorm:"primary_key;auto_increment" json:"id"`
	OrderId      uint   `json:"order_id"`
	ShipperId    uint   `json:"shipper_id"`
	ShipperName  string `json:"shipper_name"`
	ShipperCode  string `json:"shipper_code"`
	LogisticCode string `json:"logistic_code"`
	Traces       string `json:"traces"`
	IsFinish     uint8  `json:"is_finish"`
	RequestCount int    `json:"request_count"`
	RequestTime  int    `json:"request_time"`
	AddTime      int    `json:"add_time"`
	UpdateTime   int    `json:"update_time"`
	ExpressType  uint8  `json:"express_type"`
	RegionCode   string `json:"region_code"`
}

type OrderGoods struct {
	Id                          uint    `gorm:"primary_key;auto_increment" json:"id"`
	OrderId                     uint    `json:"order_id"`
	GoodsId                     uint    `json:"goods_id"`
	GoodsName                   string  `json:"goods_name"`
	GoodsAka                    string  `json:"goods_aka"`
	ProductId                   uint    `json:"product_id"`
	Number                      uint8   `json:"number"`
	RetailPrice                 float64 `json:"retail_price"`
	GoodsSpecificationNameValue string  `json:"goods_specification_name_value"`
	GoodsSpecificationIds       string  `json:"goods_specification_ids"`
	ListPicUrl                  string  `json:"list_pic_url"`
	IsDelete                    uint8   `json:"is_delete"`
}
