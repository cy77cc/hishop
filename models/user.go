package models

import (
	"time"
)

type User struct {
	Model
	Nickname      string    `json:"nickname"`
	Name          string    `json:"name"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Gender        uint8     `json:"gender"`
	Birthday      uint      `json:"birthday"`
	RegisterTime  time.Time `json:"register_time"`
	LastLoginTime time.Time `json:"last_login_time"`
	LastLoginIp   int       `json:"last_login_ip"`
	Mobile        string    `json:"mobile"`
	RegisterIp    int       `json:"register_ip"`
	Avatar        string    `json:"avatar"`
	WeixinOpenid  string    `json:"weixin_openid"`
	NameMobile    int8      `json:"name_mobile"`
	Country       string    `json:"country"`
	Province      string    `json:"province"`
	City          string    `json:"city"`
}
