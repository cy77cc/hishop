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

	// 数据库连接语句
	dbURi = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		configs.User,
		configs.Password,
		configs.Host,
		configs.Port,
		configs.Name)
	// new一个mysql连接
	dialector = mysql.New(mysql.Config{
		DSN:                       dbURi, // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	})

	// 用gorm打开这个连接
	conn, err := gorm.Open(dialector, &gorm.Config{ // 如果不用设置表前缀及复数，nil
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: configs.TablePrefix, // set table prefix
		},
	})
	if err != nil {
		log.Println(err.Error())
	}
	sqlDB, err := conn.DB()
	if err != nil {
		log.Println("connect db server failed.")
	}
	// SetMaxOpenConns 设置数据库的最大打开连接数。
	sqlDB.SetMaxOpenConns(100)
	// SetMaxIdleConns 设置空闲连接池中的最大连接数。
	sqlDB.SetMaxIdleConns(10)
	// SetConnMaxLifetime 设置连接可以重用的最长时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

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
