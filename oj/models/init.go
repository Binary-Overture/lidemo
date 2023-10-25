package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB = Init()

func Init() *gorm.DB {
	username := "root"
	password := "root"
	host := "127.0.0.1"
	port := 3306
	Dbname := "gin_gorm_oj"
	timeout := "10s"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("gorm init error")
	}
	return db
}
