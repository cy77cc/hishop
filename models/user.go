package models

type User struct {
	Id            uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Nickname      string `json:"nickname"`
	Name          string `json:"name"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Gender        uint8  `json:"gender"`
	Birthday      uint   `json:"birthday"`
	RegisterTime  uint   `json:"register_time"`
	LastLoginTime uint   `json:"last_login_time"`
	LastLoginIp   string `json:"last_login_ip"`
	Mobile        string `json:"mobile"`
	RegisterIp    string `json:"register_ip"`
	Avatar        string `json:"avatar"`
	WeixinOpenid  string `json:"weixin_openid"`
	NameMobile    int8   `json:"name_mobile"`
	Country       string `json:"country"`
	Province      string `json:"province"`
	City          string `json:"city"`
}
