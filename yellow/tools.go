package yellow

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

type Video struct {
	ID       int       `gorm:"id"`
	Name     string    `gorm:"name,uniqueIndex"`
	URL      string    `gorm:"url"`
	Picture  string    `gorm:"picture"`
	Duration string    `gorm:"duration"` // 时长
	Heart    int       `gorm:"heart"`
	Views    int       `gorm:"views"`
	Publish  time.Time `gorm:"publish"`
}

func InitPG() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=127.0.0.1 port=5432 user=postgres password=123456 dbname=aaa sslmode=disable TimeZone=Asia/Shanghai")
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
		},
		Logger: gormLogger.Default.LogMode(gormLogger.Error), // info 打印sql日志
	}
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		panic(fmt.Sprintf("gorm database engine creation failed: %s", err.Error()))
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("gorm database ConnPool create failed: %s", err.Error()))
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(10)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err = sqlDB.Ping(); err != nil {
		panic("database ping failed:" + err.Error())
	}

	fmt.Println("successfully connect to postgresql database.")

	if err = db.AutoMigrate(&Video{}); err != nil {
		fmt.Printf("AutoMigrate error:" + err.Error())
	}
	return db
}
