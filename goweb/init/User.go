package init

type User struct {
	UUID     int64  `gorm:"column:uuid"` // 主键
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	//创建时间，时间戳
	CreateTime int64 `gorm:"column:createtime"`
}
