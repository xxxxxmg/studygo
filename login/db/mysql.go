package db

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlDB *gorm.DB

func InitMysql(host string) {
	var err error
	mysqlDB, err = gorm.Open(mysql.Open(host), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := mysqlDB.DB()
	mysqlDB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 auto_increment=1")
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(59 * time.Second)
}
func Mysql() *gorm.DB {

	return mysqlDB
}
