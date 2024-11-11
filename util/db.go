package util

import (
	"github.com/cy77cc/hioshop/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

// DBOpenShop 用于连接shop数据库
func DBOpenShop() *gorm.DB {
	dsn := "root:" + configs.SqlPassword + "@tcp(127.0.0.1:3306)/hiolabsdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "hiolabs_",
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("database connect shop error")
	}
	return db
}
