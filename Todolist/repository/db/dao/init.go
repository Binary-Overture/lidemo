package dao

import (
	"Todolist/config"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

var _db *gorm.DB

func MysqlInit() {
	//设置默认Mysql配置
	mConfig := config.Config.Mysql["default"]
	//连接mysql
	conn := strings.Join([]string{mConfig.UserName, ":", mConfig.Password, "@tcp(", mConfig.DbHost,
		":", mConfig.DbPort, ")/", mConfig.DbName, "?charset=", mConfig.Charset, "&parseTime=true"}, "")
	var ormLogger = logger.Default
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       conn,
		DefaultStringSize:         256,   //string默认字段
		DisableDatetimePrecision:  true,  // mysql5.6 版本之前不支持，禁用datetime精度
		DontSupportRenameIndex:    true,  //重命名索引的时候采用删除并重建的方式，mysql5.7之前不支持
		DontSupportRenameColumn:   true,  //mysql8之前不支持
		SkipInitializeWithVersion: false, //不根据版本进行自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, //表名命名不自动加s
		},
	})
	if err != nil {
		panic(err)
	}
	//设置连接池
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(20)  //设置连接池
	sqlDb.SetMaxOpenConns(100) //设置最大打开数
	sqlDb.SetConnMaxLifetime(time.Second * 30)
	_db = db
	autoMigrate()
}
