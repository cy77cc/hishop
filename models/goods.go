package models

type Goods struct {
	Model
	CategoryId        uint   `json:"category_id" gorm:"default:0"`
	IsOnSale          uint8  `json:"is_on_sale" gorm:"default:1"`
	Name              string `json:"name"`
	GoodsNumber       uint   `json:"goods_number"`
	SellVolume        uint   `json:"sell_volume"`
	Keywords          string `json:"keywords"`
	RetailPrice       string `json:"retail_price"`
	MinRetailPrice    string `json:"min_retail_price"`
	CostPrice         string `json:"cost_price"`
	MinCostPrice      string `json:"min_cost_price"`
	GoodsBrief        string `json:"goods_brief"`
	GoodsDesc         string `json:"goods_desc" gorm:"type:text"`
	SortOrder         uint   `json:"sort_order"`
	IsIndex           int8   `json:"is_index"`
	IsNew             int8   `json:"is_new"`
	GoodsUnit         string `json:"goods_unit"`
	HttpsPicUrl       string `json:"https_pic_url"`
	ListPicUrl        string `json:"list_pic_url"`
	FreightTemplateId int    `json:"freight_template_id"`
	FreightType       int    `json:"freight_type"`
	IsDelete          uint8  `json:"is_delete"`
	HasGallery        int8   `json:"has_gallery"`
	HasDone           int8   `json:"has_done"`
}

type GoodsGallery struct {
	Model
	GoodsId   uint   `json:"goods_id"`
	ImgUrl    string `json:"img_url"`
	ImgDesc   string `json:"img_desc"`
	SortOrder uint   `json:"sort_order"`
	IsDelete  int8   `json:"is_delete"`
}

type GoodsSpecification struct {
	Model
	GoodsId         uint   `json:"goods_id"`
	SpecificationId uint   `json:"specification_id"`
	Value           string `json:"value"`
	PicUrl          string `json:"pic_url"`
	IsDelete        int8   `json:"is_delete"`
}
