package main

import (
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"time"
)

func main() {
	// :clickhouse://localhost:8123
	dsn := "clickhouse://@localhost:8123/gorm?dial_timeout=10s&read_timeout=20s"
	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// 自动迁移 (这是GORM自动创建表的一种方式--译者注)
	db.AutoMigrate(&User{})
	// 设置表选项
	db.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&User{})

	user := User{}
	// 插入
	db.Create(&user)

	// 查询
	db.Find(&user, "id = ?", 10)

	// 批量插入
	//var users = []User{user1, user2, user3}
	//db.Create(&users)
	// ...
}

type User struct {
	gorm.Model
	Username string
	Password string
	Age      int
	Birthday time.Time
	Phone    string
	Email    string
	Gender   int
}
