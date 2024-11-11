package util

import (
	"fmt"
	"github.com/cy77cc/hioshop/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var db *gorm.DB

// gorm v2
func init() {
	var dbURi string
	var dialector gorm.Dialector

	dbURi = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		configs.User,
		configs.Password,
		configs.Host,
		configs.Port,
		configs.Name)
	dialector = mysql.New(mysql.Config{
		DSN:                       dbURi, // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	})

	conn, err := gorm.Open(dialector, &gorm.Config{ // 如果不用设置表前缀及复数，nil
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   configs.TablePrefix, // set table prefix
			SingularTable: true,                // set table singular
		},
	})
	if err != nil {
		log.Println(err.Error())
	}
	sqlDB, err := conn.DB()
	if err != nil {
		log.Println("connect db server failed.")
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(600 * time.Second)

	db = conn
}

// open api
func GetDB() *gorm.DB {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("connect db server failed.")
	}
	if err = sqlDB.Ping(); err != nil {
		sqlDB.Close()
	}

	return db
}
