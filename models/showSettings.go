package models

type ShowSettings struct {
	Id             int   `json:"id" gorm:"primary_key;AUTO_INCREMENT;type:mediumint"`
	Banner         uint8 `json:"banner" gorm:"type:tinyint unsigned;default:0"`
	Channel        int8  `json:"channel" gorm:"type:tinyint;default:0"`
	IndexBannerImg int8  `json:"index_banner_img" gorm:"type:tinyint;default:0"`
	Notice         int8  `json:"notice" gorm:"type:tinyint;default:0"`
}