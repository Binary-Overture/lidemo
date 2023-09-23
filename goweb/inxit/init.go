package inxit

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Iniit() {
	username := "root"
	password := "root"
	host := "127.0.0.1"
	port := 3306
	Dbname := "goweb"
	timeout := "10s"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接错误", err)
	}
	fmt.Println(db)
	DB = db
	err = DB.AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err)
	}
}
