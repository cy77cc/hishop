package models

type Admin struct {
	Model
	Username      string `json:"username" gorm:"varchar(25);default:''"`
	Password      string `json:"password" gorm:"varchar(255);default:''"`
	PasswordSalt  string `json:"password_salt" gorm:"varchar(255);default:''"`
	LastLoginIp   string `json:"last_login_ip" gorm:"varchar(60);default:''"`
	LastLoginTime int    `json:"last_login_time" gorm:"int;default:0"`
	IsDelete      int8   `json:"is_delete" gorm:"tinyint;default:0"`
}
